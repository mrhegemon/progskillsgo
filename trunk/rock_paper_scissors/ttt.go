package main

//import "games"
import "os"
import "io"
import "fmt"
import "reflect"

type RPSView struct {
	inOut io.ReadWriter
	name, other string
	refComm chan string
}

func NewRPSView(inout io.ReadWriter, n string, ref chan string) *RPSView {
	view := new(RPSView)
	view.inOut = inout
	view.name = n
	view.other = ""
	view.refComm = ref
	return view
}

func (this *RPSView) Enable() {
	text := ([]byte) (this.name + "'s move (r, p, s):  ")
	if _, err := this.inOut.Write(text); err != nil {
		fmt.Fprintln(os.Stderr, "Error Writing To Stream (" + this.name + ")")
	}
}

func(this *RPSView) Set(move interface{}) {
	this.other = move.(string)
}

func(this *RPSView) Get() interface{} {
	return nil
	/*buffer := make([]byte, 2048)
	tempString := ""
	
	n, _ := inOut.Read(buffer)
	for {
		if n <= 0 { break; }
		tempString = tempString + ((string)buffer[0:n])
		n, _ = inOut.Read(buffer)
	}
	
	return tempString
}*/
}
func (this *RPSView) Loop() os.Error {
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
	view := NewRPSView(os.Stdout, "A")
	go view.Loop()
	//println(view.Get().(string))
}
