package main

import "testing"

type tc struct {
	a   string
	b   string
	num int
}

var gTestCases = []tc{
	{
		a:   "ab",
		b:   "xy",
		num: 6,
	},
	{
		a:   "",
		b:   "xyz",
		num: 1,
	},
	{
		a:   "a",
		b:   "xy",
		num: 3,
	},
	{
		a:   "ab",
		b:   "x",
		num: 3,
	},
	{
		a:   "ab",
		b:   "xyz",
		num: 10,
	},
}

func TestRecursion(t *testing.T) {
	for idx, tcase := range gTestCases {
		got := &node{}
		RecursionTree(tcase.a, tcase.b, got)
		checkResult(t, idx, &tcase, got)
	}
}

func checkResult(t *testing.T, idx int, tcase *tc, got *node) {
	allInterleavings := []string{}
	walkTreeMain(got, &allInterleavings)

	if len(allInterleavings) != tcase.num {
		t.Errorf("index=%d, number not match. expected=%d, got=%d\n", idx, tcase.num, len(allInterleavings))
	} else {
		for _, cob := range allInterleavings {
			cobCheck := isInterleaving(tcase.a, tcase.b, cob)
			if !cobCheck {
				t.Errorf("index=%d, (%s) and (%s) invalid interleaving=%s\n", idx, tcase.a, tcase.b, cob)
			}
		}
	}
}
