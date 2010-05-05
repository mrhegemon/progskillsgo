/*
Authors: William Broza, Tym Lipari
Rock Paper Scissors game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	rps-stdin [terminal]
*/

package main

import ("rps"; "os"; "view")
import . "sstruct"

func main() {
	term := os.Stdin
	var err os.Error
	
	if len(os.Args) > 1 {
		term, err = os.Open(os.Args[1], os.O_RDWR, 511)
		if err != nil { term = os.Stdin }
	}
	
	aComm := make(chan StringStruct)
	bComm := make(chan StringStruct)
	
	//investigate adding quit. right now, it crashes
	//because all the goroutines are asleep.
	a := view.NewGView(os.Stdin, "A", "r, p, s", aComm, aComm)
	b := view.NewGView(term, "B", "r, p, s", bComm, bComm)
	
	go a.Loop()
	go b.Loop()

	rps.Ref(aComm, bComm)
}
