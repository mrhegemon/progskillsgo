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

var( RANGE_ERROR os.Error)

func init(){
	RANGE_ERROR = os.NewError("LinkedList:  Index out of range.")
}

type node struct {
	val interface{}
	prev, next *node
}

func(this *node) getPrev() *node {
	return this.prev
}
func(this *node) getNext() *node {
	return this.next
}
func(this *node) setPrev(n *node) {
	this.prev = n
}
func(this *node) setNext(n *node) {
	this.next = n
}

func(this *node) setValue(value interface{}) interface{} {
	old := this.val
	this.val = value
	return old
}

func(this *node) getValue() interface{} { return this.val }

func newNode(val interface{}) *node{
	newNode := new(node)
	newNode.setValue(val)
	return newNode
}

func link(front, back *node) os.Error {
	if back.getPrev() != nil {
		back.getPrev().setNext(front)
	}
	back.setPrev(front)
	
	if front.getNext() != nil {
		front.getNext().setPrev(back)
	}
	front.setNext(back)
	
	return nil
}

//===========LinkedList================//

type LinkedList struct{
	head, tail *node
	length int
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
func(this *LinkedList) Remove(index int) (interface{}, os.Error) {
	if index < 0 || index >= this.length {
		return nil, RANGE_ERROR
	}

	//special cases (index = 0 or length-1)
	if index == 0 {
		link(nil, this.head.getNext())
		this.head = this.head.getNext()
	} else if index == this.length - 1 {
		link(this.tail.getPrev(), nil)
		this.tail = this.tail.getPrev()
	}

	tempNode := this.head
	for y := 1; y <= index; y++ {
		tempNode = tempNode.next
	}

	link(tempNode.getPrev(), tempNode.getNext())
	//link(tempNode.getNext(), tempNode.getPrev())

	this.length--

	return tempNode.getValue(), nil
}

func(this *LinkedList) At(index int) (interface{}, os.Error) {
	if index < 0 || index >= this.length {
		return nil, RANGE_ERROR 
	}
	
	if index < (this.length / 2) {
		tempNode := this.head
		for y:= 1; y < this.length; y++ {
			tempNode = tempNode.next
		}
		return tempNode.getValue(), nil
	} else {
		tempNode := this.tail
		for y:= this.length - 2; y > 0; y-- {
			tempNode = tempNode.prev
		}
		return tempNode.getValue(), nil
	}

	return nil, os.NewError("LinkedList:  Search Error")
}

func(this *LinkedList) ApplyToAllFromFront(action func(interface{}, int)os.Error) os.Error {
	tempNode := this.head
	for y:=0; y < this.length; y++ {
		err := action(tempNode.getValue(), y)
		if err != nil { return err }
		tempNode = tempNode.getNext()
	}
	return nil
}

func(this *LinkedList) ApplyToAllFromBack(action func(interface{}, int)os.Error) os.Error {
	tempNode := this.tail
	for y:=this.length-1; y >= 0; y-- {
		err := action(tempNode.getValue(), y)
		if err != nil { return err }
		tempNode = tempNode.getPrev()
	}
	return nil
}





