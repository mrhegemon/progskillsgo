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

func main() {
	term := os.Stdin
	var err os.Error
	
	if len(os.Args) > 1 {
		term, err = os.Open(os.Args[1], os.O_RDWR, 511)
		if err != nil { term = os.Stdin }
	}
	
	aComm := make(chan string)
	bComm := make(chan string)
	
	//investigate adding quit. right now, it crashes
	//because all the goroutines are asleep.
	a := view.NewGView(os.Stdin, "A", "n, s, e, w, c, nw, ne, sw, se", aComm)
	b := view.NewGView(term, "B", "n, s, e, w, c, nw, ne, sw, se", bComm)
	
	go a.Loop()
	go b.Loop()

	ttt.Ref(aComm, bComm)
}
