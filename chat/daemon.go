package ws_daemon

import (
  "bufio"
  "exec"
  "io"
  "websocket"
  "os"
  "strings"
)

// create a daemon to process all input.
func Daemon(daemon, network, laddr string, requests, sessions int) (*Server, os.Error) {
  factory := func(websocket.Conn) Session {
    return func(out io.Writer, in io.Reader) {
      Execute(strings.Fields(daemon), out, in)
    }
  }
  
  return New(network, laddr, factory, requests, sessions)
}

// create a daemon using the first input line, if any, as arguments.
func Daemon1(daemon, network, laddr string, requests, sessions int) (*Server, os.Error) {
  factory := func(websocket.Conn) Session {
    return func(out io.Writer, in io.Reader) {
      rdr := bufio.NewReader(in)
      if line, err := rdr.ReadString('\n'); err != nil {
        io.WriteString(out, err.String()+"\n")
      } else {
        Execute(strings.Fields(daemon+" "+line), out, rdr);
      }
    }
  }

  return New(network, laddr, factory, requests, sessions)
}

// execute argv[0] as a process.
func Execute(argv []string, out io.Writer, in io.Reader) {
  if proc, err := exec.Run(argv[0], argv, os.Environ(), "",
      exec.Pipe, exec.Pipe, exec.MergeWithStdout); err != nil {
    io.WriteString(out, err.String()+"\n")
  } else {
    go io.Copy(proc.Stdin, in)
    go io.Copy(out, proc.Stdout)
    if err := proc.Close(); err != nil {
      io.WriteString(out, err.String()+"\n")
    }
  }
}
