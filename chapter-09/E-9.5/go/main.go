package main

// Recursion function return the lenght of LCS of s1 and s2
func Recursion(s1, s2 string) int {
	s1Len, s2Len := len(s1), len(s2)

	if s1Len == 0 || s2Len == 0 {
		return 0
	}
	if s1Len == 1 {
		for i := 0; i < s2Len; i++ {
			if s1 == s2[i:i+1] {
				return 1
			}
		}
		return 0
	}
	if s2Len == 1 {
		for i := 0; i < s1Len; i++ {
			if s2 == s1[i:i+1] {
				return 1
			}
		}
		return 0
	}

	// the tail elements of s1 and s2 are same, they must be part of the LCS
	if s1[s1Len-1] == s2[s2Len-1] {
		return Recursion(s1[:s1Len-1], s2[:s2Len-1]) + 1
	}

	// choose the longer one
	s1BackRet := Recursion(s1[:s1Len-1], s2)
	s2BackRet := Recursion(s1, s2[:s2Len-1])
	if s1BackRet > s2BackRet {
		return s1BackRet
	}
	return s2BackRet
}

// NativeDP function implement the dp
func NativeDP(s1, s2 string) int {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len == 0 || s2Len == 0 {
		return 0
	}

	// dp[i][j] represents the length of LCS between s1[0,i] and s2[0,j]
	dp := make([][]int, s1Len)
	for i := range dp {
		dp[i] = make([]int, s2Len)
	}

	// fill up the first column
	for ri := 0; ri < s1Len; ri++ {
		if s2[0] == s1[ri] {
			for i := ri; i < s1Len; i++ {
				dp[i][0] = 1
			}
			break
		}
	}

	// fill up the first row
	for ci := 0; ci < s2Len; ci++ {
		if s1[0] == s2[ci] {
			for i := ci; i < s2Len; i++ {
				dp[0][i] = 1
			}
			break
		}
	}

	// same logic with Recursion function
	for ri := 1; ri < s1Len; ri++ {
		for ci := 1; ci < s2Len; ci++ {
			if s1[ri] == s2[ci] {
				dp[ri][ci] = dp[ri-1][ci-1] + 1
			} else {
				up := dp[ri-1][ci]
				left := dp[ri][ci-1]

				if up > left {
					dp[ri][ci] = up
				} else {
					dp[ri][ci] = left
				}
			}
		}
	}

	return dp[s1Len-1][s2Len-1]
}

func main() {

}
