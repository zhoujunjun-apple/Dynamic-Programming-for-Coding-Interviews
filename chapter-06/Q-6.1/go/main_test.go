package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	ts := []struct {
		n   int
		exp int32
	}{
		{n: 1, exp: 1},
		{n: 2, exp: 1},
		{n: 3, exp: 2},
		{n: 4, exp: 3},
		{n: 5, exp: 5},
		{n: 6, exp: 8},
		{n: 7, exp: 13},
		{n: 8, exp: 21},
		{n: 9, exp: 34},
		{n: 10, exp: 55},
		{n: 11, exp: 89},
		{n: 12, exp: 144},
		{n: 13, exp: 233},
		{n: 14, exp: 377},
		{n: 15, exp: 610},
		{n: 16, exp: 987},
	}
	for _, tt := range ts {
		got := fib(tt.n)
		if got != tt.exp {
			t.Errorf("n=%d expected=%d, got=%d\n", tt.n, tt.exp, got)
		}
	}
}
