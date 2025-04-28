// https://leetcode.com/problems/spiral-matrix/description/
// example 1:
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]
// Explanation: The elements in spiral order are [1,2,3,6,9,8,7,4,5].
// example 2:
// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]
// Explanation: The elements in spiral order are [1,2,3,4,8,12,11,10,9,5,6,7].

package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{} // Return an empty slice if the input matrix is empty
	}
	result := []int{}                                                 // Initialize an empty slice to store the result
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1 // Define boundaries

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ { // Traverse from left to right
			result = append(result, matrix[top][i])
		}
		top++ // Move the top boundary down

		for i := top; i <= bottom; i++ { // Traverse from top to bottom
			result = append(result, matrix[i][right])
		}
		right-- // Move the right boundary left

		if top <= bottom {
			for i := right; i >= left; i-- { // Traverse from right to left
				result = append(result, matrix[bottom][i])
			}
			bottom-- // Move the bottom boundary up
		}

		if left <= right {
			for i := bottom; i >= top; i-- { // Traverse from bottom to top
				result = append(result, matrix[i][left])
			}
			left++ // Move the left boundary right
		}
	}
	return result // Return the result slice containing elements in spiral order
}
func main() {
	// Test cases
	testCases := [][][]int{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},             // General case
		{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, // Rectangular matrix
		{{1}},            // Single element
		{},               // Empty matrix
		{{1, 2}, {3, 4}}, // Square matrix with two rows and two columns
	}

	for _, testCase := range testCases {
		result := spiralOrder(testCase)
		fmt.Println(result) // Print the result for each test case
	}
}
