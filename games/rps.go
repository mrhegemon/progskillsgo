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
import . "sstruct"
			
func Ref(aIn, aOut, bIn, bOut chan StringStruct) {
	
	stillPlaying := true
	
	for stillPlaying {
		aIn <- Make("enable")
		bIn <- Make("enable")
		
		aIn <- Make("get")
		bIn <- Make("get")
		
		aMove := strings.TrimSpace((<- aOut).S)
		if aMove == "q" { bOut <- Make("quit") }
		
		bMove := strings.TrimSpace((<- bOut).S) 
		if bMove == "q" { aIn <- Make("quit") }
		
		if stillPlaying {
			aIn <- Make("other")
			bIn <- Make("other")
		
			aIn <- Make("B's move:  " + bMove + "\n")
			bIn <- Make("A's move:  " + aMove + "\n")
		
			aIn <- Make("display")
			bIn <- Make("display")
			
			aIn <- Make("result")
			bIn <- Make("result")
			
			aResult := Make(strconv.Itoa((int) (winner(aMove.S, bMove.S))))
			bResult := Make(strconv.Itoa((int) (winner(bMove.S, aMove.S))))	
			
			aIn <- Make(aResult)
			bIn <- Make(bResult)
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
