/*
Authors: William Broza, Tym Lipari
Matrix Testing program

tests add and mult matrices

usage:
	matrix
*/
package list

import "os"
//import "strconv"

var (
	RANGE_ERROR os.Error
)

func init() {
	RANGE_ERROR = os.NewError("LinkedList:  Index out of range.")
}

type node struct {
	val        interface{}
	prev, next *node
}

func (this *node) getPrev() *node {
	return this.prev
}
func (this *node) getNext() *node {
	return this.next
}
func (this *node) setPrev(n *node) {
	this.prev = n
}
func (this *node) setNext(n *node) {
	this.next = n
}

func (this *node) setValue(value interface{}) interface{} {
	old := this.val
	this.val = value
	return old
}

func (this *node) getValue() interface{} { return this.val }

func newNode(val interface{}) *node {
	newNode := new(node)
	newNode.setValue(val)
	return newNode
}

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

type LinkedList struct {
	head, tail *node
	length     int
}

func (this *LinkedList) Len() int { return this.length }

func (this *LinkedList) Init() { this.length = 0; this.head = nil; this.tail = nil }

func (this *LinkedList) Front() interface{} { return this.head.getValue() }

func (this *LinkedList) Back() interface{} { return this.tail.getValue() }

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

func (this *LinkedList) At(index int) (interface{}, os.Error) {
	foundNode, err := this.getNode(index)

	if foundNode != nil {
		return foundNode.getValue(), err
	}
	return nil, err
}

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
