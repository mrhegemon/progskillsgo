/*
Authors: William Broza, Tym Lipari
Matrix addition and  multiplacation

Matrix evaluates addition and multiplacation of matrices
*/

package matrix

import ("fmt"; "os"; "container/vector")

// Implements a Matrix.
type Matrix struct {
		mrow int
		mcol int
		mtx []*vector.Vector
}

func(m *Matrix) GetRow() int {
	return m.mrow
}

func(m *Matrix) GetCol() int {
	return m.mcol
}

func(m *Matrix) SetRow(x int) {
	m.mrow = x
}

func(m *Matrix) SetCol(y int) {
	m.mcol = y
}

func(m *Matrix) Build(x int, y int) {
	m.SetRow(x)
	m.SetCol(y)
	m.mtx = make([] *vector.Vector, x)

	for n, _ := range m.mtx {
		m.mtx[n] = new(vector.Vector)
		for b := 0; b < y; b++ {
			m.mtx[n].Push(0)
		}
	}
}

func(m *Matrix) Set(x int, y int, v float) {
	m.mtx[x].Set(y, v)
}

func(m *Matrix) Get(x int, y int) float {
	return m.mtx[x].At(y).(float)
}

func Add(A *Matrix, B *Matrix) (*Matrix, os.Error) {
	if A.GetRow() != B.GetRow() || A.GetCol() != B.GetCol() {

		//will not accept nil!! FIX
		return A, os.NewError("Matrixes are not same dimensions")
	}
	rm := new(Matrix)
	rm.Build(B.GetRow(), A.GetCol())
	var result float
	result = 0.0
	for row := 0; row < B.GetRow(); row++ {
		for col := 0; col < A.GetCol();	col++ {
			result = A.Get(row, col) + B.Get(row, col)
			rm.Set(row, col, result) 
		}
	}
	return rm, nil
}

func Mult(A *Matrix, B *Matrix) (*Matrix, os.Error) {
	if (B.GetRow() != A.GetCol()) {
		
		//will not accept nil!! FIX
		return A,  os.NewError("matrices are not the correct dimensions")
	}
	rm := new(Matrix)
	rm.Build(A.GetRow(), B.GetCol())
	var result float
	result = 0.0
	for row := 0; row < A.GetRow(); row++ {
		for col := 0; col < B.GetCol(); col++ {
    	result = 0.0
     	for k := 0; k < B.GetCol(); k++ {
      	result = result + A.Get(row, k)*B.Get(k, col)
			}
			rm.Set(row, col, result)
		}
	}
	return rm, nil
}

func(m *Matrix) Print() os.Error {
	if m == nil {
		return os.NewError("Matrix is null")
	} else {
		for col := 0; col < m.GetCol(); col++ {
			fmt.Print("| ")
			for row := 0; row < m.GetRow();	row++ {
				fmt.Print(m.Get(row, col))
				fmt.Print(" ")
			}			
				fmt.Print("|")
		fmt.Println()
		}
	}
	return nil
}

func (m *Matrix) BuildTestMatrix(x int, y int, genNum float, genNum2 float) *Matrix {
	rm := new(Matrix)
	rm.Build(x, y)
	for a := 0; a < x; a++ {
		for b := 0; b < y; b++ {
			rm.Set(a, b, genNum)
			genNum = (genNum2 * 1.23456789) / (2 /genNum)
		}
	}
	return rm
}


