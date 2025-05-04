// Binary Serach

package main

import "fmt"

func BinarySearch(array []int, target int) int {
	return bsHelper(array, target, 0, len(array)-1)
}

func bsHelper(array []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if target == array[mid] {
		return mid
	}
	if array[mid] > target {
		return bsHelper(array, target, left, mid-1)
	} else {
		return bsHelper(array, target, mid+1, right)
	}
}

func main() {
	array := []int{1, 9, 12, 6, 5, 2, 8, 7, 15, 13, 14}
	actual := BinarySearch(array, 8)
	fmt.Println(actual)
}
