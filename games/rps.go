/*
Authors: William Broza, Tym Lipari
Rock Paper Scissors game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	rps-stdin [terminal]
*/

package main

import "games"
import "os"
//import "io"
import "strconv"
import "view"
import "strings"
			
func main() {

	term := os.Stdin
	var err os.Error
	
	println(strconv.Itoa(len(os.Args)))
	if len(os.Args) > 1 {
		term, err = os.Open(os.Args[1], os.O_RDWR, 511)
		if err != nil { term = os.Stdin }
	}
	
	aComm := make(chan string)
	bComm := make(chan string)
	
	//investigate adding quit. right now, it crashes
	//because all the goroutines are asleep.
	a := view.NewGView(os.Stdin, "A", "r, p, s", aComm)
	b := view.NewGView(term, "B", "r, p, s", bComm)
	
	go a.Loop()
	go b.Loop()
	
	stillPlaying := true
	
	for stillPlaying {
		aComm <- "enable"
		bComm <- "enable"
		
		aComm <- "get"
		bComm <- "get"
		
		aMove := strings.TrimSpace(<- aComm)
		if aMove == "q" { bComm <- "quit" }
		
		bMove := strings.TrimSpace(<- bComm) 
		if bMove == "q" { aComm <- "quit" }
		
		if stillPlaying {
			aComm <- "other"
			bComm <- "other"
		
			aComm <- "B's move:  " + bMove + "\n"
			bComm <- "A's move:  " + aMove + "\n"
		
			aComm <- "display"
			bComm <- "display"
			
			aComm <- "result"
			bComm <- "result"
			
			aResult := strconv.Itoa((int) (winner(aMove, bMove))) + "\n"
			bResult := strconv.Itoa((int) (winner(bMove, aMove))) + "\n"
			
			aComm <- aResult
			bComm <- bResult
		}
	}
}

func winner(a, b string) games.Outcome {
	if a == b { 
		return games.Draw 
	} else if a == "p" && b == "r" {
		return games.Win
	} else if a == "r" && b == "s" {
		return games.Win
	} else if a == "s" && b == "p" {
		return games.Win 
	}
	return games.Lose
}
