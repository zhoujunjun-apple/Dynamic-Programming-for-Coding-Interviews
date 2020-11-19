package main

import (
	"testing"
)

var ts = []struct {
	total int
	ways  int
}{
	{total: 0, ways: 0},
	{total: 1, ways: 0},
	{total: 2, ways: 0},
	{total: 3, ways: 1},
	{total: 4, ways: 0},
	{total: 5, ways: 1},
	{total: 6, ways: 1},
	{total: 7, ways: 0},
	{total: 8, ways: 2},
	{total: 9, ways: 1},
	{total: 10, ways: 2},
	{total: 11, ways: 3},
	{total: 12, ways: 1},
	{total: 13, ways: 5},
	{total: 14, ways: 4},
	{total: 15, ways: 4},
	{total: 16, ways: 9},
	{total: 17, ways: 5},
	{total: 18, ways: 11},
	{total: 19, ways: 14},
	{total: 20, ways: 11},
	{total: 21, ways: 23},
	{total: 22, ways: 20},
	{total: 23, ways: 27},
	{total: 24, ways: 41},
	{total: 25, ways: 35},
}

func TestReachWaysRecursive(t *testing.T) {
	for i, tt := range ts {
		got := ReachWaysRecursive(tt.total)
		if got != tt.ways {
			t.Errorf("index %d: input=%d, expected=%d, got=%d", i, tt.total, tt.ways, got)
		}
	}
}

func TestReachWaysDP(t *testing.T) {
	for i, tt := range ts {
		got := ReachWaysDP(tt.total)
		if got != tt.ways {
			t.Errorf("index %d: input=%d, expected=%d, got=%d", i, tt.total, tt.ways, got)
		}
	}
}
