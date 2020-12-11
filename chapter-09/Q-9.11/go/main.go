package main

// NativeDP function implement dp with two extra auxiliary array,
// recording the LBS in both ascending trend and descreasing trend
func NativeDP(arr []int) int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arrLen
	}

	// tp[i] represents the length of LBS ending with arr[i] with ascending trend
	// this LBS is in a monotonically ascending order
	ap := make([]int, arrLen)
	ap[0] = 1

	// dp[i] represents the length of LBS ending with arr[i] with descreasing trend
	// this LBS is in a monotonically descreasing order or a first-ascending-then-descreasing order
	dp := make([]int, arrLen)
	dp[0] = 1

	for dpi := 1; dpi < arrLen; dpi++ {
		amax, dmax := 1, 1
		arc := arr[dpi] // current integer on arr

		// looking backward
		for bki := dpi - 1; bki >= 0; bki-- {
			arb := arr[bki] // bki-th backward integer on arr

			if arc > arb {
				// current integer larger than bki-th backward integer
				// arc can only append upon the ascending LBS in bki-th position
				amax = selectMax(amax, ap[bki]+1)
			} else if arc < arb {
				// current integer less than bki-th backward integer
				// arc can append upon the descreasing LBS in bki-th position,
				dmax = selectMax(dmax, dp[bki]+1)
				// ... or upon the ascending LBS in bki-th position, forming a first-ascending-then-descreasing order
				dmax = selectMax(dmax, ap[bki]+1)
				// both forming a new descreasing LBS in dpi-th position, with length dmax
			}
		}

		// record length of the longest LBS in both ascending order and descreasing order
		ap[dpi] = amax
		dp[dpi] = dmax
	}

	max := selectMax(findMax(ap), findMax(dp))
	return max
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

func selectMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
