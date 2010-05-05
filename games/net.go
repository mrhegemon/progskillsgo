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
		aComm := make(chan StringStruct)
		bIn := make(chan StringStruct)
		bOut := make(chan StringStruct)

		err := os.NewError("Not Nil")
		val := new(netchan.Exporter)
		for err != nil {
			val, err = netchan.NewExporter("tcp", flag.Arg(0))
			//errchk(err)
		}
		errchk(val.Export("BOut", &bOut, netchan.Send, new(StringStruct)))

		imp := new(netchan.Importer)
		err = os.NewError("Not Nil")
		for err != nil {
			imp, err = netchan.NewImporter("tcp", flag.Arg(1))
			//errchk(err2)
		}
		errchk(imp.Import("BIn", &bIn, netchan.Recv, new(StringStruct)))

		aView := view.NewGView(os.Stdin, "A", "r, p, s", aComm, aComm)
		go aView.Loop()
		
		rps.Ref(aComm, aComm, bIn, bOut)
	} else {   //client model
		iChan := make(chan StringStruct)
		oChan := make(chan StringStruct)

		imp := new(netchan.Importer)
		err := os.NewError("Not Nil")
		for err != nil {
			imp, err = netchan.NewImporter("tcp", flag.Arg(0))
			//errchk(err)
		}
		errchk(imp.Import("BOut", &oChan, netchan.Recv, new(StringStruct)))
		
		exp := new(netchan.Exporter)
		err = os.NewError("Not Nil")
		for err != nil {
			exp, err = netchan.NewExporter("tcp", flag.Arg(1))
			//errchk(err2)
		}
		errchk(exp.Export("BIn", &iChan, netchan.Send, new(StringStruct)))

		myView := view.NewGView(os.Stdin, "B", "r, p, s", iChan, oChan)
		go myView.Loop()
		for { }
	}
}
func errchk(err os.Error) {
	if err != nil {
		println(err.String())
		os.Exit(-1)
	}
}
