/*
Authors: William Broza, Tym Lipari
Matrix addition and  multiplacation

expr evaluates expressions simular to the unix command shell

usage:
	matrix [arg1] ... [argn]
*/

package main

import ("fmt"; "os"; "strconv"; "flag"; "container/vector")

//constants for interpreting
const( PLUS = "+"
       TIMES = "*"
     )

//program main
func main(){
//?
}

func add(matrix A, matrix B) {
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

