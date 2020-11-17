package main

import (
	"fmt"
	"strconv"
)

// BalanceSumSubstringPartDP function use dp to optimize 'sum' procedure
func BalanceSumSubstringPartDP(s string) (int, error) {
	slen := len(s)
	// convert char digits to int. eg: '123' -> [1, 2, 3]
	snum := make([]int, slen)
	for i, c := range s {
		if ic, err := strconv.Atoi(string(c)); err == nil {
			snum[i] = ic
		} else {
			return -1, fmt.Errorf("invalid char %c", c)
		}
	}

	// use 2-dim dp array to record history sum
	// initialize the diagonal value to s[i] since sum[i, i] = s[i]
	dp := make([][]int, slen)
	for i := 0; i < slen; i++ {
		dp[i] = make([]int, slen)
		dp[i][i] = snum[i]
	}

	ret := 0
	// fill the dp array value at up-right area
	for step := 1; step < slen; step++ {
		for left := 0; left < slen-step; left++ {
			right := left + step
			mid := left + int(step/2)

			sumLeftHalf := dp[left][mid]
			sumRightHalf := dp[mid+1][right]
			dp[left][right] = sumLeftHalf + sumRightHalf

			if step%2 == 1 && sumLeftHalf == sumRightHalf && ret < step+1 {
				ret = step + 1
			}
		}
	}
	return ret, nil
}

func main() {

}
