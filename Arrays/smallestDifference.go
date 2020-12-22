// Given Two arrays find the smallest difference between two pair
// of numbers from two different list
package main

import (
	"fmt"
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)

	// O/P of sorted array1 and array2
	fmt.Println(array1)
	fmt.Println(array2)

	numsdiff := []int{}
	smallestDiff := math.MaxInt32
	currDiff := math.MaxInt32
	index1, index2 := 0, 0

	for index1 < len(array1) && index2 < len(array2) {
		first, second := array1[index1], array2[index2]
		if array1[index1] > array2[index2] {
			currDiff = array1[index1] - array2[index2]
			index2++
		} else if array1[index1] < array2[index2] {
			currDiff = array2[index2] - array1[index1]
			index1++
		} else {
			return []int{array1[index1], array2[index2]}
		}
		if smallestDiff > currDiff {
			smallestDiff = currDiff
			numsdiff = []int{first, second}
			fmt.Println(smallestDiff, numsdiff)
		}
	}
	return numsdiff
}

func main() {
	output := SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})
	fmt.Println(output)
}
