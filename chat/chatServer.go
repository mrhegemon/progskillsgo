package chat



func ChatServer(ws *websocket.Conn) {
	peertable.newPeer(ws)
}