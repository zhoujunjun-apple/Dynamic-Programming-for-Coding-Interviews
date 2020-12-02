package main

// Recursion function checks if subset of arr has sum of 'sum'
func Recursion(arr []int, sum int) bool {
	arrLen := len(arr)

	if arrLen == 0 && sum > 0 {
		return false
	}
	if sum == 0 {
		return true
	}
	if sum < 0 {
		return false
	}

	// the last element of arr may be the part of 'sum' or not
	return Recursion(arr[:arrLen-1], sum) || Recursion(arr[:arrLen-1], sum-arr[arrLen-1])
}

func nativeDP(arr []int, sum int) bool {
	arrLen := len(arr)

	// dp[i][j] 表示arr[0, i]中是否有部分元素的和等于j
	dp := make([][]bool, arrLen)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	// fill up the first column of dp
	for ri := 0; ri < arrLen; ri++ {
		if arr[ri] == 0 {
			for i := ri; i < arrLen; i++ {
				dp[i][0] = true
			}
			break
		} else {
			dp[ri][0] = false
		}
	}

	// fill up the first row of dp
	for ci := 0; ci <= sum; ci++ {
		if arr[0] == ci {
			dp[0][ci] = true
		} else {
			dp[0][ci] = false
		}
	}

	for ri := 1; ri < arrLen; ri++ {
		for ci := 1; ci <= sum; ci++ {
			up := dp[ri-1][ci]
			sub := ci - arr[ri]

			// the rest sum if less than zero, arr[ri] cannot be part of ci
			if sub < 0 {
				dp[ri][ci] = up
			} else {
				dp[ri][ci] = up || dp[ri-1][sub]
			}
		}
	}

	return dp[arrLen-1][sum]
}

func main() {

}
