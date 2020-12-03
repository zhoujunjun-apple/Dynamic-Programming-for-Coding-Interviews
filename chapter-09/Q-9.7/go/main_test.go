package main

import "testing"

var gTestCases = [][]int{
	[]int{1, 6, 5, 3, 5, 7, 2, 4, 5},
	[]int{1, 4, 3},
	[]int{1, 2, 3, 4},
	[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 2, 4, 6, 8},
	[]int{1, 3, 5, 7, 9, 5, 2, 4, 6, 8, 0},
}

var sumTestCases = []struct {
	arr []int    // the input array
	sum int      // the input sum
	has bool     // the expected result: Are there two elements from arr have the sum 'sum'
	exp [][2]int // the expected two elements, possibly more than one pair
}{
	{
		arr: []int{9, 7, 5, 3, 1, 2, 4, 6, 8},
		sum: 10,
		has: true,
		exp: [][2]int{
			{1, 9}, {9, 1},
			{2, 8}, {8, 2},
			{3, 7}, {7, 3},
			{4, 6}, {6, 4},
		},
	},
	{
		arr: []int{9, 7, 5, 3, 1, 2, 4, 6, 8},
		sum: 16,
		has: true,
		exp: [][2]int{
			{7, 9}, {9, 7},
		},
	},
	{
		arr: []int{9, 7, 5, 3, 1, 2, 4, 6, 8},
		sum: 18,
		has: false,
		exp: [][2]int{},
	},
	{
		arr: []int{9, 7, 5, 3, 1, 2, 4, 6, 8},
		sum: 2,
		has: false,
		exp: [][2]int{},
	},
	{
		arr: []int{9, 7, 5, 3, 1, 2, 4, 6, 8},
		sum: 1,
		has: false,
		exp: [][2]int{},
	},
	{
		arr: []int{9, 7, 5, 3, 1, 2, 4, 6, 8},
		sum: 5,
		has: true,
		exp: [][2]int{
			{1, 4}, {4, 1},
		},
	},
}

func TestFindTwoSum(t *testing.T) {
	for idx, tc := range sumTestCases {
		gotBool, gotLeft, gotRight := findTwoSum(tc.arr, tc.sum)
		if gotBool != tc.has {
			t.Errorf("index=%d, expected=%t, got=%t\n", idx, tc.has, gotBool)
			continue
		}

		if gotBool == tc.has && gotBool == true {
			found := false
			for _, item := range tc.exp {
				if gotLeft == item[0] && gotRight == item[1] {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("index=%d, expected sum=%d, invalid got=(%d, %d)\n", idx, tc.sum, gotLeft, gotRight)
			}
		}
	}
}

func TestQuickSort(t *testing.T) {
	for idx, arr := range gTestCases {
		quickSort(&arr, 0, len(arr)-1)
		checkAscending(t, idx, &arr)
	}
}

func checkAscending(t *testing.T, idx int, arr *[]int) {
	for i := 1; i < len(*arr); i++ {
		before, now := (*arr)[i-1], (*arr)[i]
		if before > now {
			t.Errorf("index=%d, arr: %v, got: %d>%d\n", idx, *arr, before, now)
		}
	}
}

func TestPartition(t *testing.T) {
	for idx, tarr := range gTestCases {
		arrLen := len(tarr)
		pidx := partition(&tarr, 0, arrLen-1, (arrLen-1)/2)
		checkPartition(t, &tarr, idx, pidx)
	}
}

func checkPartition(t *testing.T, arr *[]int, idx, pidx int) {
	pVal := (*arr)[pidx]
	for leftIdx := 0; leftIdx < pidx; leftIdx++ {
		if (*arr)[leftIdx] > pVal {
			t.Errorf("index=%d, arr: %v, larger value: %d on the left of pivot value: %d\n",
				idx, *arr, (*arr)[leftIdx], pVal)
		}
	}
	for rIdx := pidx; rIdx < len(*arr); rIdx++ {
		if (*arr)[rIdx] < pVal {
			t.Errorf("index=%d, arr: %v, lower value: %d on the right of pivot value: %d\n",
				idx, *arr, (*arr)[rIdx], pVal)
		}
	}
}
