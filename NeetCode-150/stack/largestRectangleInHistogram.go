//https://leetcode.com/problems/largest-rectangle-in-histogram/description/
//Largest Rectangle in Histogram
// Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.
// Example 1:
// Input: heights = [2,1,5,6,2,3]
// Output: 10
// Explanation: The largest rectangle has an area of 10, formed by heights[2] = 5 and heights[3] = 6.
// Example 2:
// Input: heights = [2,4]
// Output: 4
// Explanation: The largest rectangle has an area of 4, formed by heights[0] = 2 and heights[1] = 4.

package main

func largestRectangleAreaBruteForce(heights []int) int { //complexity O(n^2)
	n := len(heights)
	maxArea := 0

	for i := 0; i < n; i++ {
		height := heights[i]

		rightMost := i + 1
		for rightMost < n && heights[rightMost] >= height {
			rightMost++
		}

		leftMost := i
		for leftMost >= 0 && heights[leftMost] >= height {
			leftMost--
		}

		rightMost--
		leftMost++

		area := height * (rightMost - leftMost + 1)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func largestRectangleArea(heights []int) int { //complexity O(n)

	// Initialize a stack to store indices and a variable to track the maximum area
	stack := []int{}
	maxArea := 0
	//[2,1,5,6,7,2,3,0] // Add a sentinel value to handle remaining bars in the stack
	// Iterate through the heights array
	for i := 0; i <= len(heights); i++ {
		h := 0
		if i < len(heights) {
			h = heights[i]
		}

		// While the stack is not empty and the current height is less than the height of the bar
		// at the index stored at the top of the stack
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			// Pop the top index from the stack
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Calculate the width of the rectangle
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			// Calculate the area of the rectangle and update maxArea
			area := heights[top] * width
			if area > maxArea {
				maxArea = area
			}
		}

		// Push the current index onto the stack
		stack = append(stack, i)
	}

	return maxArea
}

// Dry Run Table for Input: [2, 1, 5, 6, 7, 2, 3]
/*
Index | Height (h) | Stack         | Popped | Width | Area | Max Area
---------------------------------------------------------------------
0     | 2          | [0]           | -      | -     | -    | 0
1     | 1          | [1]           | 0      | 1     | 2    | 2
2     | 5          | [1, 2]        | -      | -     | -    | 2
3     | 6          | [1, 2, 3]     | -      | -     | -    | 2
4     | 7          | [1, 2, 3, 4]  | -      | -     | -    | 2
5     | 2          | [1, 5]        | 4      | 1     | 7    | 7
      |            |               | 3      | 2     | 12   | 12
      |            |               | 2      | 3     | 15   | 15
6     | 3          | [1, 5, 6]     | -      | -     | -    | 15
7     | 0 (sentinel)| [7]          | 6      | 1     | 3    | 15
      |            |               | 5      | 5     | 10   | 15
      |            |               | 1      | 7     | 7    | 15
*/
