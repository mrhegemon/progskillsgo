/*
Authors: William Broza, Tym Lipari
Tick Tack Toe game

Written in a pattern of serial turn taking.  TTT must extend 
asymmetrically when using channels over a network.

usage:
	matrix
*/

package main

//import "games"
import "os"
import "io"
import "fmt"

type GView struct {
	inOut io.ReadWriter
	name, other string
	refComm chan string
}

//bool gametype = false; //false for RPS, True for TTT

func NewGView(inout io.ReadWriter, n string, ref chan string) *GView {
	view := new(GView)
	view.inOut = inout
	view.name = n
	view.other = ""
	view.refComm = ref
	return view
}

func (this *GView) Enable() {
	text := ([]byte) (this.name + "'s move (r, p, s):  ")
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
			if n <= 0 { return tempString }
			temp := (string) (buffer[0:n])
			tempString = tempString + temp
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
			case "ICLOSED": return os.NewError("Input stream has been closed")
			case "OCLOSED": return os.NewError("Output stream has been closed")
			default: this.refComm <- val.(string)
			}
		}
		case "display": this.Display()
		case "quit": done = true
		}
	}
	return nil
}

func(this *GView) Display() {
	//display opponent's move
}
			
func main() {
	commChan := make(chan string)
	view := NewGView(os.Stdout, "A", commChan)
	go view.Loop()
	//println(view.Get().(string))
}
