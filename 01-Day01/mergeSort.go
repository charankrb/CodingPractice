package main

import "fmt"

// MergeSort function sorts an array using the merge sort algorithm
func MergeSort(array []int) []int {
	if array == nil || len(array) <= 1 {
		return array // Return the array if it's nil or has 1 or 0 elements
	}

	mid := len(array) / 2           // Find the middle index
	left := MergeSort(array[:mid])  // Recursively sort the left half
	right := MergeSort(array[mid:]) // Recursively sort the right half

	return merge(left, right) // Merge the sorted halves	}
}

func merge(left, right []int) []int {
	sortedarray := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0 // Initialize indices for left, right, and sorted arrays
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sortedarray[k] = left[i] // If left element is smaller, add it to sorted array
			i++
		} else {
			sortedarray[k] = right[j] // If right element is smaller, add it to sorted array
			j++
		}
		k++
	}
	//	Add remaining elements from left array
	for i < len(left) {
		sortedarray[k] = left[i]
		i++
		k++
	}
	// Add remaining elements from right array
	for j < len(right) {
		sortedarray[k] = right[j]
		j++
		k++
	}
	return sortedarray // Return the merged sorted array
}

func main() {
	// Test cases
	testCases := [][]int{
		{71, 1, 17, -7, -17, -27, 12, 200, 8, 5, 70}, // General case
		{},               // Empty array
		{5},              // Single element
		{2, 2, 2, 2},     // All duplicates
		{-1, -2, -3, -4}, // All negative numbers
		{1, 2, 3, 4, 5},  // Already sorted
		{5, 4, 3, 2, 1},  // Reverse sorted
		nil,              // Nil input
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: %v -> Sorted: %v\n", testCase, MergeSort(testCase))
	}
}
