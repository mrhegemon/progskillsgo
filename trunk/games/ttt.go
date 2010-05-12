/*
Authors: William Broza, Tym Lipari
Tick Tack Toe game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

*/

package ttt

import "games"
import "os"
import . "sstruct"
import "strconv"
import "strings"
import "container/vector"
			
var (
	x int = 3
	y int = 3
	mtx []*vector.StringVector
	full int = 0
)

//game referee
func Ref(aIn, aOut, bIn, bOut chan StringStruct) {
	for {
		stillPlaying := true
	
		//makes gameboard
		mtx = make([] *vector.StringVector, x)
		//populates gameboard
		for n, _ := range mtx {
			mtx[n] = new(vector.StringVector)
			for b := 0; b < y; b++ {
				mtx[n].Push("")
			}
		}
	
		var resA, resB string = "", ""
	
		for stillPlaying {
	
			AMOVE: aIn <- Make("enable")
			aIn <- Make("get")
	
			aMove := strings.TrimSpace((<- aOut).S)
			if aMove == "q" { bIn <- Make("quit") }
	
			if setGame(aMove.S, "A") != nil {
				println("goto")
				//tell A it's move was bad
				//GO BACK and repeat A's move
				goto AMOVE
			}
			bIn <- Make("other")
			bIn <- Make("A's move:  " + aMove + "\n")
			bIn <- Make("display")
	
			//check for win state
			winA := winner("A")
			winB := winner("B")
			if winA == games.Lose && winB == games.Lose && full == 9 {
				resA = strconv.Itoa((int) (games.Draw))
				resB = strconv.Itoa((int) (games.Draw))
				stillPlaying = false
			} else if winA == games.Win {
				resA = strconv.Itoa((int) (games.Win))
				resB = strconv.Itoa((int) (games.Lose))
				stillPlaying = false
			} else if winB == games.Win {
				resA = strconv.Itoa((int) (games.Win)) 
				resB = strconv.Itoa((int) (games.Lose))
				stillPlaying = false
			}
	
			if stillPlaying {
				BMOVE: bIn <- Make("enable")
	
				bIn <- Make("get")
			
				bMove := strings.TrimSpace((<- bOut).S)
				if bMove == "q" { aIn <- Make("quit") }
			
				if setGame(bMove.S, "B") != nil {
					//tell B it's move was bad
					//GO BACK and repeat B's move
					goto BMOVE
				}
	
				aIn <- Make("other")
				aIn <- Make("B's move:  " + bMove + "\n")
	
				aIn <- Make("display")
	
				//check for win state
				winA := winner("A")
				winB := winner("B")
				if winA == games.Lose && winB == games.Lose && full == 9 {
					resA = strconv.Itoa((int) (games.Draw))
					resB = strconv.Itoa((int) (games.Draw))
					stillPlaying = false
				} else if winA == games.Win {
					resA = strconv.Itoa((int) (games.Win))
					resB = strconv.Itoa((int) (games.Lose))
					stillPlaying = false
				} else if winB == games.Win {
					resA = strconv.Itoa((int) (games.Win))
					resB = strconv.Itoa((int) (games.Lose))
					stillPlaying = false
				}
			}	
		}
	
		aIn <- Make("result")
		bIn <- Make("result")
		aIn <- Make(resA)
		bIn <- Make(resB)
	}
}

//sets a move in the gameboard and returns error if filled
func setGame(m, p string) os.Error {
	switch m {	
		case "n" : {
			if mtx[0].At(1) != "" {
				return os.NewError("Space Filled!")
			} else {
					mtx[0].Set(1, p)
			}
		}
		case "nw" : {
			if mtx[0].At(0) != "" {
				return os.NewError("Space Filled!")
			} else {
					mtx[0].Set(0, p)
			}
		}
		case "w" : {
			if mtx[1].At(0) != "" {
				return os.NewError("Space Filled!")
			} else {
					mtx[1].Set(0, p)
			}
		}
		case "sw" : {
			if mtx[2].At(0) != "" {
				return os.NewError("Space Filled!")
			} else {
					mtx[2].Set(0, p)
			}
		}
		case "s" : {
			if mtx[2].At(1) != "" {
				return os.NewError("Space Filled!")
			} else {
					mtx[2].Set(1, p)
			}
		}
		case "se" : {
			if mtx[2].At(2) != "" {
				return os.NewError("Space Filled!")
			} else {
					mtx[2].Set(2, p)
			}
		}
		case "e" : {
			if mtx[1].At(2) != "" {
				return os.NewError("Space Filled!")
			} else {
					mtx[1].Set(2, p)
			}
		}
		case "ne" : {
			if mtx[0].At(2) != "" {
				return os.NewError("Space Filled!")
			} else {
					mtx[0].Set(2, p)
			}
		}
		case "c" : {
			if mtx[1].At(1) != "" {
				return os.NewError("Space Filled!")
			} else {
					mtx[1].Set(1, p)
			}
		}
		default : {
			return os.NewError("Illegal Move")
			full--
		}
	}
	full++
	return nil
}
	
//checks board for a winner
func winner(name string) games.Outcome {
	//TTT GAME LOGIC
	winner := false
	for row := 0; row < x; row++ {
		if (mtx[row].At(0) == name && mtx[row].At(1) == name && mtx[row].At(2) == name) {
			winner = true
		}
		if (mtx[0].At(row) == name && mtx[1].At(row) == name && mtx[2].At(row) == name) {
			winner = true
		}
	}
	if (mtx[0].At(0) == name && mtx[1].At(1) == name && mtx[2].At(2) == name) {
			winner = true
	}
	if (mtx[2].At(0) == name && mtx[1].At(1) == name && mtx[0].At(2) == name) {
			winner = true
	}

	//if winner
	if winner {
		return games.Win
	}
	return games.Lose
}
