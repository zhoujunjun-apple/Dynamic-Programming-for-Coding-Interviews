package main

import "fmt"

type void struct{}

// Recursion function use recursive method
// check blocked path
func Recursion(x, y int, dead *map[string]void) int {
	if x < 0 || y < 0 {
		return 0
	}
	if x == 0 && y == 0 {
		return 1
	}

	// y-axis only need to consider the upward path
	if x == 0 && y > 0 {
		for yi := 1; yi <= y; yi++ {
			upRoad := fmt.Sprintf("(%d,%d)->(%d,%d)", x, yi-1, x, yi)
			if _, upBlocked := (*dead)[upRoad]; upBlocked {
				// if previous path is blocked, then no valid path to (x, y)
				return 0
			}
		}
		return 1
	}

	// x-axis only need to consider the rightward path
	if x > 0 && y == 0 {
		for xi := 1; xi <= x; xi++ {
			rightRoad := fmt.Sprintf("(%d,%d)->(%d,%d)", xi-1, y, xi, y)
			if _, rightBlocked := (*dead)[rightRoad]; rightBlocked {
				// if left path is blocked, then no valid path to (x, y)
				return 0
			}
		}
		return 1
	}

	// check if (x-1,y)->(x,y) and (x,y-1)->(x,y) is blocked
	rightRoad := fmt.Sprintf("(%d,%d)->(%d,%d)", x-1, y, x, y)
	upRoad := fmt.Sprintf("(%d,%d)->(%d,%d)", x, y-1, x, y)

	_, rightBlocked := (*dead)[rightRoad]
	_, upBlocked := (*dead)[upRoad]

	// if some path is blocked, ignore that road
	paths := 0
	if rightBlocked && upBlocked {
		// both directions are blocked
		paths = 0
	} else if rightBlocked && !upBlocked {
		// rightward direction is blocked. consider only upward direction
		paths = Recursion(x, y-1, dead)
	} else if !rightBlocked && upBlocked {
		// upward direction is blocked. consider only rightward direction
		paths = Recursion(x-1, y, dead)
	} else {
		// both directions are not blocked. consider both directions.
		paths = Recursion(x-1, y, dead) + Recursion(x, y-1, dead)
	}

	return paths
}

// NativeDP function use dp
// same path-checking logic with Recursion function
func NativeDP(x, y int, dead *map[string]void) (int, error) {
	if x < 0 || y < 0 {
		return 0, fmt.Errorf("x, y must not less than zero, but got x=%d, y=%d", x, y)
	}
	if x == 0 && y == 0 {
		return 1, nil
	}
	if x == 0 && y > 0 {
		for yi := 1; yi <= y; yi++ {
			upRoad := fmt.Sprintf("(%d,%d)->(%d,%d)", x, yi-1, x, yi)
			if _, upBlocked := (*dead)[upRoad]; upBlocked {
				return 0, nil
			}
		}
		return 1, nil
	}
	if x > 0 && y == 0 {
		for xi := 1; xi <= x; xi++ {
			rightRoad := fmt.Sprintf("(%d,%d)->(%d,%d)", xi-1, y, xi, y)
			if _, rightBlocked := (*dead)[rightRoad]; rightBlocked {
				return 0, nil
			}
		}
		return 1, nil
	}

	// note the dp has (x+1)*(y+1) dimension
	dp := make([][]int, x+1)
	for i := range dp {
		dp[i] = make([]int, y+1)
	}

	// check the path from (0,0) to (i,0)
	for xi := 0; xi <= x; xi++ {
		rightRoad := fmt.Sprintf("(%d,%d)->(%d,%d)", xi-1, 0, xi, 0)
		if _, rightBlocked := (*dead)[rightRoad]; rightBlocked {
			for i := xi; i <= x; i++ {
				dp[i][0] = 0
			}
			break
		}
		dp[xi][0] = 1
	}

	// check the path from (0,0) to (0,i)
	for yi := 0; yi <= y; yi++ {
		upRoad := fmt.Sprintf("(%d,%d)->(%d,%d)", 0, yi-1, 0, yi)
		if _, upBlocked := (*dead)[upRoad]; upBlocked {
			for i := yi; i <= y; i++ {
				dp[0][i] = 0
			}
			break
		}
		dp[0][yi] = 1
	}

	for xi := 1; xi <= x; xi++ {
		for yi := 1; yi <= y; yi++ {
			rightRoad := fmt.Sprintf("(%d,%d)->(%d,%d)", xi-1, yi, xi, yi)
			upRoad := fmt.Sprintf("(%d,%d)->(%d,%d)", xi, yi-1, xi, yi)

			_, rightBlocked := (*dead)[rightRoad]
			_, upBlocked := (*dead)[upRoad]

			paths := 0
			if rightBlocked && upBlocked {
				paths = 0
			} else if rightBlocked && !upBlocked {
				paths = dp[xi][yi-1]
			} else if !rightBlocked && upBlocked {
				paths = dp[xi-1][yi]
			} else {
				paths = dp[xi-1][yi] + dp[xi][yi-1]
			}

			dp[xi][yi] = paths
		}
	}

	return dp[x][y], nil
}

func main() {

}
