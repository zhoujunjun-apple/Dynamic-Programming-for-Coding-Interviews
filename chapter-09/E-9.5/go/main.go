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

func main() {

}
