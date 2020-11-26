package main

import (
	"fmt"
	"math"
)

type cell struct {
	x int
	y int
}

const (
	xMax              = 7 // the maximum valid x of chess
	yMax              = 7 // the maximum valid y of chess
	recursionMaxDepth = 6 // the maximum of recursive depth in the question since the maximum moves in a chess is 5
)

// Recursion function represent the minimum number of moves from start cell to end cell
// parameter `depth` controls the recursive depth
func Recursion(start, end cell, depth int) int {
	if depth >= recursionMaxDepth {
		return math.MaxUint32
	}
	if start.x == end.x && start.y == end.y {
		return 0
	}
	if isOneStepAway(start, end) {
		return 1
	}

	exAddOne, exMinusOne, exAddTwo, exMinusTwo := end.x+1, end.x-1, end.x+2, end.x-2
	eyAddOne, eyMinusOne, eyAddTwo, eyMinusTwo := end.y+1, end.y-1, end.y+2, end.y-2

	mins := []int{}

	ndepth := depth + 1
	if isXValid(exAddOne) {
		nx := exAddOne

		// (i+1, j)
		mins = append(mins, Recursion(start, cell{x: nx, y: end.y}, ndepth))
		if isYValid(eyAddOne) {
			// (i+1, j+1)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyAddOne}, ndepth))
		}
		if isYValid(eyAddTwo) {
			// (i+1, j+2)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyAddTwo}, ndepth))
		}
		if isYValid(eyMinusOne) {
			// (i+1, j-1)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyMinusOne}, ndepth))
		}
		if isYValid(eyMinusTwo) {
			// (i+1, j-2)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyMinusTwo}, ndepth))
		}
	}
	if isXValid(exMinusOne) {
		nx := exMinusOne
		// (i-1, j)
		mins = append(mins, Recursion(start, cell{x: nx, y: end.y}, ndepth))
		if isYValid(eyAddOne) {
			// (i-1, j+1)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyAddOne}, ndepth))
		}
		if isYValid(eyAddTwo) {
			// (i-1, j+2)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyAddTwo}, ndepth))
		}
		if isYValid(eyMinusOne) {
			// (i-1, j-1)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyMinusOne}, ndepth))
		}
		if isYValid(eyMinusTwo) {
			// (i-1, j-2)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyMinusTwo}, ndepth))
		}
	}
	if isXValid(exAddTwo) {
		nx := exAddTwo
		if isYValid(eyAddOne) {
			// (i+2, j+1)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyAddOne}, ndepth))
		}
		if isYValid(eyMinusOne) {
			// (i+2, j-1)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyMinusOne}, ndepth))
		}
	}
	if isXValid(exMinusTwo) {
		nx := exMinusTwo
		if isYValid(eyAddOne) {
			// (i-2, j+1)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyAddOne}, ndepth))
		}
		if isYValid(eyMinusOne) {
			// (i-2, j-1)
			mins = append(mins, Recursion(start, cell{x: nx, y: eyMinusOne}, ndepth))
		}
	}
	if isYValid(eyAddOne) {
		// (i, j+1)
		mins = append(mins, Recursion(start, cell{x: end.x, y: eyAddOne}, ndepth))
	}
	if isYValid(eyMinusOne) {
		// (i, j-1)
		mins = append(mins, Recursion(start, cell{x: end.x, y: eyMinusOne}, ndepth))
	}

	m, err := getMin(mins)
	if err == nil {
		return m + 1
	}
	return 0
}

// check if two cell in one step away
func isOneStepAway(a, b cell) bool {
	dx, dy := a.x-b.x, a.y-b.y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}

	// dx, dy now no less than zero
	if dx == 0 && dy == 0 {
		return false
	}
	if dx <= 1 && dy <= 1 {
		return true
	}
	if (dx == 2 && dy == 1) || (dx == 1 && dy == 2) {
		return true
	}
	return false
}

// isXValid function check if cell.x is valid
func isXValid(x int) bool {
	if x >= 0 && x <= xMax {
		return true
	}
	return false
}

// isYValid function check if cell.y is valid
func isYValid(y int) bool {
	if y >= 0 && y <= yMax {
		return true
	}
	return false
}

// getMin function return the minimum of a int slice
func getMin(is []int) (int, error) {
	if len(is) < 1 {
		return -1, fmt.Errorf("input slice cannot empty")
	}

	min := is[0]
	for i := 1; i < len(is); i++ {
		if is[i] < min {
			min = is[i]
		}
	}

	return min, nil
}

// NativeDP function try use to implement dp
// NOT IMPLEMENTED!
func NativeDP(start, end cell) int {
	if start.x == end.x && start.y == end.y {
		return 0
	}
	if isOneStepAway(start, end) {
		return 1
	}

	dp := make([][]int, xMax)
	for i := range dp {
		dp[i] = make([]int, yMax)
	}

	// have no idea about how to initializing the dp matrix
	// also have no clue about how to fill up the dp: what order do I use to fill up dp matrix?
	return 0
}

func main() {

}
