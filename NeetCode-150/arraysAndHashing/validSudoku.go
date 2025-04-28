//https://leetcode.com/problems/valid-sudoku/description/
// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

//Each row must contain the digits 1-9 without repetition.
//Each column must contain the digits 1-9 without repetition.
//Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//Note:
//A Sudoku board (partially filled) could be valid but is not necessarily solvable.
//Only the filled cells need to be validated according to the mentioned rules.

// example 1:
// Input: board = [["5","3",".",".","7",".",".",".","."],
// ["6",".",".","1","9","5",".",".","."],
// [".","9","8",".",".",".",".","6","."],
// [8",".",".",".","6",".",".",".","3"],
// ["4",".",".","8",".","3",".",".","1"],
// ["7",".",".",".","2",".",".",".","6"],
// [".","6",".",".",".",".","2","8","."],
// [".",".",".","4","1","9",".",".","5"],
// [".",".",".",".","8",".",".","7","9"]]
// Output: true
// Explanation: A partially filled Sudoku which is valid.
// example 2:
// Input: board = [["8","3",".",".","7",".",".",".","."],
// [".","9","8",".",".",".",".","6","."],
// [".",".",".","2","4","9",".",".","."],
// [8",".",".",".","6",".",".",".","3"],
// ["4",".",".","8",".","3",".",".","1"],
// ["7",".",".",".","2",".",".",".","6"],
// [".","6",".",".",".",".","2","8","."],
// [".",".",".","4","1","9",".",".","5"],
// [".",".",".",".","8",".",".","7","9"]]
// Output: false
// Explanation: Certain rows or columns have repeated digits.

package main

import "fmt"

func isValidSudoku(board [][]byte) bool { // O(n^2) time, O(1) space
	// Create a map to track the seen numbers in rows, columns, and boxes
	rows := make([]map[byte]bool, 9)  // Tracks numbers seen in each row
	cols := make([]map[byte]bool, 9)  // Tracks numbers seen in each column
	boxes := make([]map[byte]bool, 9) // Tracks numbers seen in each 3x3 box

	// Initialize the maps for rows, columns, and boxes
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	// Iterate through each cell in the board
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' { // Skip empty cells
				num := board[r][c]        // Current number
				boxIndex := (r/3)*3 + c/3 // Calculate the 3x3 box index

				// Check if the number is already seen in the row, column, or box
				if rows[r][num] || cols[c][num] || boxes[boxIndex][num] {
					return false // If duplicate is found, the board is invalid
				}

				// Mark the number as seen in the row, column, and box
				rows[r][num] = true
				cols[c][num] = true
				boxes[boxIndex][num] = true
			}
		}
	}
	return true // If no duplicates are found, the board is valid
}

/*
Dry Run Example:
Input:
board = [][]byte{
    {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
    {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
    {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
    {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
    {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
    {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
    {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
    {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
    {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
}

Execution Steps:
1. Initialize `rows`, `cols`, and `boxes` as empty maps for tracking seen numbers.

2. Process each cell in the board:
   - **Cell (0, 0)**: '5'
     - `boxIndex = (0/3)*3 + 0/3 = 0`
     - '5' is not in `rows[0]`, `cols[0]`, or `boxes[0]`.
     - Mark '5' as seen in `rows[0]`, `cols[0]`, and `boxes[0]`.

   - **Cell (0, 1)**: '3'
     - `boxIndex = (0/3)*3 + 1/3 = 0`
     - '3' is not in `rows[0]`, `cols[1]`, or `boxes[0]`.
     - Mark '3' as seen in `rows[0]`, `cols[1]`, and `boxes[0]`.

   - **Cell (0, 4)**: '7'
     - `boxIndex = (0/3)*3 + 4/3 = 1`
     - '7' is not in `rows[0]`, `cols[4]`, or `boxes[1]`.
     - Mark '7' as seen in `rows[0]`, `cols[4]`, and `boxes[1]`.

   - Continue this process for all cells in the board.

3. Validation:
   - For each cell, check if the number is already seen in the corresponding row, column, or box.
   - If a duplicate is found, return `false`.
   - If no duplicates are found after processing all cells, return `true`.

Output:
For the given input, the function will return:
true
*/

func main() {
	// Test cases
	testCases := []struct {
		board   [][]byte
		isValid bool
	}{
		// Example 1: Valid Sudoku
		{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			isValid: true,
		},
		// Example 2: Invalid Sudoku (duplicate in row)
		{
			board: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			isValid: false,
		},
		// Example 3: Invalid Sudoku (duplicate in column)
		{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'5', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			isValid: false,
		},
		// Example 4: Invalid Sudoku (duplicate in 3x3 box)
		{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '8'},
			},
			isValid: false,
		},
		// Example 5: Empty board (valid)
		{
			board: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			},
			isValid: true,
		},
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := isValidSudoku(testCase.board)
		fmt.Printf("Test Case %d: Expected: %v, Got: %v\n", i+1, testCase.isValid, result)
	}
}
