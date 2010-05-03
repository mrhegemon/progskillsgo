package main

import (
    "strconv"
    "netchan"
    "os"
)

func main() {
	ch := make(chan string)

	im, err := netchan.NewImporter("tcp", ":0")
	im.Import("CliExp", ch, netchan.Recv, new(string))
	if err != nil {
        println("new importer:" + err.String())
        os.Exit(0)
  }

  for y := 0; y < 10; y++ {
		v := <-ch
		println(v)
  }
 	
	for x := 0; x < 10; x++ { 
		ch <-("Happy Pants" + strconv.Itoa(x))
	}
}

