package main

// void type represents a empty type
type void struct{}

// result type represents all possible LCS with the key
type result map[string]void

// Recursion function return the  LCS of s1 and s2
func Recursion(s1, s2 string) string {
	s1Len, s2Len := len(s1), len(s2)

	if s1Len == 0 || s2Len == 0 {
		return ""
	}
	if s1Len == 1 {
		for i := 0; i < s2Len; i++ {
			if s1 == s2[i:i+1] {
				return s1
			}
		}
		return ""
	}
	if s2Len == 1 {
		for i := 0; i < s1Len; i++ {
			if s2 == s1[i:i+1] {
				return s2
			}
		}
		return ""
	}

	// the tail elements of s1 and s2 are same, they must be part of the LCS
	if s1[s1Len-1] == s2[s2Len-1] {
		backLcs := Recursion(s1[:s1Len-1], s2[:s2Len-1])
		return backLcs + s1[s1Len-1:s1Len]
	}

	// choose the longer one
	s1BackLcs := Recursion(s1[:s1Len-1], s2)
	s2BackLcs := Recursion(s1, s2[:s2Len-1])
	if len(s1BackLcs) > len(s2BackLcs) {
		return s1BackLcs
	}
	return s2BackLcs
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
