package main

import (
    "strconv"
    "netchan"
    "os"
)

func main() {
	ch := make(chan string)

	ex, err := netchan.NewExporter("tcp", ":4567") 	
	if err != nil {
 		println("exportSend:" + err.String())
 		os.Exit(0)
 	}
  err2 := ex.Export("CliExp", ch, netchan.Send, new(string))
 	if err2 != nil {
 		println("exportSend:" + err2.String())
 		os.Exit(0)
 	}
 	
	for x := 0; x < 10; x++ { 
		ch <-("Happy Pants" + strconv.Itoa(x))
	}
	
  for y := 0; y < 10; y++ {
		v := <-ch
		println(v)
  }
}

