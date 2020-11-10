package main

func wayNumRecursive(rowIdx, colIdx int) int {
	if rowIdx < 0 || colIdx < 0 {
		return 0
	}
	if rowIdx == 0 || colIdx == 0 {
		return 1
	}
	return wayNumRecursive(rowIdx, colIdx-1) + wayNumRecursive(rowIdx-1, colIdx)
}

func wayNumDP(rows, cols int) int {
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	for ri := 0; ri < rows; ri++ {
		dp[ri][0] = 1
	}
	for ci := 0; ci < cols; ci++ {
		dp[0][ci] = 1
	}
	for ri := 1; ri < rows; ri++ {
		for ci := 1; ci < cols; ci++ {
			dp[ri][ci] = dp[ri-1][ci] + dp[ri][ci-1]
		}
	}
	return dp[rows-1][cols-1]
}

func main() {

}
