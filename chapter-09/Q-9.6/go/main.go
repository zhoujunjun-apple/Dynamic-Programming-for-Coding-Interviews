package main

// checkInterleaving function checks if c is interleaving from a and b
// a and b shouldn't have same characters
func checkInterleaving(a, b, c string) bool {
	aLen, bLen, cLen := len(a), len(b), len(c)
	if aLen+bLen != cLen {
		return false
	}

	var ai, bi, ci int
	for ci = 0; ci < cLen; ci++ {
		cc := c[ci]
		if ai < aLen && cc == a[ai] {
			ai++
		} else if bi < bLen && cc == b[bi] {
			bi++
		} else {
			return false
		}
	}

	return true
}

func main() {

}
