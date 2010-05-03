/*
Authors: William Broza, Tym Lipari
Rock Paper Scissors game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	rps-stdin [terminal]
*/

package rps

import "games"
//import "io"
import "strconv"
import "strings"
			
func Ref(aComm, bComm chan string) {
	
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
			
			aResult := strconv.Itoa((int) (winner(aMove, bMove)))
			bResult := strconv.Itoa((int) (winner(bMove, aMove)))	
			
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
