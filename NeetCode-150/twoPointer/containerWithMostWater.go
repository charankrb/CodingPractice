//https://leetcode.com/problems/container-with-most-water/description/
// Container With Most Water
//You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
//Find two lines that together with the x-axis form a container, such that the container contains the most water.
//Return the maximum amount of water a container can store.
// Example 1:
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
// In this case, the max area of water the container can contain is 49.

package main

import "fmt"

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0 // Return 0 if the input array has less than 2 elements
	}
	maxArea := 0                    // Initialize maxArea to 0
	left, right := 0, len(height)-1 // Initialize two pointers: left and right
	for left < right {              // Loop until the two pointers meet
		// Calculate the width and height of the container
		width := right - left
		h := min(height[left], height[right])
		// Calculate the area and update maxArea if it's larger
		area := width * h
		if area > maxArea {
			maxArea = area
		}
		// Move the pointer pointing to the shorter line inward
		if height[left] < height[right] {
			left++ // Move left pointer to the right
		} else {
			right-- // Move right pointer to the left
		}
	}
	return maxArea // Return the maximum area found
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},      // General case with multiple heights
		{[]int{1, 1}, 1},                            // Smallest valid input
		{[]int{4, 3, 2, 1, 4}, 16},                  // Symmetric heights
		{[]int{1, 2, 1}, 2},                         // Heights with a peak in the middle
		{[]int{1, 2, 4, 3}, 4},                      // Heights with increasing and decreasing values
		{[]int{}, 0},                                // Empty array
		{[]int{1}, 0},                               // Single element array
		{[]int{1, 2}, 1},                            // Two elements
		{[]int{1, 8, 100, 2, 100, 4, 8, 3, 7}, 200}, // Large height difference
		{[]int{1, 2, 3, 4, 5, 25, 24, 3, 4}, 24},    // Heights with a large peak
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := maxArea(testCase.height)
		fmt.Printf("Test Case %d: Input: %v, Expected: %d, Got: %d\n",
			i+1, testCase.height, testCase.expected, result)
	}
}
