package list

struct node {
	val {}interface
	prev, next *node
}

func(this *node) linkToFront(n *node) {
	if n != nil {
		n.linkToFront(this.prev)
		n.linkToBack(this)
	}
	this.prev = n
}

func(this *node) linkToBack(n *node) {
	if n != nil {
		n.linkToFront(this)
		n.linkToBack(this.next)
	}
	this.next = n
}

func(this *node) setValue(value {}interface) {}interface {
	old := this.val
	this.val = value
	return old
}

func(this *node) getValue() {}interface { return this.val }

func newNode(val {}interface) *node{
	newNode := new(node)
	newNode.setValue(val)
}


struct LinkedList {
	head, tail *node
	length int
}

func (this *LinkedList) Len() int { return this.length }
func (this *LinkedList) Init() { this.length = 0; head, tail = nil }
func (this *LinkedList) Front() {}interface { return this.head.getValue() }
func (this *LinkedList) Back() {}interface { return this.tail.getValue() }
func (this *LinkedList) PushFront(val {}interface) {
	nnode := newNode(val)

	if this.head == nil {
		this.tail = nnode
	} else {
		nnode.linkToBack(this.head)
	}
	this.head = nnode
	this.length++
}
func (this *LinkedList) PushBack(val {}interface) {
	nnode := newNode(val)

	if this.head == nil {
		this.head = nnode
	} else {
		nnode.linkToFront(this.tail)
	}
	this.tail = nnode
	this.length++
}




