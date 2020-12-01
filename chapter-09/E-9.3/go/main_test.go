package main

import "testing"

var gTestCases = []struct {
	a   string
	b   string
	c   string
	exp bool
}{
	{
		a:   "xyz",
		b:   "abcd",
		c:   "xabyczd",
		exp: true,
	},
	{
		a:   "xyz",
		b:   "abcd",
		c:   "xabyczdz",
		exp: false,
	},
	{
		a:   "xyz",
		b:   "abcd",
		c:   "abxczyd",
		exp: false,
	},
	{
		a:   "x",
		b:   "a",
		c:   "ax",
		exp: true,
	},
	{
		a:   "x",
		b:   "a",
		c:   "xa",
		exp: true,
	},
	{
		a:   "xy",
		b:   "ab",
		c:   "xyab",
		exp: true,
	},
	{
		a:   "xy",
		b:   "ab",
		c:   "abxy",
		exp: true,
	},
	{
		a:   "xy",
		b:   "ab",
		c:   "xayb",
		exp: true,
	},
	{
		a:   "xy",
		b:   "ab",
		c:   "yaxb",
		exp: false,
	},
	{
		a:   "abcdefg",
		b:   "1234567",
		c:   "123abc4567defg",
		exp: true,
	},
	{
		a:   "abcd",
		b:   "abcd",
		c:   "aabcbdcd",
		exp: true,
	},
	{
		a:   "abcd",
		b:   "dcba",
		c:   "adcbbacd",
		exp: true,
	},
	{
		a:   "abcd",
		b:   "abce",
		c:   "abcecdab",
		exp: false,
	},
	{
		a:   "abcde",
		b:   "xyz",
		c:   "gabcdexy",
		exp: false,
	},
	{
		a:   "abcde",
		b:   "xyz",
		c:   "abcxydfe",
		exp: false,
	},
}

func TestRecursion(t *testing.T) {
	for i, tc := range gTestCases {
		got := Recursion(&tc.a, &tc.b, &tc.c, len(tc.a)-1, len(tc.b)-1, len(tc.c)-1)
		if got != tc.exp {
			t.Errorf("%d-th: %s|%s|%s expected=%t, got=%t\n",
				i, tc.a, tc.b, tc.c, tc.exp, got)
		}
	}
}

func TestNativeDP(t *testing.T) {
	for i, tc := range gTestCases {
		got := NativeDP(tc.a, tc.b, tc.c)
		if got != tc.exp {
			t.Errorf("%d-th: %s|%s|%s expected=%t, got=%t\n",
				i, tc.a, tc.b, tc.c, tc.exp, got)
		}
	}
}

func TestSimpleDP(t *testing.T) {
	for i, tc := range gTestCases {
		got := SimpleDP(tc.a, tc.b, tc.c)
		if got != tc.exp {
			t.Errorf("%d-th: %s|%s|%s expected=%t, got=%t\n",
				i, tc.a, tc.b, tc.c, tc.exp, got)
		}
	}
}
