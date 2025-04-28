// Problem: Longest Consecutive Sequence
// URL: https://leetcode.com/problems/longest-consecutive-sequence/description/
//
// Example 1:
// Input: nums = [100, 4, 200, 1, 3, 2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore, its length is 4.
//
// Example 2:
// Input: nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
// Output: 9
// Explanation: The longest consecutive elements sequence is [0, 1, 2, 3, 4, 5, 6, 7, 8]. Therefore, its length is 9.
package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0 // Return 0 if the input array is empty
	}

	// Step 1: Create a map to store the numbers for O(1) lookups
	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true // Add each number to the map
	}
	// Example: nums = [100, 4, 200, 1, 3, 2]
	// After this step, numsMap = {100: true, 4: true, 200: true, 1: true, 3: true, 2: true}

	// Step 2: Initialize the longest sequence length to 0
	longestCon := 0

	// Step 3: Iterate through the map to find the longest consecutive sequence
	for num := range numsMap {
		// Check if the current number is the start of a sequence
		if !numsMap[num-1] { // If the previous number is not in the map, this is the start of a sequence
			currentNum := num           // Start a new sequence from the current number
			currentCon := 1             // Initialize the current sequence length to 1
			for numsMap[currentNum+1] { // Check for consecutive numbers in the map
				currentNum++ // Increment the current number
				currentCon++ // Increment the current sequence length
			}
			// Example dry run:
			// For num = 1:
			// - currentNum = 1, currentCon = 1
			// - currentNum = 2, currentCon = 2
			// - currentNum = 3, currentCon = 3
			// - currentNum = 4, currentCon = 4
			// Longest sequence so far: [1, 2, 3, 4], currentCon = 4

			// Update the longest sequence length if needed
			longestCon = max(longestCon, currentCon)
			// Example: longestCon = max(0, 4) = 4
		}
	}

	// Step 4: Return the length of the longest consecutive sequence
	return longestCon
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := [][]int{
		{100, 4, 200, 1, 3, 2},               // General case
		{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},       // Longest sequence spans the entire array
		{},                                   // Empty array
		{1},                                  // Single element
		{1, 2, 0, 1},                         // Duplicates in the array
		{10, 5, 12, 3, 55, 30, 11, 4, 6},     // Disjoint sequences
		{-1, -2, -3, -4, -5},                 // Negative numbers
		{1, 2, 3, 4, 5},                      // Already consecutive
		{5, 4, 3, 2, 1},                      // Reverse consecutive
		{1, 3, 5, 7, 9},                      // No consecutive sequence longer than 1
		{-10, -9, -8, -7, -6, 0, 1, 2, 3, 4}, // Mixed positive and negative numbers
	}

	// Run the function on each test case and print the results
	for i, testCase := range testCases {
		fmt.Printf("Test Case %d: Input: %v -> Output: %d\n", i+1, testCase, longestConsecutive(testCase))
	}
}
