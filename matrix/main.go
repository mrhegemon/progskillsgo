
/*
Authors: William Broza, Tym Lipari
Matrix Testing program

tests add and mult matrices

usage:
	matrix
*/
package main

import ("fmt"; "matrix")

//program main
func main() {
	//test matrices
	m1 := new(matrix.Matrix) 
	m2 := new(matrix.Matrix) 
	m3 := new(matrix.Matrix) 
	m4 := new(matrix.Matrix) 
	m5 := new(matrix.Matrix) 
	m6 := new(matrix.Matrix) 
	m7 := new(matrix.Matrix) 
	m8 := new(matrix.Matrix) 
	m9 := new(matrix.Matrix) 
	m10 := new(matrix.Matrix) 
	m11 := new(matrix.Matrix)
	m12 := new(matrix.Matrix)
	//return matrices
	rm1 := new(matrix.Matrix) 
	rm2 := new(matrix.Matrix) 
	rm3 := new(matrix.Matrix) 
	rm4 := new(matrix.Matrix) 
	rm5 := new(matrix.Matrix) 
	rm6 := new(matrix.Matrix) 

	//TEST 1
	fmt.Print("add test 1")
	m1.BuildTestMatrix(10, 7, 24.68, 13.579)
	m2.BuildTestMatrix(10, 7, 90.36, 179.2)

	rm1, _ = matrix.Add(m1, m2)
	rm1.Print()

	//TEST 2
	fmt.Print("mult test 1")
	m3.BuildTestMatrix(4, 7, 24.068, 13.579)
	m4.BuildTestMatrix(7, 11, 24.868, 13.579)

	rm2, _ = matrix.Mult(m3, m4)
	rm2.Print()

	//TEST 3
	fmt.Print("add test 2")
	m5.BuildTestMatrix(5, 30, 24.65, 5.579)
	m6.BuildTestMatrix(5, 30, 14.68, 153.54)

	rm3, _ = matrix.Add(m5, m6)
	rm3.Print()

	//TEST 4
	fmt.Print("mult test 2")
	m7.BuildTestMatrix(2, 100, 234.68, 1390.0)
	m8.BuildTestMatrix(100, 17, 4.68, 13579.0)

	rm4, _ = matrix.Mult(m7, m8)
	rm4.Print()

	//TEST 5
	fmt.Print("return add error")
	m9.BuildTestMatrix(5, 31, 24.65, 5.579)
	m10.BuildTestMatrix(5, 30, 14.68, 153.54)

	rm5, _ = matrix.Add(m9, m10)
	rm5.Print()

	//TEST 6
	fmt.Print("return mult error")
	m11.BuildTestMatrix(2, 10, 234.68, 1390.0)
	m12.BuildTestMatrix(100, 17, 4.68, 13579.0)

	rm6, _ = matrix.Mult(m11, m12)
	rm6.Print()
}


