/*
Authors: William Broza, Tym Lipari
Matrix addition and  multiplacation

expr evaluates expressions simular to the unix command shell

usage:
	matrix [arg1] ... [argn]
*/

package matrix

import ("fmt"; "os"; "strconv"; "flag"; "container/vector")

//constants for interpreting
const( PLUS = "+"
       TIMES = "*"
     )

// Implements a Matrix.
type Matrix struct {
		x int
		y int
		mtx []*Vector
}

func(m *Matrix) build(int x, int y) {
	m.x = x
	m.y = y
	m.mtx = make([] *Vector, x)

	for _, a := range m.mtx {
		a = new(Vector)
		for b := 0; b < y; b++{
			a.Push(0);
	}
}

func (m *matrix) set(int x, int y, float v) {
	

}

func add(matrix A, matrix B) (matrix C) {
/*
double matsum[5][5];
int row, col;
	for(row=0; row<5; row++) {
		for(col=0; col<5; col++) {
		matsum[5][5]=(x[row][col]+y[row][col]);
	}
}
	return matsum[5][5];
*/

}

func mult(matrix A, matrix B) (matrix C) {
/*
 int i,j,k;
 for (i=1; i<=3; i++)
   for (j=1; j<=3; j++)
     {
     sum = 0;
     for (k=1; k<=3; k++)
       sum = sum + A[i][k]*B[k][j];
     }
     C[i][j] = sum;
*/

}

