package main

import "testing"

var testCases = []struct {
	x    int
	y    int              // target point (x, y)
	dead *map[string]void // blocked path. key:(2,1)->(3,1) represents path from (2,1) to (3,1) is blocked
	n    int              // expected paths number
}{
	{x: 3, y: 4, n: 24,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 0, y: 0, n: 1,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},

	{x: 0, y: 1, n: 1,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 0, y: 2, n: 0,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 0, y: 3, n: 0,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 0, y: 100, n: 0,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 2, y: 0, n: 1,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 3, y: 0, n: 0,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 200, y: 0, n: 0,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 1, y: 1, n: 2,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 2, y: 1, n: 3,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 2, y: 2, n: 5,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 3, y: 1, n: 3,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 3, y: 2, n: 8,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 3, y: 4, n: 24,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 5, y: 8, n: 754,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 4, y: 7, n: 196,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
	{x: 10, y: 10, n: 82548,
		dead: &map[string]void{
			"(0,1)->(0,2)": void{},
			"(2,0)->(3,0)": void{},
			"(4,3)->(5,3)": void{},
			"(6,6)->(6,7)": void{},
		},
	},
}

func TestRecursion(t *testing.T) {
	for _, tc := range testCases {
		got := Recursion(tc.x, tc.y, tc.dead)
		if got != tc.n {
			t.Errorf("(x,y)=(%d,%d), expected=%d, got=%d\n", tc.x, tc.y, tc.n, got)
		}
	}
}

func TestNativeDP(t *testing.T) {
	for _, tc := range testCases {
		got, _ := NativeDP(tc.x, tc.y, tc.dead)
		if got != tc.n {
			t.Errorf("(x,y)=(%d,%d), expected=%d, got=%d\n", tc.x, tc.y, tc.n, got)
		}
	}
}
