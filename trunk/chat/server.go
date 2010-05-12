package chat

import (
	"Vector"
	0

type peer struct {
	vector queue

/*package main

import (
  "daemon"
  "flag"
  "os"
  "os/signal"
  "syscall"
)

const (
  DATE = "/bin/date"
  FINGER = "/usr/bin/finger"
  SHELL = "/bin/sh -i"
  UNAME = "/usr/bin/uname -a"
  UPTIME = "/usr/bin/uptime"
  WHO = "/usr/bin/who"
)

func main() {
  date := flag.String("date", "", "port/path for date server, if any")
  echo := flag.String("echo", "", "port/path for echo server, if any")
  finger := flag.String("finger", "", "port/path for finger server, if any")
  shell := flag.String("shell", "", "port/path for shell server (ouch!), if any")
  uname := flag.String("uname", "", "port/path for uname server, if any")
  uptime := flag.String("uptime", "", "port/path for uptime server, if any")
  who := flag.String("who", "", "port/path for who server, if any")
  
  requests := flag.Int("r", 0, "limits requests, if > 0")
  sessions := flag.Int("s", -1, "limits concurrent sessions, if > 0")
  unix := flag.Bool("u", false, "use unix domain sockets")
  flag.Parse()
  
  network := "tcp"; pre := ":"
  
  var remove chan string // deferred removals, if unix
  if *unix {
    network = "unix"; pre = ""
    remove = make(chan string, flag.NFlag()) // buffered for all
  }

  nServers := 0 // count for termination
  errors := make(chan os.Error)
  
  if len(*echo) > 0 {
    nServers++
    go func() {
      echod, err := daemon.Echod(network, pre+*echo,
        *requests, *sessions, *echo+":")
      if err == nil {
        if *unix { remove <- *echo }
        err = echod.Run()
      }
      errors <- err
    }()
  }
  
  d := func(name string, port *string) {
    if len(*port) > 0 {
      nServers++
      go func() {
        d, err := daemon.Daemon(name, network, pre+*port, *requests, *sessions)
        if err == nil {
          if *unix { remove <- *port }
          err = d.Run()
        }
        errors <- err
      }()
    }
  }
  d(DATE, date)
  d(SHELL, shell)
  d(UNAME, uname)
  d(UPTIME, uptime)
  d(WHO, who)
  
  d1 := func(name string, port *string) {
    if len(*port) > 0 {
      nServers++
      go func() {
        d, err := daemon.Daemon1(name, network, pre+*port, *requests, *sessions)
        if err == nil {
          if *unix { remove <- *port }
          err = d.Run()
        }
        errors <- err
      }()
    }
  }
  d1(FINGER, finger)

  // wait for completion or a signal
  wait: for nServers > 0 {
    select {
    case err := <- errors: // server is done
      if err != nil { println(err.String()) }
      nServers--
      
    case s := <- signal.Incoming: // signal: terminate only on some
      switch s.(signal.UnixSignal) {
      case syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT,
           syscall.SIGKILL, syscall.SIGTERM:
        break wait
      default:
        println("ignoring "+s.String())
      }
    }
  }
  
  // remove unix domain sockets if any
  if *unix {
    for {
      if path, ok := <- remove; ok {
        println("removing "+path)
        os.Remove(path)
      } else
        break
    }
  }
}
*/
