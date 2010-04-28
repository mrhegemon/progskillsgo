/*
Authors: William Broza, Tym Lipari
Tic Tac Toe game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	matrix
*/

package view

import "games"
import "os"
import "io"
import "fmt"
import "strconv"

type GView struct {
	inOut io.ReadWriter
	name, other string
	refComm chan string
	instr chan string
	directions string
}

//inout = input/output
//n = this view's name
//dir = the directions for this game
//ref = the channel that Loop() will communicate with
func NewGView(inout io.ReadWriter, n, dir string, ref chan string) *GView {
	view := new(GView)
	view.inOut = inout
	view.name = n
	view.other = ""
	view.refComm = ref
	view.directions = dir
	return view
}

func (this *GView) Enable() {
	text := ([]byte) (this.name + "'s move (" + this.directions + "):  ")
	if _, err := this.inOut.Write(text); err != nil {
		fmt.Fprintln(os.Stderr, "Error Writing To Stream (" + this.name + ")")
	}
}

func(this *GView) Set(move interface{}) {
	this.other = move.(string)
}

func(this *GView) Get() interface{} {
if this == nil {
	return nil
} else {
	buffer := make([]byte, 2048)
		tempString := ""
		n, _ := this.inOut.Read(buffer)
		for {
			temp := (string) (buffer[0:n])
			tempString = tempString + temp
			if n < 2048 { return tempString }
			n, _ = this.inOut.Read(buffer)
		}
		return tempString
	}
return ""
}

func (this *GView) Loop() os.Error {
	done := false
	for !done {
		command := <- this.refComm
		switch command {
		case "enable": this.Enable()
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

func(this *GView) Display() {
	this.inOut.Write(([]byte) (this.other))
}

func(this *GView) Done(youWin games.Outcome) {
	if youWin == games.Win {
		this.inOut.Write(([]byte) (this.name + " won.\n"))
	} else if youWin == games.Draw {
		this.inOut.Write(([]byte) ("There was a tie.\n"))
	} else {
		this.inOut.Write(([]byte) (this.name + " lost.\n"))
	}
}