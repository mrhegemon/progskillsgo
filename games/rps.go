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
import . "sstruct"
import "strconv"
import "strings"
			
func Ref(aComm, bComm chan StringStruct) {
	
	stillPlaying := true
	
	for stillPlaying {
		aComm <- Make("enable")
		bComm <- Make("enable")
		
		aComm <- Make("get")
		bComm <- Make("get")
		
		var aMove := Make(strings.TrimSpace((<- aComm).S))
		if aMove.S == "q" { bComm <- Make("quit") }
		
		var bMove := Make(strings.TrimSpace((<- bComm).S)) 
		if bMove.S == "q" { aComm <- Make("quit") }
		
		if stillPlaying {
			aComm <- Make("other")
			bComm <- Make("other")
		
			aComm <- Make("B's move:  " + bMove + "\n")
			bComm <- Make("A's move:  " + aMove + "\n")
		
			aComm <- Make("display")
			bComm <- Make("display")
			
			aComm <- Make("result")
			bComm <- Make("result")
			
			aResult := Make(strconv.Itoa((int) (winner(aMove.S, bMove.S))))
			bResult := Make(strconv.Itoa((int) (winner(bMove.S, aMove.S))))	
			
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
