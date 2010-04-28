/*
Authors: William Broza, Tym Lipari
Matrix Testing program

Written in a pattern of parallel turntaking.

usage:
	matrix
*/

package main

//import "games"
import "os"
import "io"
import "fmt"
import "reflect"

type GView struct {
	inOut io.ReadWriter
	name, other string
	refComm chan string
}

//bool gametype = false; //false for RPS, True for TTT

func NewGView(inout io.ReadWriter, n string, ref chan string) *RPSView {
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
) else {
	buffer := make([]byte, 2048)
		tempString := ""
		n, _ := inOut.Read(buffer)
		for {
			if n <= 0 { break; }
			tempString = tempString + ((string)buffer[0:n])
			n, _ = inOut.Read(buffer)
		}
		return tempString
	}
}

func (this *GView) Loop() os.Error {
	done := false
	for !done {
		command := <- this.refComm
		switch command {
		case "enable": this.Enable()
		case "get": {
			val := this.Get()
			switch reflect.Typeof(val){ 
			//is an os.Error
			case *reflect.StructType: return val
			case *reflect.StringType: this.refComm <- val.(string)
			}
		}
		case "display": this.Display()
		case "quit": done = true
		}
	}
	return nil
}
			
func main() {
	view := NewGView(os.Stdout, "A")
	go view.Loop()
	//println(view.Get().(string))
}
