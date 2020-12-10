package main

import (
	"strings"
)

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

// cell type represent a element of dp matrix.
// the index pair (ri, ci) used as a direction of the backward routing when combing the LCS.
// the index pair (ri, ci) records the indexs of the 'previous' dp element.
// the 'previous' dp element depends on the updating logic of 'length' field of current dp element.
// eg: the dp[i][j]'s index pair (ri, ci) means that dp[i][j].length is calculated based on the dp[ri][ci].length.
type cell struct {
	length int // the length of current LCS
	ri     int // the row index of the 'previous' dp element
	ci     int // the column index of the 'previous' dp element
}

// NativeDP function implement the dp
func NativeDP(s1, s2 string) string {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len == 0 || s2Len == 0 {
		return ""
	}

	// dp[i][j].length represents the length of LCS between s1[0,i] and s2[0,j]
	// dp[i][j].ri represents the row index of the 'previous' dp element
	// dp[i][j].ci represents the column index of the 'previous' dp element
	dp := make([][]cell, s1Len)
	for i := range dp {
		dp[i] = make([]cell, s2Len)
	}

	// fill up the first column
	for ri := 0; ri < s1Len; ri++ {
		if s2[0] == s1[ri] {
			// the first possible character of LCS
			// make sure: dp[ri][ci].ri == ri and dp[ri][ci].ci == ci
			dp[ri][0] = cell{length: 1, ri: ri, ci: 0}
			for i := ri + 1; i < s1Len; i++ {
				dp[i][0] = cell{
					length: 1,
					ri:     i - 1,
					ci:     0,
				}
			}
			break
		} else {
			dp[ri][0] = cell{length: 0, ri: ri - 1, ci: 0}
		}
	}

	// fill up the first row
	for ci := 0; ci < s2Len; ci++ {
		if s1[0] == s2[ci] {
			// the first possible character of LCS
			// make sure: dp[ri][ci].ri == ri and dp[ri][ci].ci == ci
			dp[0][ci] = cell{length: 1, ri: 0, ci: ci}
			for i := ci + 1; i < s2Len; i++ {
				dp[0][i] = cell{
					length: 1,
					ri:     0,
					ci:     i - 1,
				}
			}
			break
		} else {
			dp[0][ci] = cell{length: 0, ri: 0, ci: ci - 1}
		}
	}

	// same logic with Recursion function
	for ri := 1; ri < s1Len; ri++ {
		for ci := 1; ci < s2Len; ci++ {
			if s1[ri] == s2[ci] {
				// found a character of LCS
				// the 'previous' element is from the adjacent up-left corner
				dp[ri][ci] = cell{
					length: dp[ri-1][ci-1].length + 1,
					ri:     ri - 1, // make sure (ri, ci) points to the up-left element
					ci:     ci - 1,
				}
			} else {
				up := dp[ri-1][ci]
				left := dp[ri][ci-1]

				if up.length > left.length {
					// the 'previous' element is from up
					dp[ri][ci] = cell{
						length: up.length,
						ri:     ri - 1, // make sure (ri, ci) points to the up element
						ci:     ci,
					}
				} else {
					// the 'previous' element is from left
					dp[ri][ci] = cell{
						length: left.length,
						ri:     ri,
						ci:     ci - 1, // make sure (ri, ci) points to the left element
					}
				}
			}
		}
	}

	s := "" // the LCS string
	ri, ci := s1Len-1, s2Len-1
	sLen := dp[ri][ci].length // the length of LCS

	// combing the LCS according to dp matrix only when LCS isn't empty
	if sLen > 0 {
		// lcs slice records every characters of LCS
		lcs := make([]string, sLen)
		for {
			now := dp[ri][ci]
			if now.ri == ri && now.ci == ci {
				// reach the first character of LCS
				lcs[0] = s1[ri : ri+1]
				break
			}

			if now.ri == ri-1 && now.ci == ci-1 {
				// found a character of LCS, save it at the tail of lcs slice
				lcs[now.length-1] = s1[ri : ri+1]
			}
			// jump back to the 'previous' dp element
			ri = now.ri
			ci = now.ci
		}

		s = strings.Join(lcs, "")
	}
	return s
}

func main() {

}
