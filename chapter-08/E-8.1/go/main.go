package main

type rowType []int
type costType []rowType

// MinCostPathDP function implement dp
func MinCostPathDP(cm costType) int {
	rowNum, colNum := len(cm), len(cm[0])

	dp := make([][]int, rowNum)
	for i := range dp {
		dp[i] = make([]int, colNum)
	}

	// initialize the first row
	sum := 0
	for ci := 0; ci < colNum; ci++ {
		sum += cm[0][ci]
		dp[0][ci] = sum
	}

	// initialize the first column
	sum = 0
	for ri := 0; ri < rowNum; ri++ {
		sum += cm[ri][0]
		dp[ri][0] = sum
	}

	// fill up the dp
	for ri := 1; ri < rowNum; ri++ {
		for ci := 1; ci < colNum; ci++ {
			leftPath := dp[ri][ci-1]
			upPath := dp[ri-1][ci]

			// can only from cell(i-1, j) or cell(i, j-1) to cell(i, j)
			// obviously, choose the minmum cost path
			if leftPath < upPath {
				dp[ri][ci] = leftPath + cm[ri][ci]
			} else {
				dp[ri][ci] = upPath + cm[ri][ci]
			}
		}
	}

	return dp[rowNum-1][colNum-1]
}

// MinCostPathRecursionMain function is the main entry for recursive implementation
func MinCostPathRecursionMain(cm costType) int {
	rowNum, colNum := len(cm), len(cm[0])
	return MinCostPathRecursion(cm, rowNum-1, colNum-1)
}

// MinCostPathRecursion function implement recursive solution
// f(i, j) represents the minmum cost from cell(0, 0) to cell(i, j)
func MinCostPathRecursion(cm costType, i, j int) int {
	if i == 0 && j == 0 {
		return cm[0][0]
	}
	if i == 0 {
		return cm[0][j] + MinCostPathRecursion(cm, 0, j-1)
	}
	if j == 0 {
		return cm[i][0] + MinCostPathRecursion(cm, i-1, 0)
	}

	left := MinCostPathRecursion(cm, i, j-1)
	up := MinCostPathRecursion(cm, i-1, j)

	ret := 0
	if left < up {
		ret = left + cm[i][j]
	} else {
		ret = up + cm[i][j]
	}

	return ret
}

// MinCostPathRecursionMemoryMain function is the main entry for MinCostPathRecursionMemory function
func MinCostPathRecursionMemoryMain(cm costType) int {
	rowNum, colNum := len(cm), len(cm[0])

	memory := make([][]int, rowNum)
	for i := range memory {
		memory[i] = make([]int, colNum)
	}

	for ri := 0; ri < rowNum; ri++ {
		for ci := 0; ci < colNum; ci++ {
			memory[ri][ci] = -1
		}
	}

	return MinCostPathRecursionMemory(cm, rowNum-1, colNum-1, memory)
}

// MinCostPathRecursionMemory function like MinCostPathRecursion,
// but add memory to get rid of duplicate function call.
func MinCostPathRecursionMemory(cm costType, i, j int, mem [][]int) int {
	memVal := mem[i][j]
	if memVal >= 0 {
		return memVal
	}

	now := cm[i][j]
	if i == 0 && j == 0 {
		mem[i][j] = now
	} else if i == 0 {
		mem[i][j] = now + MinCostPathRecursion(cm, 0, j-1)
	} else if j == 0 {
		mem[i][j] = now + MinCostPathRecursion(cm, i-1, 0)
	} else {
		left := MinCostPathRecursion(cm, i, j-1)
		up := MinCostPathRecursion(cm, i-1, j)

		ret := 0
		if left < up {
			ret = left + cm[i][j]
		} else {
			ret = up + cm[i][j]
		}

		mem[i][j] = ret
	}

	return mem[i][j]
}

func main() {

}
