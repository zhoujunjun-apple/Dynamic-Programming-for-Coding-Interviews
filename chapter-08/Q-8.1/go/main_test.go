package main

import "testing"

var ts = []struct {
	costs costMatrix
	min   costType
}{
	{
		costs: costMatrix{
			costRowType{1, 3, 5, 8},
			costRowType{4, 2, 1, 7},
			costRowType{4, 3, 2, 3},
		},
		min: 7,
	},
	{
		costs: costMatrix{
			costRowType{1, 1, 1, 2, 3},
		},
		min: 8,
	},
	{
		costs: costMatrix{
			costRowType{1}, {2}, {3}, {4},
		},
		min: 10,
	},

	{
		costs: costMatrix{
			costRowType{1, 2, 3, 4, 5, 6},
			costRowType{1, 3, 4, 8, 1, 2},
			costRowType{2, 1, 1, 2, 1, 8},
			costRowType{3, 2, 1, 1, 3, 2},
		},
		min: 9,
	},
	{
		costs: costMatrix{
			costRowType{4, 3, 1, 2, 4, 5, 2},
			costRowType{2, 1, 3, 3, 2, 2, 1},
			costRowType{3, 1, 2, 1, 2, 2, 1},
			costRowType{1, 2, 1, 1, 1, 1, 2},
			costRowType{1, 1, 2, 2, 1, 3, 4},
		},
		min: 14,
	},
}

func TestMinPathCostDP(t *testing.T) {
	for i, tt := range ts {
		got := minPathCostDP(tt.costs)
		if got != tt.min {
			t.Errorf("index %d: expected=%d, got=%d", i, tt.min, got)
		}
	}
}
