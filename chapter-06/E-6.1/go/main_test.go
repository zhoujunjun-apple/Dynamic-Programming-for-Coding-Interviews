package main

import (
	"testing"
)

func Test_BalanceSumSubstringPartDP(t *testing.T) {
	ts := []struct {
		input string
		exput int
	}{
		{
			input: "142124",
			exput: 6,
		},
		{
			input: "9430723",
			exput: 4,
		},
		{
			input: "11203009",
			exput: 6,
		},
		{
			input: "123",
			exput: 0,
		},
		{
			input: "1422",
			exput: 2,
		},
		{
			input: "1322",
			exput: 4,
		},
		{
			input: "12345678123456789",
			exput: 16,
		},
	}

	for _, tc := range ts {
		got, err := BalanceSumSubstringPartDP(tc.input)
		if err != nil {
			t.Errorf(err.Error())
		}
		if got != tc.exput {
			t.Errorf("input=%s, expected=%d, got=%d", tc.input, tc.exput, got)
		}
	}
}
