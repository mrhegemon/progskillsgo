                                                                     
                                                                     
                                                                     
                                             
/*
Authors: William Broza, Tym Lipari
Rock Paper Scissors game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	rps-stdin /dev/ttys000 or /dev/ttys001
*/

package main

//import "games"
import "os"
import "io"
import "fmt"

type GView struct {
	inOut io.ReadWriter
	name, other string
	refComm chan string
	instr chan string
}
			
func main() {
	commChan := make(chan string)
	view := NewGView(os.Stdout, "A", commChan)
	go view.Loop()
	//println(view.Get().(string))
}
