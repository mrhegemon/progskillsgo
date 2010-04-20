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
	var verbose bool = false
	setArgs := func() {
		for y := 1; y < len(os.Args); y++ {
			if os.Args[y] == "-v" {
				verbose = true
			}
		}
	}

	setArgs()

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
	test := ""
	scanner := func(val interface{}, index int) os.Error {
		test += (val.(string))
		if verbose {
			println("[" + Itoa(index) + "] = " + (val.(string)))
		}
		return nil
	}

	toTest.ApplyToAllFromFront(scanner)

	if test != "hello" {
		println("Test Failed (ApplyToFront - 1) test4 != \"hello\"")
	}

	//Test 5: (AT)
	value, err := toTest.At(0)
	if err == nil && value.(string) != "hello" {
		println("Test Failed (AT) toTest.At(0) != \"hello\"")
	}

	//Test 6: Add Another Item To Front:

	//clear the list
	toTest.Init()

	//Add some random strings. Leading numbers indicate
	//expected order
	toTest.PushFront("3World")
	toTest.PushBack("4Worldz")
	toTest.PushBack("5kWorldz")
	toTest.PushFront("2Hello")
	toTest.PushBack("6Helloz")
	toTest.PushFront("1Hey")

	if verbose {
		println("Size:  " + Itoa(toTest.Len()))
	}
	//reset the test variable
	test = ""
	toTest.ApplyToAllFromFront(scanner)

	if test != "1Hey2Hello3World4Worldz5kWorldz6Helloz" {
		println("Test Failed (ApplyToFront - several) test != \"1Hey2Hello3World4Worldz5kWorldz6Helloz\"")
	}


	//Test 7: Remove the head of the list
	toTest.Remove(0)

	if verbose {
		println("Size:  " + Itoa(toTest.Len()))
	}
	test = ""
	toTest.ApplyToAllFromBack(scanner)

	if test != "6Helloz5kWorldz4Worldz3World2Hello" {
		println("Test Failed (Remove - 1) test != \"6Helloz5kWorldz4Worldz3World2Hello\"")
	}

	//Test 8: Remove several arbitrary
	if _, err := toTest.Remove(3); err != nil {
		println(err.String())
	}
	if _, err := toTest.Remove(1); err != nil {
		println(err.String())
	}

	if verbose {
		println("Size:  " + Itoa(toTest.Len()))
	}
	test = ""
	toTest.ApplyToAllFromFront(scanner)

	if test != "2Hello4Worldz6Helloz" {
		println("Test Failed (Remove - several) test != \"2Hello4Worldz6Helloz\"")
	}
}
