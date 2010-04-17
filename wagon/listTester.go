package main

import "list"

func main() {
	toTest := new(list.LinkedList)
	
	//Test 1: Empty List
	if toTest.Len() != 0 {
		println("Test Failed (Empty List):   Len != 0")
	}
	
	//Test 2: Add An Item To Front
	toTest.PushFront("hello")
	if toTest.Len() != 1 {
		println("Test Failed (Add to Front Len != 1")
	}
	
	//Test 3: Verify Item (FRONT)
	item := toTest.Front().(string)
	if item != "hello" {
		println("Test Failed (Verify Item 1) item != \"hello\"")
	}
}
