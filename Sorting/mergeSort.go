// Merge Sort
package main

import (
	"fmt"
)

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	middle := len(array) / 2
	leftHalf := MergeSort(array[:middle])
	rightHalf := MergeSort(array[middle:])
	return mergeSortedArrays(leftHalf, rightHalf)
}

func mergeSortedArrays(left, right []int) []int {
	sortedArray := make([]int, len(left)+len(right))
	k, i, j := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sortedArray[k] = left[i]
			i++
		} else {
			sortedArray[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		sortedArray[k] = left[i]
		k++
		i++
	}
	for j < len(right) {
		sortedArray[k] = right[j]
		j++
		k++
	}
	return sortedArray
}

func main() {
	array := []int{71, 1, 17, -7, -17, -27, 12, 200, 8, 5, 70}
	actual := MergeSort(array)
	fmt.Println(actual)
}
