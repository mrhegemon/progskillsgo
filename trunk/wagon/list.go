/*
Authors: William Broza, Tym Lipari
Linked list and node library
*/

package list

import "os"
//import "strconv"

//range error variable
var (
	RANGE_ERROR os.Error
)

//initialize range error
func init() {
	RANGE_ERROR = os.NewError("LinkedList: Index out of range.")
}

//node structure
type node struct {
	val        interface{}
	prev, next *node
}

//node getPrev()
//returns: previous node
func (this *node) getPrev() *node {
	return this.prev
}

//node getNext()
//returns: next node
func (this *node) getNext() *node {
	return this.next
}

//node setPrev()
//sets the previous node
func (this *node) setPrev(n *node) {
	this.prev = n
}

//node setNext()
//sets the next node
func (this *node) setNext(n *node) {
	this.next = n
}

//node setValue()
//sets the node value
//returns: returns old value
func (this *node) setValue(value interface{}) interface{} {
	old := this.val
	this.val = value
	return old
}

//node getValue()
//gets the value
func (this *node) getValue() interface{} { return this.val }

//newNode()
//makes a new node
//returns: the made node
func newNode(val interface{}) *node {
	newNode := new(node)
	newNode.setValue(val)
	return newNode
}

//link()
//links two nodes in order
//returns: error
func link(front, back *node) os.Error {
	if back != nil {
		if back.getPrev() != nil {
			back.getPrev().setNext(front)
		}
		back.setPrev(front)
	}
	if front != nil {
		if front.getNext() != nil {
			front.getNext().setPrev(back)
		}
		front.setNext(back)
	}
	return nil
}

//===========LinkedList================//
//linked list structure
type LinkedList struct {
	head, tail *node
	length     int
}

//LinkedList Len()
//returns lenght of list
func (this *LinkedList) Len() int { return this.length }

//LinkedList Init()
//initializes the linked list
func (this *LinkedList) Init() { this.length = 0; this.head = nil; this.tail = nil }

//LinkedList Front()
//returns: value from the front of the list
func (this *LinkedList) Front() interface{} { return this.head.getValue() }

//LinkedList Back()
//returns: value from the back of the list
func (this *LinkedList) Back() interface{} { return this.tail.getValue() }

//LinkedList PushFront()
//adds value to the front of the list
func (this *LinkedList) PushFront(val interface{}) {
	nnode := newNode(val)
	if this.head == nil {
		this.tail = nnode
	} else {
		link(nnode, this.head)
	}
	this.head = nnode
	this.length++
}

//LinkedList PushBack()
//adds value to the back of the list
func (this *LinkedList) PushBack(val interface{}) {
	nnode := newNode(val)

	if this.head == nil {
		this.head = nnode
	} else {
		link(this.tail, nnode)
	}
	this.tail = nnode
	this.length++
}

//LinkedList Remove()
//removes value at the index
func (this *LinkedList) Remove(index int) (interface{}, os.Error) {
	if index < 0 || index >= this.length {
		return nil, RANGE_ERROR
	}
	tempNode := new(node)
	if tempNode.getPrev() != nil {
		tempNode.getPrev().setNext(tempNode.getNext())
	}
	if tempNode.getNext() != nil {
		tempNode.getNext().setPrev(tempNode.getPrev())
	}
	//special cases:
	if index == 0 {
		this.head = this.head.getNext()
	} else if index == this.Len()-1 {
		this.tail = this.tail.getPrev()
	}
	this.length--
	return tempNode.getValue(), nil
}

//LinkedList At()
//returns: value at the index
func (this *LinkedList) At(index int) (interface{}, os.Error) {
	foundNode, err := this.getNode(index)
	if foundNode != nil {
		return foundNode.getValue(), err
	}
	return nil, err
}

//LinkedList getNode()
//returns: node at the index
func (this *LinkedList) getNode(index int) (*node, os.Error) {
	if index < 0 || index >= this.length {
		return nil, RANGE_ERROR
	}
	if index < (this.length / 2) {
		tempNode := this.head
		for y := 1; y < index; y++ {
			tempNode = tempNode.getNext()
		}
		return tempNode, nil
	} else {
		tempNode := this.tail
		for y := this.length - 2; y > index; y-- {
			tempNode = tempNode.getPrev()
		}
		return tempNode, nil
	}

	return nil, os.NewError("LinkedList:  Search Error")
}

//LinkedList ApplyToAllFromFront()
//applys function in order from front
//returns: error
func (this *LinkedList) ApplyToAllFromFront(action func(interface{}, int) os.Error) os.Error {
	tempNode := this.head
	for y := 0; y < this.length; y++ {
		err := action(tempNode.getValue(), y)
		if err != nil {
			return err
		}
		tempNode = tempNode.getNext()
	}
	return nil
}

//LinkedList ApplyToAllFromBack()
//applys function in order from back
//returns: error
func (this *LinkedList) ApplyToAllFromBack(action func(interface{}, int) os.Error) os.Error {
	tempNode := this.tail
	for y := this.length - 1; y >= 0; y-- {
		err := action(tempNode.getValue(), y)
		if err != nil {
			return err
		}
		tempNode = tempNode.getPrev()
	}
	return nil
}
