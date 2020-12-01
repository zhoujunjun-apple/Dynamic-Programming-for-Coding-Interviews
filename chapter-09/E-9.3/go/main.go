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

// SimpleDP function like NativeDP, but only consider 2-dim dp.
// Since a[0,ri] and b[0, ki] are the only two free variables, c[0, ri+ki] isn't.
// It' much easy to understand compared with NativeDP, in which 3-dim dp matrix is used.
func SimpleDP(a, b, c string) bool {
	aLen, bLen, cLen := len(a), len(b), len(c)
	if aLen+bLen != cLen {
		return false
	}

	// dp[ri][ki]表示a的前ri个字符 和 b的前ki个字符是否能够按照interleaving的方式构成 c的前ri+ki个字符
	dp := make([][]bool, aLen+1)
	for i := range dp {
		dp[i] = make([]bool, bLen+1)
	}

	// ignore b, only consider a and c
	for ri := 1; ri <= aLen; ri++ {
		if a[:ri] == c[:ri] {
			dp[ri][0] = true
		} else {
			dp[ri][0] = false
		}
	}

	// ignore a, only consider b and c
	for ki := 1; ki <= bLen; ki++ {
		if b[:ki] == c[:ki] {
			dp[0][ki] = true
		} else {
			dp[0][ki] = false
		}
	}

	// fill up the dp matrix
	for ri := 1; ri <= aLen; ri++ {
		for ki := 1; ki <= bLen; ki++ {
			// the index of a, b and c, respectively
			ai, bi, ci := ri-1, ki-1, ri+ki-1
			if a[ai] == c[ci] && b[bi] == c[ci] {
				dp[ri][ki] = dp[ri-1][ki] || dp[ri][ki-1]
			} else if a[ai] == c[ci] && b[bi] != c[ci] {
				dp[ri][ki] = dp[ri-1][ki]
			} else if a[ai] != c[ci] && b[bi] == c[ci] {
				dp[ri][ki] = dp[ri][ki-1]
			} else {
				dp[ri][ki] = false
			}
		}
	}

	return dp[aLen][bLen]
}

func main() {

}
