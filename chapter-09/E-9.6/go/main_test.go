package main

import "testing"

var gTestCases = []struct {
	s1  string
	s2  string
	exp result // all the possible LCS of s1 and s2
}{
	{
		s1:  "",
		s2:  "abc",
		exp: result{"": void{}},
	},
	{
		s1:  "ABCD",
		s2:  "ABED",
		exp: result{"ABD": void{}},
	},
	{
		s1: "abc123def",
		s2: "321cbafed",
		exp: result{
			"3d": void{},
			"3e": void{},
			"3f": void{},
			"2d": void{},
			"2e": void{},
			"2f": void{},
			"1d": void{},
			"1e": void{},
			"1f": void{},
			"cd": void{},
			"ce": void{},
			"cf": void{},
			"bd": void{},
			"be": void{},
			"bf": void{},
			"ad": void{},
			"ae": void{},
			"af": void{},
		},
	},
	{
		s1: "abc123abc",
		s2: "ac3",
		exp: result{
			"ac3": void{},
		},
	},
	{
		s1: "abcd1234abcd",
		s2: "24adabcd",
		exp: result{
			"24abcd": void{},
		},
	},
	{
		s1:  "abc",
		s2:  "123",
		exp: result{"": void{}},
	},
	{
		s1: "abc",
		s2: "cde",
		exp: result{
			"c": void{},
		},
	},
	{
		s1: "abcdef",
		s2: "abcdef",
		exp: result{
			"abcdef": void{},
		},
	},
	{
		s1: "123456",
		s2: "654321",
		exp: result{
			"1": void{},
			"2": void{},
			"3": void{},
			"4": void{},
			"5": void{},
			"6": void{},
		},
	},
	{
		s1: "aaabbbccc123123",
		s2: "abc2233",
		exp: result{
			"abc233": void{},
			"abd223": void{},
		},
	},
	{
		s1: "aabbcc112233def",
		s2: "ac13fb2de",
		exp: result{
			"ac13de": void{},
		},
	},
	{
		s1: "abcdefg123456",
		s2: "ace135bdf246",
		exp: result{
			"abdf246": void{},
			"ace1356": void{},
		},
	},
}

func TestRecursion(t *testing.T) {
	for i, tc := range gTestCases {
		got := Recursion(tc.s1, tc.s2)
		if _, exist := tc.exp[got]; !exist {
			t.Errorf("index=%d, s1=%s, s2=%s, unexpectly got=%s\n", i, tc.s1, tc.s2, got)
		}
	}
}

// func TestNativeDP(t *testing.T) {
// 	for i, tc := range gTestCases {
// 		got := NativeDP(tc.s1, tc.s2)
// 		// if got != tc.exp {
// 		// 	t.Errorf("index=%d, s1=%s, s2=%s, expected=%d, got=%d\n", i, tc.s1, tc.s2, tc.exp, got)
// 		// }
// 	}
// }
