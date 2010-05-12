package chat

import("websocket"; "bufio"; "sync")

type Peer struct {
	msgs chan string
	table PeerTable
	id int
	conn *websocket.Conn
	sync.Mutex
}

func(this *Peer) sendMessage(message string) {
	this.msgs <- message
}

func(this *Peer) printLoop() {
	for {
		message := <- msgs
		_, err := this.conn.Write([]byte(message))
		if err != nil {
			this.table.removePeer(this.id)
			conn.Close()
			break
		}
	}
}

func(this *Peer) readLoop() {
	for {
		reader := bufio.NewReader(conn)
		message, err := reader.ReadString([]byte("\n"))
		if err != nil {
			println(err.String())
			conn.Close()
			table.removePeer(id)
			break
		} else {
			table.sendMessageToAll(string(message), id)
		}
	}
}