/*
Authors: William Broza, Tym Lipari
Matrix Testing program

tests add and mult matrices

usage:
	matrix
*/
package main

import ("fmt"; "matrix"; "os")
//import . "strconv"

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

	//TEST 1
	fmt.Println("add test 1")
	m1.BuildTestMatrix(10, 7, 24.68, 13.579)
	m2.BuildTestMatrix(10, 7, 90.36, 179.2)
	
	m1.Print()
	println("\n")
	m2.Print()
	println("\n")

	rm1, err1 := matrix.Add(m1, m2)
	if err1 != nil { fmt.Fprintln(os.Stderr, err1.String()) }
	perr1 := rm1.Print()
	if perr1 != nil { fmt.Fprintln(os.Stderr, perr1.String()) }

	//TEST 2
	fmt.Println("mult test 1")
	m3.BuildTestMatrix(4, 7, 24.068, 13.579)
	m4.BuildTestMatrix(7, 11, 24.868, 13.579)

	m3.Print()
	println("\n")
	m4.Print()
	println("\n")

	rm2, err2 := matrix.Mult(m3, m4)
	if err2 != nil { fmt.Fprintln(os.Stderr, err2.String()) }
	perr2 := rm2.Print()
	if perr2 != nil { fmt.Fprintln(os.Stderr, perr2.String()) }

	//TEST 3
	fmt.Println("add test 2")
	m5.BuildTestMatrix(5, 30, 24.65, 5.579)
	m6.BuildTestMatrix(5, 30, 14.68, 153.54)
	
	m5.Print()
	println("\n")
	m6.Print()
	println("\n")

	rm3, err3 := matrix.Add(m5, m6)
	if err3 != nil { fmt.Fprintln(os.Stderr, err3.String()) }
	perr3 := rm3.Print()
	if perr3 != nil { fmt.Fprintln(os.Stderr, perr3.String()) }

	//TEST 4
	fmt.Println("mult test 2")
	m7.BuildTestMatrix(2, 100, 234.68, 4390.02)
	m8.BuildTestMatrix(100, 17, 4.68, 13579.04)

	m7.Print()
	println("\n")
	m8.Print()
	println("\n")

	rm4, err4 := matrix.Mult(m7, m8)
	if err4 != nil { fmt.Fprintln(os.Stderr, err4.String()) }
	perr4 := rm4.Print()
	if perr4 != nil { fmt.Fprintln(os.Stderr, perr4.String()) }

	//TEST 5
	fmt.Println("return add error")
	m9.BuildTestMatrix(5, 31, 24.65, 5.579)
	m10.BuildTestMatrix(5, 30, 14.68, 153.54)

	m9.Print()
	println("\n")
	m10.Print()
	println("\n")

	rm5, err5 := matrix.Add(m9, m10)
	if err5 != nil { fmt.Fprintln(os.Stderr, err5.String()) }
	perr5 := rm5.Print()
	if perr5 != nil { fmt.Fprintln(os.Stderr, err5.String()) }

	//TEST 6
	fmt.Println("return mult error")
	m11.BuildTestMatrix(2, 10, 234.68, 8790.20)
	m12.BuildTestMatrix(100, 17, 4.68, 3579.450)
	
	m11.Print()
	println("\n")
	m12.Print()
	println("\n")

	rm6, err6 := matrix.Mult(m11, m12)
	if err6 != nil { fmt.Fprintln(os.Stderr, err6.String()) }
	perr6 := rm6.Print()
	if perr6 != nil { fmt.Fprintln(os.Stderr, perr6.String()) }
}


