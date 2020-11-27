package main

// Recursion function implement backward recursion solution
// check if c[0:ci] is interleaving of a[0:ai] and b[0:bi]
func Recursion(a, b, c *string, ai, bi, ci int) bool {
	if ai+bi+1 != ci {
		return false
	}
	if ai < 0 {
		if bi == 0 {
			return (*b)[bi] == (*c)[ci]
		}
		return (*b)[:bi] == (*c)[:ci]
	}
	if bi < 0 {
		if ai == 0 {
			return (*a)[ai] == (*c)[ci]
		}
		return (*a)[:ai] == (*c)[:ci]
	}

	ret := false

	// check the last character
	aLast, bLast, cLast := (*a)[ai], (*b)[bi], (*c)[ci]
	if cLast == aLast && cLast != bLast {
		// the last character c[ci] comes from a[ai]
		ret = Recursion(a, b, c, ai-1, bi, ci-1)
	} else if cLast != aLast && cLast == bLast {
		// the last character c[ci] comes from b[bi]
		ret = Recursion(a, b, c, ai, bi-1, ci-1)
	} else if cLast == aLast && cLast == bLast {
		// the last character may comes from a[ai] or b[bi]
		ret = Recursion(a, b, c, ai-1, bi, ci-1) || Recursion(a, b, c, ai, bi-1, ci-1)
	}

	return ret
}

// NativeDP function use 3-dim dp
func NativeDP(a, b, c string) bool {
	aLen, bLen, cLen := len(a), len(b), len(c)
	if aLen+bLen != cLen {
		return false
	}
	if aLen == 0 {
		return b == c
	}
	if bLen == 0 {
		return a == c
	}

	// dp[ai][bi][ci] represents if c's first ci characters c[0, ci-1] is interleaving
	// from a's first ai characters a[0, ai-1] and b's first bi characters b[0, bi-1]
	dp := make([][][]bool, aLen+1)
	for ai := range dp {
		dp[ai] = make([][]bool, bLen+1)
		for bi := range dp[ai] {
			dp[ai][bi] = make([]bool, cLen+1)
		}
	}

	// only consider part b and part c of same length
	bcMin := getMin(bLen, cLen)
	for i := 1; i <= bcMin; i++ {
		dp[0][i][i] = b[:i] == c[:i]
	}

	// only consider part a and part c of same length
	acMin := getMin(aLen, cLen)
	for i := 1; i <= acMin; i++ {
		dp[i][0][i] = a[:i] == c[:i]
	}

	// note: ai, bi represents the first ai and bi characters of a and b, respectively
	for ai := 1; ai <= aLen; ai++ {
		for bi := 1; bi <= bLen; bi++ {
			ci := ai + bi

			ac, bc, cc := a[ai-1], b[bi-1], c[ci-1]
			if ac == cc && bc == cc {
				dp[ai][bi][ci] = dp[ai-1][bi][ci-1] || dp[ai][bi-1][ci-1]
			} else if ac == cc && bc != cc {
				dp[ai][bi][ci] = dp[ai-1][bi][ci-1]
			} else if ac != cc && bc == cc {
				dp[ai][bi][ci] = dp[ai][bi-1][ci-1]
			} else {
				dp[ai][bi][ci] = false
			}
		}
	}

	return dp[aLen][bLen][cLen]
}

func getMin(is ...int) int {
	min := is[0]
	for i := 1; i < len(is); i++ {
		if is[i] < min {
			min = is[i]
		}
	}

	return min
}

func main() {

}
