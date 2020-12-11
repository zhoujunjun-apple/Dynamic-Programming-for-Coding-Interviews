package main

import "testing"

var gTestCases = []struct {
	arr []int // the input integer array
	exp int   // the expected longest length of LBS
}{
	{
		arr: []int{},
		exp: 0,
	},
	{
		arr: []int{2},
		exp: 1,
	},
	{
		arr: []int{4, 4, 4},
		exp: 1,
	},
	{
		arr: []int{1, 2},
		exp: 2,
	},
	{
		arr: []int{1, 2, 2, 2},
		exp: 2,
	},
	{
		arr: []int{1, 3, 2},
		exp: 3,
	},
	{
		arr: []int{1, 2, 2, 3, 3, 3},
		exp: 3,
	},
	{
		arr: []int{1, 2, 2, 0, 0, 1},
		exp: 3,
	},
	{
		arr: []int{1, 3, 5, 4, 3, 5, 3},
		exp: 5,
	},
	{
		arr: []int{1, 2, 3, 2, 1},
		exp: 5,
	},
	{
		arr: []int{1, 2, 3, 2, 3, 4, 2},
		exp: 5,
	},
	{
		arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		exp: 10,
	},
}

func TestNativeDP(t *testing.T) {
	for i, tc := range gTestCases {
		got := NativeDP(tc.arr)
		if got != tc.exp {
			t.Errorf("index=%d, arr=%v, expected=%d, got=%d\n", i, tc.arr, tc.exp, got)
		}
	}
}
