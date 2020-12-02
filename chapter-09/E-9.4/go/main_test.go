package main

import "testing"

var gTestCases = []struct {
	array []int // input array
	sum   int   // the expected positive sum
	exp   bool  // the expected result
}{
	{
		array: []int{3, 2, 7, 1},
		sum:   6,
		exp:   true,
	},
	{
		array: []int{3, 2, 7, 1},
		sum:   4,
		exp:   true,
	},
	{
		array: []int{3, 2, 7, 1},
		sum:   10,
		exp:   true,
	},
	{
		array: []int{3, 2, 7, 1},
		sum:   14,
		exp:   false,
	},
	{
		array: []int{1, 2, 3, 4, 5, 6},
		sum:   10,
		exp:   true,
	},
	{
		array: []int{1, 2, 1, 2, 1, 2, 3, 4, 5},
		sum:   17,
		exp:   true,
	},
	{
		array: []int{1, 2, 1, 2, 1, 2, 3, 4, 5},
		sum:   18,
		exp:   true,
	},
	{
		array: []int{1, 3, 2, 4, 5, 7},
		sum:   15,
		exp:   true,
	},
	{
		array: []int{1, 4, 7, 9},
		sum:   6,
		exp:   false,
	},
	{
		array: []int{2, 5, 8},
		sum:   9,
		exp:   false,
	},
	{
		array: []int{2, 5, 8},
		sum:   11,
		exp:   false,
	},
	{
		array: []int{2, 5, 8},
		sum:   14,
		exp:   false,
	},
	{
		array: []int{1, 3, 5, 7, 9},
		sum:   2,
		exp:   false,
	},
}

func TestRecursion(t *testing.T) {
	for idx, tc := range gTestCases {
		got := Recursion(tc.array, tc.sum)
		if got != tc.exp {
			t.Errorf("index=%d, arr: %v, sum: %d, expected=%t, got=%t\n", idx, tc.array, tc.sum, tc.exp, got)
		}
	}
}

func TestNativeDP(t *testing.T) {
	for idx, tc := range gTestCases {
		got := nativeDP(tc.array, tc.sum)
		if got != tc.exp {
			t.Errorf("index=%d, arr: %v, sum: %d, expected=%t, got=%t\n", idx, tc.array, tc.sum, tc.exp, got)
		}
	}
}
