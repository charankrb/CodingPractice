// https://leetcode.com/problems/two-sum/description/
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Explanation: Because nums[1] + nums[2] == 6, we return [1, 2].

package main

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil // Return nil if the input array has less than 2 elements
	}
	numsMap := make(map[int]int) // Create a map to store the numbers and their indices
	for i, num := range nums {
		complement := target - num                      // Calculate the complement of the current number
		if index, found := numsMap[complement]; found { // Check if the complement exists in the map
			return []int{index, i} // Return the indices of the two numbers that add up to target
		}
		numsMap[num] = i // Add the current number and its index to the map
	}
	return nil // Return nil if no solution is found (should not happen as per problem statement)
}

func main() {
	// Test cases
	testCases := [][]int{
		{2, 7, 11, 15}, // General case
		{3, 2, 4},      // Two numbers that add up to target
		{3, 3},         // Same number twice
		{1, 2, 3, 4},   // No solution
		{0, -1, 2, -3}, // Negative numbers
	}

	targets := []int{
		9,
		6,
		6,
		5,
		-1,
	}

	for i, testCase := range testCases {
		result := twoSum(testCase, targets[i])
		if result != nil {
			println("Input:", testCase, "Target:", targets[i], "Output:", result[0], result[1])
		} else {
			println("Input:", testCase, "Target:", targets[i], "Output: No solution")
		}
	}
}
