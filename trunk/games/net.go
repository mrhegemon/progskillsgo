package main

import ("flag"; "os"; "view"; "rps"; "netchan")

func main() {
	var server *bool = flag.Bool("s", true, "server mode")
	var usage *bool = flag.Bool("x", false, "print usage")
	
	flag.Parse()

	if *usage {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if flag.NArg() == 0 {  //not enough parameters
		//print usage
		os.Exit(-1);
	} else if *server {  //server mode
		aTemp := make(chan string)
		bTemp := make(chan string)
		aComm := &aTemp
		bComm := &bTemp
		
		val, err := netchan.NewExporter("tcp", flag.Arg(0))
		if err != nil {
			println(err.String())
			os.Exit(-1)
		}
		val.Export("B", bComm, netchan.Send, bComm)

		if flag.NArg() == 1 {
			aView := view.NewGView(os.Stdin, "A", "r, p, s", *aComm)
			go aView.Loop()
		} else {
			a, err := netchan.NewImporter("tcp", flag.Arg(1))
			if err != nil {
				println(err.String())
				os.Exit(-1)
			}
			a.Import("A", aComm, netchan.Recv, aComm)
		}
		rps.Ref(*aComm, *bComm)
	} else {   //client mode
		myChan := make(chan string)

		imp, err := netchan.NewImporter("tcp", flag.Arg(0))
		if err != nil {
			println(err.String())
			os.Exit(-1)
		}
		imp.Import("B", myChan, netchan.Recv, myChan)

		myView := view.NewGView(os.Stdin, "B", "r, p, s", myChan)
		go myView.Loop()
	}
}
