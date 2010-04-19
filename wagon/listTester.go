/*
Authors: William Broza, Tym Lipari
Matrix Testing program

tests add and mult matrices

usage:
	matrix
*/
package main

import "list"
import "os"
import . "strconv"

func main() {
	toTest := new(list.LinkedList)
	
	//Test 1: Empty List
	if toTest.Len() != 0 {
		println("Test Failed (Empty List): Len != 0")
	}
	
	//Test 2: Add An Item To Front
	toTest.PushFront("hello")
	if toTest.Len() != 1 {
		println("Test Failed (Add to Front Len != 1, Len =" + Itoa(toTest.Len()))
	}
	
	//Test 3: Verify Item (FRONT)
	item := toTest.Front().(string)
	if item != "hello" {
		println("Test Failed (Verify Item 1) item != \"hello\"")
	}
	
	//Test 4: Applying Actions (FROM FRONT)
	test4 := ""
	printer := func(val interface{}, index int) os.Error {
		//test4 += (val.(string))
		println("[" + Itoa(index) + "] = " + (val.(string)))
		return nil
	}
	
	toTest.ApplyToAllFromFront(printer)
	
	if test4 != "hello" {
		println("Test Failed (ApplyToFront) test4 != \"hello\"")
	}
	
	//Test 5: (AT)
	value, err := toTest.At(0)
	if err == nil && value.(string) != "hello" {
		println("Test Failed (AT) toTest.At(0) != \"hello\"")
	}

	toTest2 := new(list.LinkedList)

	//Test 6: Add Another Item To Front:
	toTest2.PushFront("2World")
	toTest2.PushFront("1Worldz")
	toTest2.PushBack("3kWorldz")

	println("Back:  " + toTest2.Back().(string))
	println("Size:  " + Itoa(toTest2.Len()))
	toTest2.ApplyToAllFromBack(printer)
}


