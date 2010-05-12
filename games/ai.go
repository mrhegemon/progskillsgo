/*
Authors: William Broza, Tym Lipari
Tic Tac Toe game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	matrix
*/

package ai

import ("games"; "os"; "io"; "fmt"; "strconv"; "rand")

var(
	rps [3]string = [3]string{"r", "p", "s"}
	ttt [9]string = [9]string{"nw", "n", "ne",
					 "w", "c", "e",
					 "sw", "s", "se"}
)

type AIView struct {
	inOut io.ReadWriter
	name, other string
	refComm chan string
	instr chan string
	directions string
	isRps bool
}

//inout = input/output
//n = this view's name
//dir = the directions for this game
//ref = the channel that Loop() will communicate with
func NewAIView(isRps bool, inout io.ReadWriter, n, dir string, ref chan string) *AIView {
	view := new(AIView)
	view.inOut = inout
	view.name = n
	view.other = ""
	view.refComm = ref
	view.directions = dir
	view.isRps = isRps
	return view
}

func (this *AIView) Enable() {
	text := ([]byte) (this.name + "'s move (" + this.directions + "):  ")
	if _, err := this.inOut.Write(text); err != nil {
		fmt.Fprintln(os.Stderr, "Error Writing To Stream (" + this.name + ")")
	}
}

func(this *AIView) Set(move interface{}) {
	this.other = move.(string)
}

func(this *AIView) Get() interface{} {
	move := ""
	if this.is {
		move = rps[rand.Intn(len(rps))]
	}
	else {
		move = ttt[rand.Intn(len(ttt))]
	}
	this.inOut.Write(([]byte)(move))
	return move
}

func (this *AIView) Loop() os.Error {
	done := false
	for !done {
		command := <- this.refComm
		switch command {
		case "enable": this.Enable()
		case "name": {
			nname := <- this.refComm
			this.name = nname
		}
		case "get": {
			val := this.Get()
			switch val.(string) { 
			//is an os.Error
			case "ICLOSED": {
				this.refComm <- "q"
				return os.NewError("Input stream has been closed")
			}
			case "OCLOSED": {
				this.refComm <- "q"
				return os.NewError("Output stream has been closed")
			}
			default: this.refComm <- val.(string)
			}
		}
		case "other": this.Set(<- this.refComm)
		case "display": this.Display()
		case "result":{
			val := <- this.refComm
			intVal, _ := strconv.Atoi(val)
			this.Done(games.Outcome(intVal))
		}
		case "quit": done = true
	}
	}
	return nil
}

func(this *AIView) Display() {
	this.inOut.Write(([]byte) (this.other))
}

func(this *AIView) Done(youWin games.Outcome) {
	if youWin == games.Win {
		this.inOut.Write(([]byte) (this.name + " won.\n"))
	} else if youWin == games.Draw {
		this.inOut.Write(([]byte) ("There was a tie.\n"))
	} else {
		this.inOut.Write(([]byte) (this.name + " lost.\n"))
	}
}
