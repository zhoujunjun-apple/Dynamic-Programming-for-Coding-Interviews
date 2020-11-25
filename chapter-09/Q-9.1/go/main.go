package main

import "fmt"

// Recursion function use recursive method
func Recursion(x, y int) int {
	if x < 0 || y < 0 {
		return 0
	}
	if x == 0 || y == 0 {
		return 1
	}
	return Recursion(x-1, y) + Recursion(x, y-1)
}

// NativeDP function use dp
func NativeDP(x, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, fmt.Errorf("x, y must not less than zero, but got x=%d, y=%d", x, y)
	}
	if x == 0 || y == 0 {
		return 1, nil
	}

	// note the dp has (x+1)*(y+1) dimension
	dp := make([][]int, x+1)
	for i := range dp {
		dp[i] = make([]int, y+1)
	}

	// only one path from (0,0) to (i,0)
	for i := 0; i <= x; i++ {
		dp[i][0] = 1
	}

	// only one path from (0,0) to (0,i)
	for i := 0; i <= y; i++ {
		dp[0][i] = 1
	}

	for xi := 1; xi <= x; xi++ {
		for yi := 1; yi <= y; yi++ {
			dp[xi][yi] = dp[xi-1][yi] + dp[xi][yi-1]
		}
	}

	return dp[x][y], nil
}

func main() {

}
