//https://leetcode.com/problems/trapping-rain-water/description/
// Trapping Rain Water
// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
// Example 1:
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6

package main

import "fmt"

func trapTwoPointer(height []int) int { // complexity O(n) and space complexity O(1)
	if len(height) < 3 {
		return 0 // Return 0 if the input array has less than 3 elements
	}
	left, right := 0, len(height)-1                  // Initialize two pointers: left and right
	leftMax, rightMax := height[left], height[right] // Initialize leftMax and rightMax to the first and last elements
	trappedWater := 0                                // Initialize trappedWater to 0

	for left < right { // Loop until the two pointers meet
		if leftMax < rightMax { // If the left max is less than the right max
			left++                      // Move the left pointer to the right
			if height[left] < leftMax { // If the current height is less than leftMax
				trappedWater += leftMax - height[left] // Calculate trapped water at this position
			} else {
				leftMax = height[left] // Update leftMax if current height is greater
			}
		} else { // If the right max is less than or equal to the left max
			right--                       // Move the right pointer to the left
			if height[right] < rightMax { // If the current height is less than rightMax
				trappedWater += rightMax - height[right] // Calculate trapped water at this position
			} else {
				rightMax = height[right] // Update rightMax if current height is greater
			}
		}
	}
	return trappedWater // Return the total trapped water calculated
}

func trapSuffixandPrifixSum(height []int) int { // complexity O(n) and space complexity O(n)
	if len(height) < 3 {
		return 0 // Return 0 if the input array has less than 3 elements
	}
	n := len(height)
	leftMax := make([]int, n)  // Create a slice to store left max heights
	rightMax := make([]int, n) // Create a slice to store right max heights

	leftMax[0] = height[0]   // Initialize the first element of leftMax
	for i := 1; i < n; i++ { // Fill the leftMax array
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]   // Initialize the last element of rightMax
	for i := n - 2; i >= 0; i-- { // Fill the rightMax array
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	trappedWater := 0        // Initialize trappedWater to 0
	for i := 0; i < n; i++ { // Calculate trapped water at each position
		trappedWater += min(leftMax[i], rightMax[i]) - height[i]
	}
	return trappedWater // Return the total trapped water calculated
}

func main() {
	// Test cases
	testCases := []struct {
		height   []int
		expected int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6}, // General case with multiple heights
		{[]int{4, 2, 0, 3, 2, 5}, 9},                   // Heights with valleys and peaks
		{[]int{1, 0, 2}, 1},                            // Smallest valid input with a valley
		{[]int{3, 0, 0, 2, 0, 4}, 10},                  // Heights with multiple valleys
		{[]int{}, 0},                                   // Empty array
		{[]int{1}, 0},                                  // Single element array
	}

	for i, testCase := range testCases {
		result := trapSuffixandPrifixSum(testCase.height)
		fmt.Printf("Test Case %d: Input: %v, Expected: %d, Got: %d\n", i+1, testCase.height, testCase.expected, result)
	}

	for i, testCase := range testCases {
		result := trapTwoPointer(testCase.height)
		fmt.Printf("Test Case %d: Input: %v, Expected: %d, Got: %d\n", i+1, testCase.height, testCase.expected, result)
	}
}
