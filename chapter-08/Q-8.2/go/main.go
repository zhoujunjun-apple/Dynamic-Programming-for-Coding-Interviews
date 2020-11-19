package main

// TileWaysRecursion function use recursion
// a 3*2 plot have 3 different ways of placing tile, so 3*f(i-2)
// a 3*4 plot have 2 different ways of placing tile in middle of plot, so 2*f(i-4)
func TileWaysRecursion(N int) int {
	ret := 0
	if N%2 == 1 || N < 0 {
		ret = 0
	} else {
		if N == 0 {
			ret = 1 // special case
		} else if N == 2 {
			ret = 3
		} else if N == 4 {
			ret = 11
		} else {
			ret = 3*TileWaysRecursion(N-2) + 2*TileWaysRecursion(N-4)
		}
	}
	return ret
}

func main() {

}
