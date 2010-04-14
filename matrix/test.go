
/*
Authors: William Broza, Tym Lipari
Matrix Testing program

tests add and mult matrices

usage:
	test
*/
package matrix

import ("fmt"; "os"; "strconv"; "flag"; "container/vector")

//program main
func main(){
	matrix m1, m2, m3, m4, m5, m6, m7, m8, m9, m10, m11, m12

	fmt.print("add test 1")
	m1.buildTestMatrix(10, 7, 24.68, 13.579)
	m2.buildTestMatrix(10, 7, 90.36, 179.2)

	matrix rm1 = add(m1, m2)
	rm1.print()

	fmt.print("mult test 1")
	m3.buildTestMatrix(4, 7, 24.068, 13.579)
	m4.buildTestMatrix(7, 11, 24.868, 13.579)

	matrix rm2 = mult(m3, m4)
	rm2.print()

	fmt.print("add test 2")
	m5.buildTestMatrix(5, 30, 24.65, 5.579)
	m6.buildTestMatrix(5, 30, 14.68, 153.54)

	matrix rm3 = add(m5, m6)
	rm3.print()

	fmt.print("mult test 2")
	m7.buildTestMatrix(2, 100, 234.68, 1390.0)
	m8.buildTestMatrix(100, 17, 4.68, 13579.0)

	matrix rm4 = mult(m7, m8)
	rm4.print()

	fmt.print("return add error")
	m9.buildTestMatrix(5, 31, 24.65, 5.579)
	m10.buildTestMatrix(5, 30, 14.68, 153.54)

	matrix rm5 = add(m9, m10)
	rm5.print()

	fmt.print("return mult error")
	m11.buildTestMatrix(2, 10, 234.68, 1390.0)
	m12.buildTestMatrix(100, 17, 4.68, 13579.0)

	matrix rm6 = mult(m11, m12)
	rm6.print()
}


