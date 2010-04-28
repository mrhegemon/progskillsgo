/*
Authors: William Broza, Tym Lipari
Tick Tack Toe game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	ttt-stdin [terminal]
*/

package main

import "games"
import "os"
//import "io"
import "strconv"
import "view"
import "strings"
import "container/vector"
			
var (
	x int = 3
	y int = 3
	mtx []*vector.StringVector
	full int = 0
)

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
	a := view.NewGView(os.Stdin, "A", "Directions: Ex: nw, s, c ", aComm)
	b := view.NewGView(term, "B", "Directions: Ex: nw, s, c ", bComm)
	
	go a.Loop()
	go b.Loop()
	
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

	resA, resB := "", ""

	for stillPlaying {

		AMOVE: aComm <- "enable"
		aComm <- "get"

		aMove := strings.TrimSpace(<- aComm)
		if aMove == "q" { bComm <- "quit" }

		if setGame(aMove, "A") != nil {
			//tell A it's move was bad
			//GO BACK and repeat A's move
			goto AMOVE
		}
		bComm <- "other"
		bComm <- "A's move:  " + aMove + "\n"
		bComm <- "display"

		//check for win state
		winA := winner("A")
		winB := winner("B")
		if winA == games.Lose && winB == games.Lose && full == 9 {
			resA = strconv.Itoa((int) (games.Draw)) + "\n"
			resB = strconv.Itoa((int) (games.Draw)) + "\n"
			stillPlaying = false
		} else if winA == games.Win {
			resA = strconv.Itoa((int) (games.Win)) + "\n"
			resB = strconv.Itoa((int) (games.Lose)) + "\n"
			stillPlaying = false
		} else if winB == games.Win {
			resA = strconv.Itoa((int) (games.Win)) + "\n"
			resB = strconv.Itoa((int) (games.Lose)) + "\n"
			stillPlaying = false
		}

		if stillPlaying {
			BMOVE: bComm <- "enable"

			bComm <- "get"
		
			bMove := strings.TrimSpace(<- bComm) 
			if bMove == "q" { aComm <- "quit" }
		
			if setGame(aMove, "B") != nil {
				//tell B it's move was bad
				//GO BACK and repeat B's move
				goto BMOVE
			}

			aComm <- "other"
			aComm <- "B's move:  " + bMove + "\n"

			aComm <- "display"

			//check for win state
			winA := winner("A")
			winB := winner("B")
			if winA == games.Lose && winB == games.Lose && full == 9 {
				resA = strconv.Itoa((int) (games.Draw)) + "\n"
				resB = strconv.Itoa((int) (games.Draw)) + "\n"
				stillPlaying = false
			} else if winA == games.Win {
				resA = strconv.Itoa((int) (games.Win)) + "\n"
				resB = strconv.Itoa((int) (games.Lose)) + "\n"
				stillPlaying = false
			} else if winB == games.Win {
				resA = strconv.Itoa((int) (games.Win)) + "\n"
				resB = strconv.Itoa((int) (games.Lose)) + "\n"
				stillPlaying = false
			}
		}	
	}
	aComm <- "result"
	bComm <- "result"
	aComm <- resA
	bComm <- resB
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
