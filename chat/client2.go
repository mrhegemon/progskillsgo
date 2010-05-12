package ws_daemon

import(
	"fmt"
	"io"
	"websocket"
	"os"
)

// connect reader/writer to network.
// trc can be nil to suppress trace.
func Client(network, raddr string, wrt io.Writer, rd io.Reader) os.Error {
  c, err := websocket.Dial(network, "", raddr)
  if err == nil {
    done := make(chan bool)
    var nW, nR int
    copy := func(pre string, n *int, out io.Writer, in io.Reader) {
      Copy(pre, n, out, in)
      done <- true
    }
    go copy("< ", &nW, wrt, c)
    go copy("> ", &nR, c, rd)
    <- done
    return c.Close()
  }
  return err
}

// taken from io.Copy; trace added.
func Copy(prefix string, n *int, dst io.Writer, src io.Reader) {
  buf := make([]byte, 32*1024)
  for {
    if nr, _ := src.Read(buf); nr > 0 {
      if nw, _ := dst.Write(buf[0:nr]); nw > 0 {
        *n += nw
        continue
      }
    }
    return
  }
}
