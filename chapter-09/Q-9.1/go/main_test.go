package main

import "testing"

var testCases = []struct {
	x int
	y int // target point (x, y)
	n int // expected paths number
}{
	{x: 0, y: 0, n: 1},
	{x: 1, y: 1, n: 2},
	{x: 2, y: 2, n: 6},
	{x: 3, y: 4, n: 35},
	{x: 5, y: 8, n: 1287},
	{x: 4, y: 7, n: 330},
	{x: 10, y: 10, n: 184756},
}

func TestRecursion(t *testing.T) {
	for _, tc := range testCases {
		got := Recursion(tc.x, tc.y)
		if got != tc.n {
			t.Errorf("x=%d, y=%d, expected=%d, got=%d\n", tc.x, tc.y, tc.n, got)
		}
	}
}

func TestNativeDP(t *testing.T) {
	for _, tc := range testCases {
		got, _ := NativeDP(tc.x, tc.y)
		if got != tc.n {
			t.Errorf("x=%d, y=%d, expected=%d, got=%d\n", tc.x, tc.y, tc.n, got)
		}
	}
}
