package main

// Recursion function checks if subset of arr has sum of 'sum'
func Recursion(arr []int, sum int) bool {
	arrLen := len(arr)

	if arrLen == 0 && sum > 0 {
		return false
	}
	if sum == 0 {
		return true
	}
	if sum < 0 {
		return false
	}

	// the last element of arr may be the part of 'sum' or not
	return Recursion(arr[:arrLen-1], sum) || Recursion(arr[:arrLen-1], sum-arr[arrLen-1])
}

func main() {

}
