/*
Authors: William Broza, Tym Lipari
Tic Tac Toe game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	matrix
*/

package view

import . "sstruct"
import ("games"; "os"; "io"; "fmt"; "strconv")

type GView struct {
	inOut io.ReadWriter
	name, other string
	irefComm, orefComm chan StringStruct
	instr chan string
	directions string
}

//inout = input/output
//n = this view's name
//dir = the directions for this game
//iref = the channel that Loop() will communicate with for input
//oref = the channel that Loop() will communicate with for output
func NewGView(inout io.ReadWriter, n, dir string, iref, oref chan StringStruct) *GView {
	view := new(GView)
	view.inOut = inout
	view.name = n
	view.other = ""
	view.irefComm = iref
	view.orefComm = oref
	view.directions = dir
	return view
}

//Enables the view, and prompts the user with directions
func (this *GView) Enable() {
	text := ([]byte) (this.name + "'s move (" + this.directions + "):  ")
	if _, err := this.inOut.Write(text); err != nil {
		fmt.Fprintln(os.Stderr, "Error Writing To Stream (" + this.name + ")")
	}
}

//Sets the opponent's move
func(this *GView) Set(move interface{}) {
	this.other = move.(string)
}

//Gets the user's move and returns it
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


//Runs the view
func (this *GView) Loop() os.Error {
	done := false
	for !done {
		command := (<- this.irefComm).S
		switch command {
		case "enable": this.Enable()
		case "name": {
			nname := (<- this.irefComm).S
			this.name = nname
		}
		case "get": {
			val := this.Get()
			switch val.(string) { 
			//is an os.Error
			case "ICLOSED": {
				var out StringStruct
				out.S = "q"
				this.orefComm <- out
				return os.NewError("Input stream has been closed")
			}
			case "OCLOSED": {
				var out StringStruct
				out.S = "q"
				this.orefComm <- out
				return os.NewError("Output stream has been closed")
			}
			default: 
				var out StringStruct
				out.S = val.(string) 
				this.orefComm <- out
			}
		}
		case "other": this.Set((<- this.irefComm).S)
		case "display": this.Display()
		case "result":{
			val := (<- this.irefComm).S
			intVal, _ := strconv.Atoi(val)
			this.Done(games.Outcome(intVal))
		}
		case "quit": done = true
	}
	}
	return nil
}

//Display's the opponent's move
func(this *GView) Display() {
	this.inOut.Write(([]byte) (this.other))
}

//Tells the user the end state
func(this *GView) Done(youWin games.Outcome) {
	if youWin == games.Win {
		this.inOut.Write(([]byte) (this.name + " won.\n"))
	} else if youWin == games.Draw {
		this.inOut.Write(([]byte) ("There was a tie.\n"))
	} else {
		this.inOut.Write(([]byte) (this.name + " lost.\n"))
	}
}
