/*
Authors: William Broza, Tym Lipari
Rock Paper Scissors game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	rps-stdin [terminal]
*/
package main

import ("ttt"; "os"; "view")
import . "sstruct"

func main() {
	term := os.Stdin
	var err os.Error
	
	//reads arguments
	if len(os.Args) > 1 {
		term, err = os.Open(os.Args[1], os.O_RDWR, 511)
		if err != nil { term = os.Stdin }
	}
	
	aComm := make(chan StringStruct)
	bComm := make(chan StringStruct)
	
	//instructions to be sent
	a := view.NewGView(os.Stdin, "A", "n, s, e, w, c, nw, ne, sw, se", aComm, aComm)
	b := view.NewGView(term, "B", "n, s, e, w, c, nw, ne, sw, se", bComm, bComm)
	
	go a.Loop()
	go b.Loop()

	ttt.Ref(aComm, aComm, bComm, bComm)
}
