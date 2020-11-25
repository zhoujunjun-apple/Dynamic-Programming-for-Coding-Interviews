package main

import (
	"math/rand"
	"testing"
)

var testCases = []struct {
	arr []int
	max int
}{
	{
		arr: []int{-2, -3, 4, -1, -2, 1, 5, -3},
		max: 7,
	},
	{
		arr: []int{1, 2},
		max: 3,
	},
	{
		arr: []int{1, 2, 3, -2},
		max: 6,
	},
	{
		arr: []int{-1, 2, -4, 3},
		max: 3,
	},
	{
		arr: []int{-1, 2, -1, 3},
		max: 4,
	},
	{
		arr: []int{1, -1, 1, 1, 0, 1},
		max: 3,
	},
	{
		arr: []int{-2, 3, 1, -1, 0, -1},
		max: 4,
	},
	{
		arr: []int{-1, -2, -3},
		max: -1,
	},
	{
		arr: []int{1, -2, 3, 1, 4, -6, -1, 0, 2, 4, 6},
		max: 13,
	},
	{
		arr: []int{0, 1, -1, 3, 4, 2, -4, 6, -3, 0, 1, -3},
		max: 11,
	},
}

func TestBruteForce(t *testing.T) {
	for i, tc := range testCases {
		got := bruteForce(tc.arr)
		if got != tc.max {
			t.Errorf("index=%d, expected=%d, got=%d\n", i, tc.max, got)
		}
	}
}

func TestNativeDP(t *testing.T) {
	for i, tc := range testCases {
		got := nativeDP(tc.arr)
		if got != tc.max {
			t.Errorf("index=%d, expected=%d, got=%d\n", i, tc.max, got)
		}
	}
}

func BenchmarkBruteForce(b *testing.B) {
	tcLen := len(testCases)
	for i := 0; i < b.N; i++ {
		ri := rand.Intn(tcLen)
		bruteForce(testCases[ri].arr)
	}
}

func BenchmarkNativeDP(b *testing.B) {
	tcLen := len(testCases)
	for i := 0; i < b.N; i++ {
		ri := rand.Intn(tcLen)
		nativeDP(testCases[ri].arr)
	}
}
