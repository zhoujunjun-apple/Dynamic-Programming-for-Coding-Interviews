package main

// RecursionMain function is the entry function that returns
// the length of LMIS of arr by checking each subarr [arr:i] and
// find the maximum length
func RecursionMain(arr []int) int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arrLen
	}

	max := 0
	for i := 1; i <= arrLen; i++ {
		got := Recursion(arr[:i])
		if got > max {
			max = got
		}
	}

	return max
}

// Recursion function returns the length of LMIS ending with arr[arrLen-1]
func Recursion(arr []int) int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arrLen
	}

	arrLast := arr[arrLen-1]
	candidate := []int{1} // the minimal length of LMIS ending with arr[arrLen-1]

	for span := arrLen - 1; span > 0; span-- {
		if arrLast > arr[span-1] {
			// arrLast can append on the LMIS ending with arr[span-1]
			candidate = append(candidate, Recursion(arr[:span])+1)
		}
	}

	return findMax(candidate)
}

func findMax(arr []int) (max int) {
	max = arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	return
}

// NativeDP function use dp with the same logic with RecursionMain function
func NativeDP(arr []int) int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arrLen
	}

	// dp[i] represents the length of LMIS ending with arr[i]
	dp := make([]int, arrLen)
	dp[0] = 1

	// fill up the dp array from left to right
	for dpi := 1; dpi < arrLen; dpi++ {
		max := 1

		// at each postion dpi, checking backward
		for bki := dpi - 1; bki >= 0; bki-- {
			if arr[dpi] > arr[bki] {
				bkv := dp[bki] + 1
				if bkv > max {
					max = bkv
				}
			}
		}

		dp[dpi] = max
	}

	return findMax(dp)
}

func main() {

}
