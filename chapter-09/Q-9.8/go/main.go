package main

// subset type represent the checking result:
// key: int. elements from input array;
// value: int. duplicate time for key
type subset map[int]int

// Recursion function return the result with recursive solution
// Check if subset of arr has sum of 'sum', if so, return the subset
func Recursion(arr []int, sum int) (bool, subset) {
	arrLen := len(arr)

	if sum == 0 {
		return true, subset{}
	}
	if sum < 0 || (sum > 0 && arrLen == 0) {
		return false, nil
	}

	// check if the last element of arr is part of the sum
	withoutBool, withoutSet := Recursion(arr[:arrLen-1], sum)
	if withoutBool {
		// the 'sum' can be calculated without the last element of arr
		return withoutBool, withoutSet
	}

	// the 'sum' cannot be calculated without the last element of arr
	arrLast := arr[arrLen-1]
	withBool, withSet := Recursion(arr[:arrLen-1], sum-arrLast)
	if withBool {
		// the 'sum' can be calculated with the last element of arr
		withSet[arrLast]++
		return withBool, withSet
	}
	return false, nil
}

func main() {

}
