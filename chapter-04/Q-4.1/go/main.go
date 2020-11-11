package main

// wayNumRecursive function use recursion
func wayNumRecursive(rowIdx, colIdx int) int {
	if rowIdx < 0 || colIdx < 0 {
		return 0
	}
	if rowIdx == 0 || colIdx == 0 {
		return 1
	}
	return wayNumRecursive(rowIdx, colIdx-1) + wayNumRecursive(rowIdx-1, colIdx)
}

// SIZE is the size of memory array
const SIZE int = 21

var memory [SIZE][SIZE]int32

// initMemory function initialize the memory cell to zero value
func initMemory() {
	for ri := 0; ri < SIZE; ri++ {
		for ci := 0; ci < SIZE; ci++ {
			memory[ri][ci] = 0
		}
	}
}

// wayNumRecursiveMemo function combines recursion and memory
func wayNumRecursiveMemo(rowIdx, colIdx int) int32 {
	hit := memory[rowIdx][colIdx]
	if hit != 0 {
		return hit
	}
	if rowIdx == 0 || colIdx == 0 {
		memory[rowIdx][colIdx] = 1
		return 1
	}
	calc := wayNumRecursiveMemo(rowIdx-1, colIdx) + wayNumRecursiveMemo(rowIdx, colIdx-1)
	memory[rowIdx][colIdx] = calc
	return calc
}

// wayNumDP function use dynamic programming
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
