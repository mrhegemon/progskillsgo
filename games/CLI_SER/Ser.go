package main

import (
    "strconv"
    "netchan"
    "os"
)

type stringStruct struct {
	s string
}


func main() {
	ch := make(chan stringStruct)

	ex, err := netchan.NewExporter("tcp", ":4567") 	
	if err != nil {
 		println("exportSend:" + err.String())
 		os.Exit(0)
 	}
  	err2 := ex.Export("CliExp", ch, netchan.Send, new(stringStruct))
 	if err2 != nil {
 		println("exportSend:" + err2.String())
 		os.Exit(0)
 	}
 	
	for x := 0; x < 10; x++ { 
		var v stringStruct
		v.s = "Happy Pants" + strconv.Itoa(x)
		ch <- v
	}

	imp, err3 := netchan.NewImporter("tcp", ":4568")
	if err3 != nil {
		println(err3.String())
		os.Exit(0)
	}
	chn := make(chan stringStruct)
	err4 := imp.Import("CliImp", chn, netchan.Recv, new(stringStruct))
	if err4 != nil {
		println(err4.String())
		os.Exit(0)
	}
	
	for y := 0; y < 10; y++ {
  		var v stringStruct
		v = <-chn
		println(v.s)
  	}
}

