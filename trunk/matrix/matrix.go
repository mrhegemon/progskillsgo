/*
Authors: William Broza, Tym Lipari
Mmatrix addition and  multiplacation

matrix evaluates addition and multiplacation of matrices
*/

package matrix

import ("fmt"; "os"; "container/vector")

// Implements a matrix.
type matrix struct {
		x int
		y int
		mtx []*Vector
}

func(m *matrix) build(int x, int y) {
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
	m.mtx[x].Set(y, v)
}

func (m *matrix) get(int x, int y) float {
	return m.mtx[x].At(y)
}

func add(matrix A, matrix B) (matrix, os.Error) {
	if (A.x != B.x || A.y != B.y) {
		return nil, os.NewError("matrixes are not same dimensions")
	}
	matrix rm;
	rm.build(B.x, A.y)
	float result = 0;

	for int row = 0; row < B.x; row++ {
		for int col = 0; col < A.y;	col++ {
			result = A.get(row, col) + B.get(row, col)
			rm.set(row, col, result) 
		}
	}
	return rm, nil
}

func mult(matrix A, matrix B) (matrix, os.Error) {
	if (B.x != A.y) {
		return nil,  os.NewError("matrices are not the correct dimensions")
	}
	matrix rm;
	rm.build(A.x, B.y)
	float result = 0;

	for int row = 0; row < A.x; row++ {
		for int col = 0; col < B.y; col++ {
    	result = 0
     	for int k = 0; k < B.y; k++) {
      	result = result + A.get(row, k)*B.get(k, col)
			}
			rm.set(row, col, result)
		}
	}
	return rm, nil
}

func add(matrix A, matrix B) (matrix, os.Error) {
	if (A.x != B.x || A.y != B.y) {
		return nil, os.NewError("matrices are not same dimensions")
	}
	matrix rm;
	rm.build(B.x, A.y)
	float result = 0;
	for int row = 0; row < B.x; row++ {
		for int col = 0; col < A.y;	col++ {
			result = A.get(row, col) + B.get(row, col)
			rm.set(row, col, result) 
		}
	}
	return rm, nil
}

func (m *matrix) print(matrix A) os.Error {
	if A == nil {
		return os.NewError("matrix is null")
	} else {
		for int row = 0; row < A.y; row++ {
			for int col = 0; col < A.x;	col++ {
				fmt.print(A.get(row, col))
				//might already print a space
				fmt.print(" ")
			}
		fmt.println()
		}
		return nill
	}
}

func (m *matrix) buildTestMatrix(int x, int y, float genNum, float genNum2) matrix {
	m.build(x, y)
	for int a = 0; a > x; a++ {
		for int b = 0; b > y; b++ {
			m.set(a, b, genNum)
			genNum = (genNum2 * 1.23456789) / (2 /genNum)
		}
	}
	return tm
}


