//https://leetcode.com/problems/search-a-2d-matrix/description/
// Search a 2D Matrix
// Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
// Integers in each row are sorted from left to right.
// Integers in each column are sorted from top to bottom.
// example 1:
// Input: matrix = [[1,3,5],[7,9,11],[12,13,15]], target = 9
// Output: true
// Example 2:
// Input: matrix = [[1,3,5],[7,9,11],[12,13,15]], target = 10
// Output: false

package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows*cols-1

	for left <= right {
		mid := left + (right-left)/2
		midValue := matrix[mid/cols][mid%cols]

		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	// Test cases
	testCases := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		// General cases
		{[][]int{{1, 3, 5}, {7, 9, 11}, {12, 13, 15}}, 9, true},   // Target exists in the matrix
		{[][]int{{1, 3, 5}, {7, 9, 11}, {12, 13, 15}}, 10, false}, // Target does not exist in the matrix

		// Edge cases
		{[][]int{}, 5, false},         // Empty matrix
		{[][]int{{}}, 5, false},       // Matrix with empty rows
		{[][]int{{1}}, 1, true},       // Single element matrix, target exists
		{[][]int{{1}}, 2, false},      // Single element matrix, target does not exist
		{[][]int{{1, 3, 5}}, 3, true}, // Single row matrix, target exists
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := searchMatrix(testCase.matrix, testCase.target)
		fmt.Printf("Test Case %d: Matrix: %v, Target: %d\n", i+1, testCase.matrix, testCase.target)
		fmt.Printf("Expected: %v, Got: %v\n", testCase.expected, result)

		// Validate the result
		if result == testCase.expected {
			fmt.Printf("Test Case %d Passed\n\n", i+1)
		} else {
			fmt.Printf("Test Case %d Failed\n\n", i+1)
		}
	}
}
