package chat

import("sync"; "strconv")//; "websocket")

type PeerTable struct {
	table map[int]*LoopPeer
	sync.Mutex
}

type LoopPeer interface {
	sendMessage(message string)
	printLoop()
	readLoop()
}

func(this *PeerTable) sendMessageToAll(msg string, pId int) {
	message := strconv.Itoa(pId) + ":  " + msg
	for id, peer := range this.table {
		if id != pId {
			peer.sendMessage(message)
		}
	}
}

func(this *PeerTable) removePeer(pId int) {
	this.table[pId] = nil, false
}

/*func(this *PeerTable) newPeer(c *websocket.Conn) *Peer {
	peer := new(Peer)
	peer.id = len(table)
	peer.table = this
	peer.conn = c
	peer.msgs = make(chan string, 20)
	
	this.table[peer.id] = peer
	
	go peer.printLoop()
	go peer.readLoop()
	
	return peer
}*/

func NewPeerTable() *PeerTable {
	tab := new(PeerTable)
	tab.table = make(map[int]*LoopPeer)
			
	return tab
}
	
