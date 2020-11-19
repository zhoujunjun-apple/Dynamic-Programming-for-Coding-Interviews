package main

// ReachWaysDP function use dp.
// each move can score 3, 5 or 10 points
// O(n) time and O(1) space
func ReachWaysDP(total int) int {
	if total < 3 { // total score can less than 3
		return 0
	}

	// dp[i] represents the reach ways of total=10*k+i
	dp := [10]int{} // the longest reliance distance is 10.
	dp[0] = 1       // score 10. special case
	dp[3] = 1       // score 3
	dp[5] = 1       // score 5

	for i := 6; i <= total; i++ {
		ni := i % 10
		part := dp[(i-3)%10] + dp[(i-5)%10]
		if i >= 10 {
			part += dp[(i-10)%10]
		}
		dp[ni] = part
	}

	return dp[total%10]
}

// ReachWaysRecursive function used as the baseline.
// test cases are generated from this function
func ReachWaysRecursive(total int) int {
	if total == 3 || total == 5 || total == 6 || total == 9 {
		return 1
	}
	if total == 8 || total == 10 {
		return 2
	}
	if total < 10 {
		return 0
	}

	ret := ReachWaysRecursive(total-3) + ReachWaysRecursive(total-5) + ReachWaysRecursive(total-10)
	return ret
}

func main() {

}
