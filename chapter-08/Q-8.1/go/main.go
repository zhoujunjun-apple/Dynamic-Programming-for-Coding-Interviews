package main

type costType int
type costRowType []costType
type costMatrix []costRowType

// minPathCostDP function implement dp
func minPathCostDP(cm costMatrix) costType {
	rowNum, colNum := len(cm), len(cm[0])

	dp := make(costMatrix, rowNum)
	for i := range dp {
		dp[i] = make(costRowType, colNum)
	}

	dp[0][0] = cm[0][0]
	for ri := 1; ri < rowNum; ri++ {
		dp[ri][0] = dp[ri-1][0] + cm[ri][0]
	}
	for ci := 1; ci < colNum; ci++ {
		dp[0][ci] = dp[0][ci-1] + cm[0][ci]
	}

	for ri := 1; ri < rowNum; ri++ {
		for ci := 1; ci < colNum; ci++ {
			dp[ri][ci] = cm[ri][ci] + minThree(dp[ri-1][ci], dp[ri][ci-1], dp[ri-1][ci-1])
		}
	}

	return dp[rowNum-1][colNum-1]
}

func minThree(a, b, c costType) costType {
	ret := a
	if a < b {
		if a < c {
			ret = a
		} else {
			ret = c
		}
	} else {
		if b < c {
			ret = b
		} else {
			ret = c
		}
	}
	return ret
}

func main() {

}
