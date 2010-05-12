package ws_daemon

import (
  "io"
  "websocket"
  "os"
  "strings"
)

// create an echo daemon.
// banner, if any, is the opening message.
func Echod(network, laddr string, requests, sessions int,
    banner ... string) (*Server, os.Error) {
  factory := func(websocket.Conn) Session {
    return func(out io.Writer, in io.Reader) {
      if len(banner) > 0 {
        io.WriteString(out, strings.Join(banner, " ")+"\n")
      }
      io.Copy(out, in)
    }
  }
  
  return New(network, laddr, factory, requests, sessions)
}
