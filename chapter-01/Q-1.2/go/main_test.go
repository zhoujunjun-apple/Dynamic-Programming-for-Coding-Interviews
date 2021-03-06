package main

import (
	"testing"
)

func TestAddPrevious(t *testing.T) {
	tt := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{0},
			output: []int{0},
		},
		{
			input:  []int{0, 1},
			output: []int{0, 1},
		},
		{
			input:  []int{0, 1, 2},
			output: []int{0, 1, 3},
		},
		{
			input:  []int{0, 0, 1, -1, 2, -2, 3, -3},
			output: []int{0, 0, 1, 0, 2, 0, 3, 0},
		},
		{
			input:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			output: []int{0, 1, 3, 6, 10, 15, 21, 28},
		},
	}
	for _, ts := range tt {
		inpdata := ts.input
		expdata := ts.output
		AddPrevious(inpdata)
		CheckResult(t, inpdata, expdata)
	}
}

func TestAddPreviousIterative(t *testing.T) {
	tt := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{0},
			output: []int{0},
		},
		{
			input:  []int{0, 1},
			output: []int{0, 1},
		},
		{
			input:  []int{0, 1, 2},
			output: []int{0, 1, 3},
		},
		{
			input:  []int{0, 0, 1, -1, 2, -2, 3, -3},
			output: []int{0, 0, 1, 0, 2, 0, 3, 0},
		},
		{
			input:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			output: []int{0, 1, 3, 6, 10, 15, 21, 28},
		},
	}
	for _, ts := range tt {
		inpdata := ts.input
		expdata := ts.output
		AddPreviousIterative(inpdata)
		CheckResult(t, inpdata, expdata)
	}
}

func CheckResult(t *testing.T, outdata []int, expdata []int) {
	if len(outdata) != len(expdata) {
		t.Fatalf("length unmatch. expected=%d, got=%d\n", len(expdata), len(outdata))
	}
	for idx, gotVal := range outdata {
		expVal := expdata[idx]
		if gotVal != expVal {
			t.Errorf("expected=%d, got=%d from expected output: %v\n", expVal, gotVal, expdata)
		}
	}
}

func BenchmarkAddPreviousIterative(b *testing.B) {
	t := []int{1, 3, 24, 5, 3, 6, 9, 8, 7, 9, 10, 5, 10, 33, 23, 45, 64, 89, 10}
	for idx := 0; idx < b.N; idx++ {
		AddPreviousIterative(t)
	}
}

func BenchmarkAddPrevious(b *testing.B) {
	t := []int{1, 3, 24, 5, 3, 6, 9, 8, 7, 9, 10, 5, 10, 33, 23, 45, 64, 89, 10}
	for idx := 0; idx < b.N; idx++ {
		AddPrevious(t)
	}
}
