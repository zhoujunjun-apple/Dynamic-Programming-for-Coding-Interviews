package main

import "testing"

var gTestCases = []struct {
	array []int    // input array
	sum   int      // the expected positive sum
	exp   bool     // the expected bool result
	set   []subset // the expected subset results
}{
	{
		array: []int{1, 2, 3, 3, 2, 1},
		sum:   11,
		exp:   true,
		set: []subset{
			{3: 2, 2: 2, 1: 1},
		},
	},
	{
		array: []int{3, 2, 7, 1},
		sum:   6,
		exp:   true,
		set: []subset{
			{3: 1, 2: 1, 1: 1},
		},
	},
	{
		array: []int{3, 2, 7, 1},
		sum:   4,
		exp:   true,
		set: []subset{
			{1: 1, 3: 1},
		},
	},
	{
		array: []int{3, 2, 7, 1},
		sum:   10,
		exp:   true,
		set: []subset{
			{3: 1, 7: 1},
			{2: 1, 1: 1, 7: 1},
		},
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
		set: []subset{
			{5: 1, 3: 1, 2: 1},
			{4: 1, 6: 1},
			{3: 1, 1: 1, 6: 1},
			{4: 1, 5: 1, 1: 1},
			{4: 1, 2: 1, 3: 1, 1: 1},
		},
	},
	{
		array: []int{1, 2, 1, 2, 1, 2, 3, 4, 5},
		sum:   17,
		exp:   true,
		set: []subset{
			{5: 1, 4: 1, 3: 1, 2: 2, 1: 1},
			{5: 1, 4: 1, 3: 1, 2: 1, 1: 3},
			{5: 1, 3: 1, 2: 3, 1: 3},
		},
	},
	{
		array: []int{1, 2, 1, 2, 1, 2, 3, 4, 5},
		sum:   18,
		exp:   true,
		set: []subset{
			{5: 1, 4: 1, 3: 1, 2: 3},
			{5: 1, 4: 1, 3: 1, 2: 2, 1: 2},
			{5: 1, 4: 1, 2: 3, 1: 3},
		},
	},
	{
		array: []int{1, 3, 2, 4, 5, 7},
		sum:   15,
		exp:   true,
		set: []subset{
			{7: 1, 5: 1, 3: 1},
			{7: 1, 5: 1, 2: 1, 1: 1},
			{7: 1, 4: 1, 3: 1, 1: 1},
			{5: 1, 4: 1, 3: 1, 2: 1, 1: 1},
		},
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
	for i, tc := range gTestCases {
		gotBool, gotSet := Recursion(tc.array, tc.sum)
		if gotBool != tc.exp {
			t.Errorf("index=%d, expected=%t, got=%t\n", i, tc.exp, gotBool)
			continue
		}
		if gotBool == true {
			var found bool = false
			for _, expSet := range tc.set {
				if len(gotSet) != len(expSet) {
					continue
				}

				this := true
				for ek, ev := range expSet {
					if gv, exist := gotSet[ek]; !exist {
						this = false
						break
					} else if gv != ev {
						this = false
						break
					}
				}
				found = this

				if found {
					break
				}
			}
			if !found {
				t.Errorf("index=%d, unexpected got=%v\n", i, gotSet)
			}
		}
	}
}
