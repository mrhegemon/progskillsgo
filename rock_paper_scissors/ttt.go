package main

//import "games"
import "os"
import "io"
import "fmt"

type RPSView struct {
	inOut io.ReadWriter
	name string
}

func NewRPSView(inout io.ReadWriter, n string) *RPSView {
	view := new(RPSView)
	view.inOut = inout
	view.name = n
	return view
}

func (this *RPSView) Enable() {
	text := ([]byte) (this.name + "'s move (r, p, s):  ")
	if _, err := this.inOut.Write(text); err != nil {
		fmt.Fprintln(os.Stderr, "Error Writing To Stream (" + this.name + ")")
	}
}

/*func(this *RPSView) Get() interface{} {
	buffer := make([]byte, 2048)
	tempString := ""
	
	n, _ := inOut.Read(buffer)
	for {
		if n <= 0 { break; }
		tempString = tempString + ((string)buffer[0:n])
		n, _ = inOut.Read(buffer)
	}
	
	return tempString
}*/

func main() {
	view := NewRPSView(os.Stdout, "A")
	view.Enable()
	//println(view.Get().(string))
}
