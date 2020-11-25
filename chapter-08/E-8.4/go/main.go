package main

// bruteForce function use brute-force to find all the sub array,
// and calculate the sum of each sub array. return the maximum sum
// O(N^3) time and O(1) space
func bruteForce(arr []int) int {
	max := -2 ^ 20
	aLen := len(arr)
	for i := 0; i < aLen; i++ {
		for j := i; j < aLen; j++ {
			cur := 0
			for k := i; k <= j; k++ {
				cur += arr[k]
			}
			if cur > max {
				max = cur
			}
		}
	}
	return max
}

// nativeDP function implement dp
// O(N) time and O(N) space
// extra：如果需要返回和最大的字数组的长度，则增加一个长度N的辅助数组L，
// L[i]表示以arr[i]结尾的和最大的子数组的长度，更新sum[i]时同步更新L[i]
func nativeDP(arr []int) int {
	arrLen := len(arr)
	if arrLen < 1 {
		return 0
	}

	// sum[i]表示以arr[i]结尾的和最大的子数组的和
	sum := make([]int, arrLen)
	sum[0] = arr[0]

	for i := 1; i < arrLen; i++ {
		arrNow := arr[i]
		sumBefore := sum[i-1]

		// sum[i] = max(sum[i-1]+arr[i], arr[i])
		if sumBefore <= 0 {
			// 如果sum[i-1]不大于0,则arr[i]单独构成一个和最大的子数组
			sum[i] = arrNow
		} else {
			// 否则，arr[i]和sum[i-1]构成一个和最大的字数组的和
			sum[i] = sumBefore + arrNow
		}
	}

	// find out the maximun sum
	max := sum[0]
	for i := 1; i < arrLen; i++ {
		if sum[i] > max {
			max = sum[i]
		}
	}

	return max
}

// Kadane function use Kadane's Algorithm.
// not really understand it.
// it only works for non-negative result
func Kadane(arr []int) int {
	arrLen := len(arr)
	if arrLen < 1 {
		return 0
	}

	maxSumSoFar := 0
	maxSumEndingHere := 0

	for i := 0; i < arrLen; i++ {
		maxSumEndingHere += arr[i]
		if maxSumEndingHere < 0 {
			maxSumEndingHere = 0
		}
		if maxSumSoFar < maxSumEndingHere {
			maxSumSoFar = maxSumEndingHere
		}
	}
	return maxSumSoFar
}

func main() {

}
