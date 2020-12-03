package main

// findTwoSum function checks if arr has two elements which have sum 'sum'
// use quickSort to sort the arr, then linearly scan the sorted arr
func findTwoSum(arr []int, sum int) (bool, int, int) {
	arrLen := len(arr)
	quickSort(&arr, 0, arrLen-1)

	left, right := 0, arrLen-1
	for left < right {
		add := arr[left] + arr[right]
		if add > sum {
			right--
		} else if add < sum {
			left++
		} else {
			return true, arr[left], arr[right]
		}
	}
	return false, -1, -1
}

// quickSort function sorts arr[left, right] by ascending order
func quickSort(arr *[]int, left, right int) {
	if left+1 == right {
		leftVal, rightVal := (*arr)[left], (*arr)[right]
		if leftVal > rightVal {
			(*arr)[left] = rightVal
			(*arr)[right] = leftVal
		}
	} else {
		var pivot int = left + (right-left)/2
		// partition函数将arr[left, right]分为三个部分
		pidx := partition(arr, left, right, pivot)
		// if use loosePartition function， quickSort function must change it's logic

		if pidx-left > 1 {
			// 递归对左边区间排序
			quickSort(arr, left, pidx-1)
		}
		if right-pidx > 1 {
			//  递归对右边区间排序
			quickSort(arr, pidx+1, right)
		}
	}
}

// partition function 将arr[i, j]区间严格的重新排列为三个子区间，依次为：
// 小于arr[pivot]的元素、等于arr[pivot]的元素 和 大于arr[pivot]的元素
// 返回值是第二个子区间的任意索引
func partition(arr *[]int, i, j, pivot int) int {
	pVal := (*arr)[pivot]

	// arr[0, i) contains the elements that less than pVal;
	// arr[i, k) contains the elements that equal to pVal;
	// arr(j, end] contains the elements that larger than pVal;
	k := i
	for k <= j {
		// move k towards to j until them meet each other
		ak := (*arr)[k]
		if ak > pVal {
			// larger than pVal, move it to arr(j, end] range and amplify its left range to arr(j-1, end)
			swap(arr, k, j)
			j--
		} else if ak < pVal {
			// less than pVal, swap it with the first element arr[i] that equals to pVal,
			// then amplify arr[0, i) to arr[0, i+1), and amplify arr[i, k) to arr[i+1, k+1)
			swap(arr, k, i)
			i++
			k++
		} else {
			// equal with pVal, only need to amplify arr[i, k) to arr[i, k+1)
			k++
		}
	}

	// i represents the first position of arr[i, k), in which elements value have the same value with pVal
	return i
}

// loosePartition function 将arr[i, j]切分为两部分：
// 不超过arr[pivot]的元素 和 不小于arr[pivot]的元素
// 返回值表示这两个部分的切分点
func loosePartition(arr *[]int, i, j, pivot int) int {
	pVal := (*arr)[pivot]

	for i <= j {
		now := (*arr)[i]

		if now < pVal {
			i++
		} else {
			swap(arr, i, j)
			j--
		}
	}

	return i
}

// swap function swaps arr[i] and arr[j] in place
func swap(arr *[]int, i, j int) {
	iv, jv := (*arr)[i], (*arr)[j]
	(*arr)[i] = jv
	(*arr)[j] = iv
}

func main() {

}
