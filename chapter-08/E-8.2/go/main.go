package main

// TileWaysDP function implement dp. It is same with Fibonacci
func TileWaysDP(N int) int {
	if N < 3 {
		return N
	}

	zero, one, two := 0, 1, 2
	for i := 3; i <= N; i++ {
		zero = one
		one = two
		two = zero + one
	}
	return two
}

func main() {

}
