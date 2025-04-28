package main

import (
	"fmt"
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {
	// Sort the input array to enable the two-pointer approach
	sort.Ints(array)

	// Initialize a slice to store the triplets that sum up to the target
	triplets := [][]int{}

	// Iterate through the array, fixing one number at a time
	for i := 0; i < len(array)-2; i++ {
		// Skip duplicate elements to avoid duplicate triplets
		if i > 0 && array[i] == array[i-1] {
			continue
		}

		// Initialize two pointers: one starting from the next element (left)
		// and the other starting from the end of the array (right)
		left, right := i+1, len(array)-1

		// Use the two-pointer approach to find pairs that sum up to (target - array[i])
		for left < right {
			// Calculate the current sum of the three numbers
			currentSum := array[i] + array[left] + array[right]

			// If the current sum equals the target, add the triplet to the result
			if currentSum == target {
				triplets = append(triplets, []int{array[i], array[left], array[right]})
				left += 1
				right -= 1

				// Skip duplicate elements for the left pointer
				for left < right && array[left] == array[left-1] {
					left += 1
				}
				// Skip duplicate elements for the right pointer
				for left < right && array[right] == array[right+1] {
					right -= 1
				}
			} else if currentSum < target {
				// If the current sum is less than the target, move the left pointer to the right
				left += 1
			} else {
				// If the current sum is greater than the target, move the right pointer to the left
				right -= 1
			}
		}
	}

	// Return the list of triplets that sum up to the target
	return triplets
}

func main() {
	// Test cases
	testCases := []struct {
		array    []int
		target   int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, 0, [][]int{{-1, -1, 2}, {-1, 0, 1}}},                                                // General case with multiple triplets
		{[]int{1, 2, 3, 4, 5}, 9, [][]int{{1, 3, 5}, {2, 3, 4}}},                                                         // General case with multiple triplets
		{[]int{0, 0, 0, 0}, 0, [][]int{{0, 0, 0}}},                                                                       // All elements are the same
		{[]int{-2, 0, 1, 1, 2}, 1, [][]int{{-2, 1, 2}, {0, 1, 1}}},                                                       // Case with duplicates
		{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 6, 6}, 5, [][]int{{-4, 3, 6}, {-2, 1, 6}, {-2, 2, 5}, {0, 1, 4}}}, // Complex case
		{[]int{}, 0, [][]int{}},                                // Empty array
		{[]int{1, 2}, 3, [][]int{}},                            // Array with less than 3 elements
		{[]int{-1, 0, 1, 2, -1, -4}, -5, [][]int{{-4, -1, 0}}}, // Negative target
		{[]int{1, 2, 3, 4, 5}, 50, [][]int{}},                  // No triplets found
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := ThreeNumberSum(testCase.array, testCase.target)
		fmt.Printf("Test Case %d: Input: %v, Target: %d\n", i+1, testCase.array, testCase.target)
		fmt.Printf("Expected: %v, Got: %v\n\n", testCase.expected, result)
	}
}
