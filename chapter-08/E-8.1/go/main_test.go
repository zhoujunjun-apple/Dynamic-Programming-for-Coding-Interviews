package main

import (
	"math/rand"
	"testing"
)

var ts = []struct {
	costs costType
	min   int
}{
	{
		costs: costType{
			rowType{1, 1, 1, 2, 3},
		},
		min: 8,
	},
	{
		costs: costType{
			rowType{1}, {2}, {3}, {4},
		},
		min: 10,
	},
	{
		costs: costType{
			rowType{1, 3, 5, 8},
			rowType{4, 2, 1, 7},
			rowType{4, 3, 2, 3},
		},
		min: 12,
	},
	{
		costs: costType{
			rowType{1, 2, 3, 4, 5, 6},
			rowType{1, 3, 4, 8, 1, 2},
			rowType{2, 1, 1, 2, 1, 8},
			rowType{3, 2, 1, 1, 3, 2},
		},
		min: 13,
	},
}

func TestMinCostPathDP(t *testing.T) {
	for i, tt := range ts {
		got := MinCostPathDP(tt.costs)
		if got != tt.min {
			t.Errorf("index=%d: expected=%d, got=%d", i, tt.min, got)
		}
	}
}

func TestMinCostPathRecursion(t *testing.T) {
	for i, tt := range ts {
		got := MinCostPathRecursionMain(tt.costs)
		if got != tt.min {
			t.Errorf("index=%d: expected=%d, got=%d", i, tt.min, got)
		}
	}
}

func BenchmarkMinCostPathDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rd := rand.Intn(len(ts))
		MinCostPathDP(ts[rd].costs)
	}
}

func BenchmarkMinCostPathRecursionMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rd := rand.Intn(len(ts))
		MinCostPathRecursionMain(ts[rd].costs)
	}
}

func BenchmarkMinCostPathRecursionMemoryMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rd := rand.Intn(len(ts))
		MinCostPathRecursionMemoryMain(ts[rd].costs) // the memory initialization may cost more time when the cost matrix is small
	}
}
