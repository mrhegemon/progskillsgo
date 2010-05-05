/*
Authors: William Broza, Tym Lipari
Rock Paper Scissors game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	rps-stdin [terminal]
*/
package main

import ("flag"; "os"; "view"; "rps"; "netchan")
import . "sstruct"

func main() {
	var server *bool = flag.Bool("c", false, "client mode")
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
		aTemp := make(chan StringStruct)
		bTemp1 := make(chan StringStruct)
		bTemp2 := make(chan StringStruct)
		aComm := &aTemp
		bIn := &bTemp1
		bOut := &bTemp2
		
		val, err := netchan.NewExporter("tcp", flag.Arg(0))
		if err != nil {
			println(err.String())
			os.Exit(-1)
		}
		err2 := val.Export("BOut", bOut, netchan.Send, new(StringStruct))

		val2, err3 := netchan.NewImporter("tcp", flag.Arg(1))
		if err3 != nil {
			println(err3.String())
			os.Exit(-1)
		}
		err4 := netchan.Import("BIn", bIn, netchan.Recv, new(StringStruct))

		aView := view.NewGView(os.Stdin, "A", "r, p, s", *aComm, *aComm)
		go aView.Loop()
		rps.Ref(*aComm, *bComm)
	} else {   //client mode
		iChan := make(chan StringStruct)
		oChan := make(chan StringStruct)
		

		imp, err := netchan.NewImporter("tcp", flag.Arg(0))
		if err != nil {
			println(err.String())
			os.Exit(-1)
		}
		imp.Import("BIn", iChan, netchan.Recv, new(StringStruct))

		exp, err3 := netchan.NewExporter("tcp", flag.Arg(0))
		if err3 != nil {
			println(err3.String())
			os.Exit(-1)
		}
		exp.Export("BOut", oChan, netchan.Send, new(StringStruct))

		myView := view.NewGView(os.Stdin, "B", "r, p, s", myChan)
		go myView.Loop()
		for { }
	}
}
