package chat

import("http"; "websocket")

type Server struct {
	peertable PeerTable
	port string
}

func ChatServer(ws *websocket.Conn) {
	peertable.newPeer(ws)
}

func ListenAndServe() {
	http.HandleFunc("/chat", websocket.Handle(ChatServer))
	for {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			panic("ListenAndServe:   " + err.String())
		}
	}
}