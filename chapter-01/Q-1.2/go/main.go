package main

// AddPrevious function
func AddPrevious(arr []int) {
	AddRecursive(arr, len(arr)-1)
}

// AddRecursive function add from arr[0] to arr[idx], and put result at arr[idx]
func AddRecursive(arr []int, idx int) {
	if idx <= 0 {
		return
	}
	AddRecursive(arr, idx-1)
	arr[idx] += arr[idx-1]
}

// AddPreviousIterative function use iterative rather than recursive
func AddPreviousIterative(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for idx := 1; idx < len(arr); idx++ {
		arr[idx] += arr[idx-1]
	}
}

func main() {

}
