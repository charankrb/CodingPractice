//https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
//Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
// find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
// Your solution must use only constant extra space.
// example 1:
// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: Because numbers[0] + numbers[1] == 9, we return [1, 2].
// example 2:
// Input: numbers = [2,3,4], target = 6
// Output: [1,3]
// Explanation: Because numbers[0] + numbers[2] == 6, we return [1, 3].

package main

import "fmt"

func twoSum(nums []int, t int) []int {
	if len(nums) < 2 {
		return nil // Return nil if the input array has less than 2 elements
	}
	start, end := 0, len(nums)-1 // Initialize two pointers: start and end
	for start < end {            // Loop until the two pointers meet
		sum := nums[start] + nums[end] // Calculate the sum of the two numbers at the pointers
		if sum == t {
			return []int{start + 1, end + 1} // If the sum equals the target, return the indices (1-indexed)
		} else if sum < t {
			start++ // If the sum is less than the target, move the start pointer to the right
		} else {
			end-- // If the sum is greater than the target, move the end pointer to the left
		}
	}
	return nil // Return nil if no such indices are found
}

func main() {
	// Test cases
	testCases := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},       // General case: valid pair exists
		{[]int{2, 3, 4}, 6, []int{1, 3}},            // General case: valid pair exists
		{[]int{-3, -1, 0, 2, 4, 6}, 3, []int{1, 6}}, // Negative numbers: valid pair exists
		{[]int{1, 2, 3, 4, 5}, 10, nil},             // No valid pair exists
		{[]int{1, 2, 3, 4, 4, 5}, 8, []int{3, 6}},   // Duplicate numbers: valid pair exists
		{[]int{1, 2}, 3, []int{1, 2}},               // Smallest valid input
		{[]int{1, 2}, 5, nil},                       // Smallest input with no valid pair
		{[]int{}, 5, nil},                           // Empty array
		{[]int{1}, 1, nil},                          // Single element array
		{[]int{1, 1, 1, 1, 1}, 2, []int{1, 5}},      // All elements are the same
		{[]int{1, 2, 3, 4, 5}, 7, []int{2, 5}},      // Valid pair in the middle of the array
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := twoSum(testCase.numbers, testCase.target)
		fmt.Printf("Test Case %d: Input: %v, Target: %d, Expected: %v, Got: %v\n",
			i+1, testCase.numbers, testCase.target, testCase.expected, result)
	}
}
