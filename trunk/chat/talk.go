package main

import (
  "ws_daemon"
  "flag"
  "os"
)

func main() {
  var err os.Error
  switch {
  case flag.NArg() == 0:
    println("usage: talk [host] port\n")
    os.Exit(1)
  
  case flag.NArg() == 1:
    err = daemon.Client("ws:\\localhost:" + flag.Arg(0), "http:\\localhost:"+flag.Arg(0), os.Stdout, os.Stdin)

  default:
    err = daemon.Client("ws:\\" + flag.Arg(0) + ":" + flag.Arg(1), "http:\\" + flag.Arg(0)+":"+flag.Arg(1), os.Stdout, os.Stdin)
  }
  if err != nil { println(err.String()); os.Exit(1) }
}

