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

	im, err := netchan.NewImporter("tcp", ":4567")
	err2 := im.Import("CliExp", ch, netchan.Recv, new(stringStruct))
	
	if err != nil {
        println("new importer:" + err.String())
        os.Exit(0)
  	}
  	if err2 != nil {
  		println(err2.String())
  		os.Exit(0)
  	}

  	for y := 0; y < 10; y++ {
  		var v stringStruct
		v = <-ch
		println(v.s)
  	}

  	exp, err3 := netchan.NewExporter("tcp", ":4568")
	if err3 != nil {
		println(err3.String())
		os.Exit(0)
	}
	chn := make(chan stringStruct)
	err4 := exp.Export("CliImp", chn, netchan.Send, new(stringStruct))
	if err4 != nil {
		println(err4.String())
		os.Exit(0)
	}
 	
	for x := 0; x < 10; x++ { 
		var v stringStruct
		v.s = "Happy Pants" + strconv.Itoa(x)
		chn <- v
	}
}

