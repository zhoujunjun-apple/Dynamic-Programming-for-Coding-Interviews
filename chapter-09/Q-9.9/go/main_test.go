package main

import "testing"

var gTestCases = []struct {
	arr []int
	exp int
}{
	{
		arr: []int{},
		exp: 0,
	},
	{
		arr: []int{1, 1, 1, 1, 1},
		exp: 1,
	},
	{
		arr: []int{3, 2, 1},
		exp: 1,
	},
	{
		arr: []int{4, 3, 3, 3, 2, 2, 1, 1, 1},
		exp: 1,
	},
	{
		arr: []int{3, 2, 1, 2, 1},
		exp: 2,
	},
	{
		arr: []int{1, 1, 5, 5, 4, 3, 2, 1, 1},
		exp: 2,
	},
	{
		arr: []int{1, 4, 1, 3, 1, 2, 1, 0},
		exp: 2,
	},
	{
		arr: []int{1, 2, 3, 1, 2, 3},
		exp: 3,
	},
	{
		arr: []int{3, 2, 6, 4, 9},
		exp: 3,
	},
	{
		arr: []int{1, 2, 3, 0, 1, 2},
		exp: 3,
	},
	{
		arr: []int{1, 2, 3, 3, 2, 1},
		exp: 3,
	},
	{
		arr: []int{1, 1, 2, 2, 3, 3, 4, 4},
		exp: 4,
	},
	{
		arr: []int{1, 3, 5, 7, 2, 4, 6, 7},
		exp: 5,
	},
}

func TestRecursion(t *testing.T) {
	for i, tc := range gTestCases {
		got := RecursionMain(tc.arr)
		if got != tc.exp {
			t.Errorf("index=%d, arr=%v, expected=%d, got=%d\n", i, tc.arr, tc.exp, got)
		}
	}
}

func TestNativeDP(t *testing.T) {
	for i, tc := range gTestCases {
		got := NativeDP(tc.arr)
		if got != tc.exp {
			t.Errorf("index=%d, arr=%v, expected=%d, got=%d\n", i, tc.arr, tc.exp, got)
		}
	}
}
