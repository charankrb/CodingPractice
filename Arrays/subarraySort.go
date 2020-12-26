// write a function that takes an array of atleast two integers that returns
// an array of starting and ending indices of the smallest sub array of the
// input array that needs to be sorted
// Ex : Input = [1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19]
//      Output = [3,9]
// The input array needs to be sorted from indices 3 to 9 inorder to make
// entire input array sorted in ascending order

package main

import (
	"fmt"
	"math"
)

func SubarraySort(array []int) []int {
	minSub, maxSub := math.MaxInt32, math.MinInt32
	for i, num := range array {
		if isOutOfOrder(array, i) {
			minSub = min(num, minSub)
			maxSub = max(num, maxSub)
		}
	}
	if minSub == math.MaxInt32 {
		return []int{-1, -1}
	}
	left, right := 0, len(array)-1
	for array[left] <= minSub {
		left++
	}
	for array[right] >= maxSub {
		right--
	}
	return []int{left, right}
}

func isOutOfOrder(array []int, i int) bool {
	if i == 0 {
		return array[i] > array[i+1]
	}
	if i == len(array)-1 {
		return array[i] < array[i-1]
	}
	return array[i] > array[i+1] || array[i] < array[i-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	output := SubarraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19})
	fmt.Println(output)
}
