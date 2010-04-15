package list

struct node {
	val {}interface
	prev, next *node
}

func(this *node) linkToFront(n *node) {
	n.linkToFront(this.prev)
	n.linkToBack(this)
	this.prev = n
}

func(this *node) linkToBack(n *node) {
	n.linkToFront(this)
	n.linkToBack(this.next)
	this.next = n
}

func(this *node) setValue(value {}interface) {}interface {
	old := this.val
	this.val = value
	return old
}

func newNode(val {}interface) *node{
	newNode := new(node)
	newNode.setValue(val)
}

