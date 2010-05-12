/*
 Referee package for Rock Paper Scissors
*/

package rps

import "games"
import "strconv"
import "strings"
import . "sstruct"

//Ref: referee function for rps
//aIn = Player A's input channel (send commands TO A)
//aOut = Player A's output channel (receives commands FROM A)
//bIn = Player B's input channel
//bOut = Player B's output channel
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
			
			aResult := strconv.Itoa((int) (winner(aMove, bMove)))
			bResult := strconv.Itoa((int) (winner(bMove, aMove)))
			
			aIn <- Make(aResult)
			bIn <- Make(bResult)
		}
	}
}

//determine the win state for the game
//(determines based on "a" value)
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
