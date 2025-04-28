package main

import "fmt"

func quicksort(arr []int) {
	quicksortHelper(arr, 0, len(arr)-1)
}

func quicksortHelper(arr []int, low, high int) {
	if low >= high {
		return // Base case: if the array has 1 or 0 elements, it's already sorted
	}
	pivotIndex := partition(arr, low, high)  // Partition the array
	quicksortHelper(arr, low, pivotIndex-1)  // Recursively sort the left subarray
	quicksortHelper(arr, pivotIndex+1, high) // Recursively sort the right subarray
}

func partition(arr []int, low, high int) int {
	pivot := arr[low] // Choose the first element as the pivot
	i, j := low, high // Initialize pointers for partitioning
	for i < j {
		for i <= high && arr[i] <= pivot {
			i++ // Move the left pointer to the right while elements are less than or equal to pivot
		}
		for j >= low && arr[j] > pivot {
			j-- // Move the right pointer to the left while elements are greater than pivot
		}
		if i < j { // If the left pointer is less than the right pointer, swap elements
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[j] = arr[j], arr[low] // Swap the pivot with the right pointer
	return j                            // Return the index of the pivot after partitioning
}

func main() {
	// Test cases
	testCases := [][]int{
		{71, 1, 17, -7, -17, -27, 12, 200, 8, 5, 70}, // General case
		{},                  // Empty array
		{5},                 // Single element
		{2, 2, 2, 2},        // All duplicates
		{-1, -2, -3, -4},    // All negative numbers
		{1, 2, 3, 4, 5},     // Already sorted
		{5, 4, 3, 2, 1},     // Reverse sorted
		{100, -100, 0},      // Mixed positive, negative, and zero
		{0, 0, 0, 0, 0},     // All zeros
		{-5, 0, 5, -10, 10}, // Mixed positive, negative, and zero
		{int(^uint(0) >> 1), -int(^uint(0) >> 1)}, // Max and min integers
	}

	// Run Quick Sort on each test case and print the results
	for _, testCase := range testCases {
		fmt.Printf("Input: %v -> ", testCase)
		quicksort(testCase)
		fmt.Printf("Sorted: %v\n", testCase)
	}
}
