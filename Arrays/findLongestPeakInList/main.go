package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	actual := LongestPeak(array)
	fmt.Println(actual)
}

func LongestPeak(array []int) int {
	longestPeakLength := 0
	i := 1
	for i < len(array)-1 {
		// find peak in the given array
		isPeak := array[i-1] < array[i] && array[i] > array[i+1]
		if !isPeak {
			// if not peak continue
			i += 1
			continue
		}
		leftIndex := i - 2
		rightIndex := i + 2
		// find the longest peak by traversing the left and right index of the peak
		for leftIndex >= 0 && array[leftIndex] < array[leftIndex+1] {
			leftIndex -= 1
		}
		for rightIndex < len(array) && array[rightIndex] < array[rightIndex-1] {
			rightIndex += 1
		}
		currentPeakLength := rightIndex - leftIndex - 1
		if currentPeakLength > longestPeakLength {
			longestPeakLength = currentPeakLength
		}
		i = rightIndex
	}
	return longestPeakLength
}
