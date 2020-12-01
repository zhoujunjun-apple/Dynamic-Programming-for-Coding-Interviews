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

func TestCheckInterleaving(t *testing.T) {
	for idx, tc := range gTestCases {
		got := checkInterleaving(tc.a, tc.b, tc.c)
		if got != tc.exp {
			t.Errorf("index=%d, a:%s, b:%s, c:%s, expected=%t, got=%t\n", idx, tc.a, tc.b, tc.c, tc.exp, got)
		}
	}
}
