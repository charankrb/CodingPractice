// https://leetcode.com/problems/contains-duplicate/description/
// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
//
//Example 1:
//Input: nums = [1,2,3,1]
//Output: true
//Explanation: The element 1 appears twice in the array.
//
//Example 2:
//Input: nums = [1,2,3,4]
//Output: false
//Explanation: Each element appears once and is distinct.

package contains_duplicate

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool) // Create a map to track seen numbers
	for _, num := range nums { // Iterate through the array
		if seen[num] { // If the number is already in the map, return true
			return true
		}
		seen[num] = true // Otherwise, add it to the map
	}
	return false // If no duplicates were found, return false
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 1},                             // Duplicate exists
		{1, 2, 3, 4},                             // No duplicates
		{},                                       // Empty array
		{1},                                      // Single element
		{1, 1, 1, 1},                             // All elements are the same
		{1, 2, 3, 4, 5, 6, 7, 8},                 // All distinct elements
		{-1, -2, -3, -1},                         // Negative numbers with duplicates
		{-1, -2, -3, -4},                         // Negative numbers without duplicates
		{0, 0},                                   // Array with only zeros
		{int(^uint(0) >> 1), int(^uint(0) >> 1)}, // Max integer duplicate
		{-int(^uint(0) >> 1), -int(^uint(0) >> 1)}, // Min integer duplicate
	}

	// Run the function on each test case and print the results
	for i, testCase := range testCases {
		fmt.Printf("Test Case %d: Input: %v -> Output: %v\n", i+1, testCase, containsDuplicate(testCase))
	}
}
