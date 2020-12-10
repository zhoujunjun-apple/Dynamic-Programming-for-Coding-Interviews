package main

import "testing"

var gTestCases = []struct {
	s1  string
	s2  string
	exp int
}{
	{
		s1:  "",
		s2:  "abc",
		exp: 0,
	},
	{
		s1:  "ABCD",
		s2:  "ABED",
		exp: 3,
	},
	{
		s1:  "abc123def",
		s2:  "321cbafed",
		exp: 2,
	},
	{
		s1:  "abc123abc",
		s2:  "ac3",
		exp: 3,
	},
	{
		s1:  "abcd1234abcd",
		s2:  "24adabcd",
		exp: 6,
	},
	{
		s1:  "abc",
		s2:  "123",
		exp: 0,
	},
	{
		s1:  "abc",
		s2:  "cde",
		exp: 1,
	},
	{
		s1:  "abcdef",
		s2:  "abcdef",
		exp: 6,
	},
	{
		s1:  "123456",
		s2:  "654321",
		exp: 1,
	},
	{
		s1:  "aaabbbccc123123",
		s2:  "abc2233",
		exp: 6,
	},
	{
		s1:  "aabbcc112233def",
		s2:  "ac13fb2de",
		exp: 6,
	},
	{
		s1:  "abcdefg123456",
		s2:  "ace135bdf246",
		exp: 7,
	},
}

func TestRecursion(t *testing.T) {
	for i, tc := range gTestCases {
		got := Recursion(tc.s1, tc.s2)
		if got != tc.exp {
			t.Errorf("index=%d, s1=%s, s2=%s, expected=%d, got=%d\n", i, tc.s1, tc.s2, tc.exp, got)
		}
	}
}

func TestNativeDP(t *testing.T) {
	for i, tc := range gTestCases {
		got := NativeDP(tc.s1, tc.s2)
		if got != tc.exp {
			t.Errorf("index=%d, s1=%s, s2=%s, expected=%d, got=%d\n", i, tc.s1, tc.s2, tc.exp, got)
		}
	}
}
