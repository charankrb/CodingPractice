// https://leetcode.com/problems/product-of-array-except-self/description/
//Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Explanation: The product of all the elements except nums[0] is 2*3*4 = 24. The product of all the elements except nums[1] is 1*3*4 = 12. The product of all the elements except nums[2] is 1*2*4 = 8. The product of all the elements except nums[3] is 1*2*3 = 6.
// Example 2:
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]
// Explanation: The product of all the elements except nums[0] is 1*0*-3*3 = 0. The product of all the elements except nums[1] is -1*0*-3*3 = 0. The product of all the elements except nums[2] is -1*1*-3*3 = 9. The product of all the elements except nums[3] is -1*1*0*3 = 0. The product of all the elements except nums[4] is -1*1*0*-3 = 0.

package main

import "fmt"

func productExceptSelfDivision(nums []int) []int { // O(n) time, O(1) space division method
	// Initialize the result array with the same length as nums
	result := make([]int, len(nums))
	// Calculate the product of all elements in nums
	product := 1
	zeroCount := 0
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			product *= num
		}
	}
	if zeroCount > 1 {
		// If there are more than one zero, all products will be zero
		return result
	} else if zeroCount == 1 {
		// If there is one zero, only the position of the zero will have a non-zero product
		for i, num := range nums {
			if num == 0 {
				result[i] = product
			} else {
				result[i] = 0
			}
		}
	} else {
		// If there are no zeros, calculate the product for each position
		for i, num := range nums {
			result[i] = product / num
		}
	}
	return result
}

func productExceptSelf(nums []int) []int { // O(n) time, O(1) space prefix and suffix product
	// Check if the input array is empty
	if len(nums) == 0 {
		return []int{} // Return an empty array if the input is empty
	}

	// Initialize the result array with the same length as nums
	result := make([]int, len(nums))

	// Step 1: Calculate the prefix product for each position
	// Example: nums = [1, 2, 3, 4]
	// prefixProduct starts as 1
	prefixProduct := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefixProduct // Store the prefix product in result[i]
		prefixProduct *= nums[i]  // Update prefixProduct by multiplying with nums[i]
		// Dry run:
		// i = 0: result = [1, _, _, _], prefixProduct = 1 * 1 = 1
		// i = 1: result = [1, 1, _, _], prefixProduct = 1 * 2 = 2
		// i = 2: result = [1, 1, 2, _], prefixProduct = 2 * 3 = 6
		// i = 3: result = [1, 1, 2, 6], prefixProduct = 6 * 4 = 24
	}

	// Step 2: Calculate the suffix product and multiply it with the prefix product
	// suffixProduct starts as 1
	suffixProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffixProduct // Multiply the current result[i] with suffixProduct
		suffixProduct *= nums[i]   // Update suffixProduct by multiplying with nums[i]
		// Dry run:
		// i = 3: result = [1, 1, 2, 6], suffixProduct = 1 * 4 = 4
		// i = 2: result = [1, 1, 8, 6], suffixProduct = 4 * 3 = 12
		// i = 1: result = [1, 12, 8, 6], suffixProduct = 12 * 2 = 24
		// i = 0: result = [24, 12, 8, 6], suffixProduct = 24 * 1 = 24
	}

	// Final result after combining prefix and suffix products
	return result
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4},       // General case
		{-1, 1, 0, -3, 3},  // Contains a zero
		{0, 0, 0, 0},       // All zeros
		{1},                // Single element
		{1, 2},             // Two elements
		{-1, -2, -3, -4},   // All negative numbers
		{1, 0, 2, 0, 3},    // Multiple zeros
		{100000, 200000},   // Large numbers
		{},                 // Empty array
		{1, 1, 1, 1},       // All elements are the same
		{1, 2, 3, 0, 4, 5}, // Mixed numbers with a zero
	}

	// Run the function on each test case and print the results
	for i, testCase := range testCases {
		fmt.Printf("Test Case %d: Input: %v -> Output: %v\n", i+1, testCase, productExceptSelfDivision(testCase))
	}

	// Test the productExceptSelf function
	for i, testCase := range testCases {
		fmt.Printf("Test Case %d: Input: %v -> Output: %v\n", i+1, testCase, productExceptSelf(testCase))
	}
}
