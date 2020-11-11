package main

// Matrix represent a 2-d matrix
type Matrix struct {
	cell00 int32 // top left value
	cell01 int32 // top right value
	cell10 int32 // bottom left value
	cell11 int32 // bottom right value
}

// fib function return the n-th Fibonacci number
func fib(n int) int32 {
	ret := fibMatrixCalc(n)
	return ret.cell01
}

// fibMatrixCalc function implement M^n
func fibMatrixCalc(n int) Matrix {
	if n == 1 {
		return Matrix{
			cell00: 1,
			cell01: 1,
			cell10: 1,
			cell11: 0,
		}
	}

	half := fibMatrixCalc(n / 2)
	power2 := multiMatrix(half, half)
	if n%2 == 0 {
		return power2
	}
	base := Matrix{
		cell00: 1,
		cell01: 1,
		cell10: 1,
		cell11: 0,
	}
	return multiMatrix(base, power2)
}

// multiMatrix function implement 2-dim leftMatrix * rightMatrix
func multiMatrix(left, right Matrix) Matrix {
	return Matrix{
		cell00: left.cell00*right.cell00 + left.cell01*right.cell10,
		cell01: left.cell00*right.cell01 + left.cell01*right.cell11,
		cell10: left.cell10*right.cell00 + left.cell11*right.cell10,
		cell11: left.cell10*right.cell01 + left.cell11*right.cell11,
	}
}

func main() {

}
