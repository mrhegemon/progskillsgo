package chat

import("sync"; "strconv")

type PeerTable struct {
	table map[int]Peer
	sync.Mutex
}

func(this *PeerTable) sendMessageToAll(msg string, pId int) {
	message := strconv.Itoa(pId) + ":  " + msg
	for id, peer := range table {
		if id != pId {
			peer.sendMessage(message)
		}
	}
}

func(this *PeerTable) removePeer(pId int) {
	table[pId] = nil, false
}

func(this *PeerTable) newPeer(c *websocket.Conn) *Peer {
	peer := new(Peer)
	peer.id = len(table)
	peer.table = this
	peer.conn = c
	
	return peer
}

func NewPeerTable() *PeerTable {
	tab := new(PeerTable)
	tab.table = make(map[int]Peer)
	
	return tab
}
	
