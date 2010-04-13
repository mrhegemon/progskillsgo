
/*
Authors: William Broza, Tym Lipari
Matrix addition and  multiplacation

atrix evaluates addition and multiplacation of matrices
*/

package atrix

import ("fmt"; "os"; "container/vector")

// Implements a atrix.
type atrix struct {
		x int
		y int
		mtx []*Vector
}

func(m *atrix) build(int x, int y) {
	m.x = x
	m.y = y
	m.mtx = make([] *Vector, x)

	for _, a := range m.mtx {
		a = new(Vector)
		for b := 0; b < y; b++{
			a.Push(0);
	}
}

func (m *atrix) set(int x, int y, float v) {
	m.mtx[x].Set(y, v)
}

func (m *atrix) get(int x, int y) float {
	return m.mtx[x].At(y)
}

func add(atrix A, atrix B) (atrix, os.Error) {
	if (A.x != B.x || A.y != B.y) {
		return nil, os.NewError("Matrixes not same dimensions")
	}

	atrix rm;
	rm.build(B.x, A.y)
	float result = 0;

	for int row = 0; row < B.x; row++ {
		for int col = 0; col < A.y;	col++ {
			result = A.At(row, col) + B.At(row, col)
			rm.set(row, col, result) 
		}
	}
	return rm, _
}

func mult(atrix A, atrix B) (atrix, os.Error) {
	if (A.x != B.x || A.y != B.y) {
		os.Error err
		return _, err
	}

	atrix rm;
	rm.build(B.x, A.y)
	float result = 0;

	for int row = 0; row < B.x; row++ {
		for int col = 0; col < A.y; col++ {
    	result = 0
     	for int k = 0; k < B.y; k++) {
      	result = result + A.get(row, k)*B.get(k, col)
			}
			rm.set(row, col, result)
		}
	}
	return rm, _
}

