package main

import "testing"

var gTestCases = []struct {
	a   cell // starting cell
	b   cell // ending cell
	exp int  // expected minimum move steps
}{
	{
		a:   cell{x: 1, y: 1},
		b:   cell{x: 1, y: 1},
		exp: 0,
	},
	// single dimension: 1 offset
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 1, y: 2},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 3, y: 2},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 2, y: 1},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 2, y: 3},
		exp: 1,
	},
	// two dimensions: 1 offset for both dimension
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 1, y: 1},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 3, y: 3},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 1, y: 3},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 3, y: 1},
		exp: 1,
	},
	// two dimensions: 1 offset for one and 2 offset for another
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 0, y: 3},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 0, y: 1},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 4, y: 3},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 4, y: 1},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 1, y: 0},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 3, y: 0},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 1, y: 4},
		exp: 1,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 3, y: 4},
		exp: 1,
	},
	// at least two moves
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 0, y: 0},
		exp: 2,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 0, y: 4},
		exp: 2,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 4, y: 0},
		exp: 2,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 4, y: 4},
		exp: 2,
	},
	// others
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 5, y: 1},
		exp: 2,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 3, y: 6},
		exp: 3,
	},
	{
		a:   cell{x: 2, y: 2},
		b:   cell{x: 7, y: 7},
		exp: 4,
	},
	{
		a:   cell{x: 0, y: 1},
		b:   cell{x: 7, y: 1},
		exp: 4,
	},
	{
		a:   cell{x: 0, y: 2},
		b:   cell{x: 7, y: 2},
		exp: 4,
	},
	{
		a:   cell{x: 0, y: 3},
		b:   cell{x: 7, y: 3},
		exp: 4,
	},
	{
		a:   cell{x: 0, y: 4},
		b:   cell{x: 7, y: 4},
		exp: 4,
	},
	{
		a:   cell{x: 0, y: 5},
		b:   cell{x: 7, y: 5},
		exp: 4,
	},
	{
		a:   cell{x: 1, y: 0},
		b:   cell{x: 1, y: 7},
		exp: 4,
	},
	{
		a:   cell{x: 2, y: 0},
		b:   cell{x: 2, y: 7},
		exp: 4,
	},
	{
		a:   cell{x: 3, y: 0},
		b:   cell{x: 3, y: 7},
		exp: 4,
	},
	{
		a:   cell{x: 4, y: 0},
		b:   cell{x: 4, y: 7},
		exp: 4,
	},
	{
		a:   cell{x: 5, y: 0},
		b:   cell{x: 5, y: 7},
		exp: 4,
	},
	{
		a:   cell{x: 6, y: 0},
		b:   cell{x: 6, y: 7},
		exp: 4,
	},
	// diagonal cells
	{
		a:   cell{x: 0, y: 0},
		b:   cell{x: 7, y: 7},
		exp: 5,
	},
	{
		a:   cell{x: 0, y: 7},
		b:   cell{x: 7, y: 0},
		exp: 5,
	},
	// pair cells with one of same axis value
	{
		a:   cell{x: 0, y: 0},
		b:   cell{x: 0, y: 7}, // lower left cell to upper left cell
		exp: 4,
	},
	{
		a:   cell{x: 0, y: 0},
		b:   cell{x: 7, y: 0}, // lower left cell to lower right cell
		exp: 4,
	},
	{
		a:   cell{x: 0, y: 7},
		b:   cell{x: 7, y: 7}, // upper left cell to upper right cell
		exp: 4,
	},
	{
		a:   cell{x: 7, y: 0},
		b:   cell{x: 7, y: 7}, // lower right cell to upper right cell
		exp: 4,
	},
}

func TestRecursion(t *testing.T) {
	for idx, tc := range gTestCases {
		got := Recursion(tc.a, tc.b, 1)
		if got != tc.exp {
			t.Errorf("%d: (%d,%d) and (%d,%d): expected=%d, got=%d\n",
				idx, tc.a.x, tc.a.y, tc.b.x, tc.b.y, tc.exp, got)
		}
	}
}

func TestIsOneStepAway(t *testing.T) {
	testCases := []struct {
		a   cell // cell a
		b   cell // cell b
		exp bool // expected result of whether cell a and cell b is one step away
	}{
		{
			a:   cell{x: 1, y: 1},
			b:   cell{x: 1, y: 1},
			exp: false,
		},
		// single dimension: 1 offset
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 1, y: 2},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 3, y: 2},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 2, y: 1},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 2, y: 3},
			exp: true,
		},
		// two dimensions: 1 offset for both dimension
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 1, y: 1},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 3, y: 3},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 1, y: 3},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 3, y: 1},
			exp: true,
		},
		// two dimensions: 1 offset for one and 2 offset for another
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 0, y: 3},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 0, y: 1},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 4, y: 3},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 4, y: 1},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 1, y: 0},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 3, y: 0},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 1, y: 4},
			exp: true,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 3, y: 4},
			exp: true,
		},
		// out of one step away
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 0, y: 2},
			exp: false,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 4, y: 2},
			exp: false,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 2, y: 0},
			exp: false,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 2, y: 4},
			exp: false,
		},
		// out of one step away: 2 offset for both dimension
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 0, y: 0},
			exp: false,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 0, y: 4},
			exp: false,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 4, y: 0},
			exp: false,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 4, y: 4},
			exp: false,
		},
		// others
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 5, y: 1},
			exp: false,
		},
		{
			a:   cell{x: 2, y: 2},
			b:   cell{x: 3, y: 6},
			exp: false,
		},
	}

	for _, tc := range testCases {
		got := isOneStepAway(tc.a, tc.b)
		if got != tc.exp {
			t.Errorf("(%d,%d) and (%d,%d): expected=%t, got=%t\n",
				tc.a.x, tc.a.y, tc.b.x, tc.b.y, tc.exp, got)
		}
	}
}
