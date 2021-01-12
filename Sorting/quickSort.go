// Quick Sort

package main

import (
	"fmt"
)

func QuickSort(array []int) []int {
	return helper(array, 0, len(array)-1)
}

func helper(array []int, start, end int) []int {
	if start >= end {
		return array
	}
	pivot := start
	left := start + 1
	right := end

	for left <= right {
		if array[left] > array[pivot] && array[right] < array[pivot] {
			array[left], array[right] = array[right], array[left]
		}
		if array[left] <= array[pivot] {
			left++
		}
		if array[right] >= array[pivot] {
			right--
		}
	}
	array[right], array[pivot] = array[pivot], array[right]

	isLeftSubarraySmallest := right-1-start < end-(right+1)

	// performing sorting in the smallest sub array first.
	if isLeftSubarraySmallest {
		helper(array, start, right-1)
		helper(array, right+1, end)
	} else {
		helper(array, right+1, end)
		helper(array, start, right-1)
	}
	return array
}
func main() {
	array := []int{71, 1, 17, -7, -17, -27, 12, 200, 8, 5, 70}
	actual := QuickSort(array)
	fmt.Println(actual)
}
