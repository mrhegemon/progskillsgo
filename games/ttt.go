/*
Authors: William Broza, Tym Lipari
Tick Tack Toe game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	ttt-stdin [terminal]
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

func Ref(aComm, bComm chan StringStruct) {
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
	
			AMOVE: aComm <- Make("enable")
			aComm <- Make("get")
	
			aMove := Make(strings.TrimSpace(<- aComm))
			if aMove.S == "q" { bComm <- Make("quit") }
	
			if setGame(aMove.S, "A") != nil {
				//tell A it's move was bad
				//GO BACK and repeat A's move
				goto AMOVE
			}
			bComm <- Make("other")
			bComm <- Make("A's move:  " + aMove + "\n")
			bComm <- Make("display")
	
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
				BMOVE: bComm <- Make("enable")
	
				bComm <- Make("get")
			
				bMove := strings.TrimSpace(<- bComm) 
				if bMove.S == "q" { aComm <- "quit" }
			
				if setGame(bMove.S, "B") != nil {
					//tell B it's move was bad
					//GO BACK and repeat B's move
					goto BMOVE
				}
	
				aComm <- Make("other")
				aComm <- Make("B's move:  " + bMove + "\n")
	
				aComm <- Make"display")
	
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
	
		aComm <- Make("result")
		bComm <- Make("result")
		aComm <- Make(resA)
		bComm <- Make(resB)
	}
}

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
