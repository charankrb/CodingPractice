//https://leetcode.com/problems/generate-parentheses/description/
//Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//Example 1:
//Input: n = 3
//Output: ["((()))","(()())","(())()","()(())","()()()"]
//Example 2:
//Input: n = 1
//Output: ["()"]

package main

import (
	"fmt"
	"reflect"
	"strings"
)

func generateParenthesisBF(n int) []string {
	res := []string{} // Initialize a slice to store all valid combinations of parentheses
	// Define a helper function `def` to generate all possible combinations of parentheses
	var def func(string)
	def = func(s string) {
		// Base case: If the length of the current string is equal to 2*n (n pairs of parentheses)
		if len(s) == (n * 2) {
			// Check if the current string is valid using the `isValid` function
			if isValid(s) {
				res = append(res, s) // Add the valid combination to the result
			}
			return
		}
		// Add '(' to the current string and recurse
		def(s + "(")
		// Add ')' to the current string and recurse
		def(s + ")")
	}
	// Start the recursive process with an empty string
	def("")
	// Return the result containing all valid combinations
	return res
}

func isValid(s string) bool {
	open := 0 // Counter to track the balance of parentheses
	// Iterate through each character in the string
	for _, c := range s {
		if c == '(' {
			open++ // Increment the counter for an opening parenthesis
		} else {
			open-- // Decrement the counter for a closing parenthesis
		}

		// If at any point the counter becomes negative, the string is invalid
		if open < 0 {
			return false
		}
	}
	// The string is valid if all opening parentheses are matched with closing ones
	return open == 0
}

func generateParenthesisBT(n int) []string { // Complexity O(2^n) time and O(n) space
	// Initialize a stack to build the current combination of parentheses
	stack := make([]string, 0)
	// Initialize a slice to store all valid combinations of parentheses
	res := make([]string, 0)
	// Define a recursive function to perform Depth-First Search (DFS)
	var dfs func(int, int)
	dfs = func(openN, closeN int) {
		// Base case: If the number of opening and closing parentheses equals `n`
		if openN == n && closeN == n {
			// Join the stack into a string and add it to the result
			res = append(res, strings.Join(stack, ""))
			return
		}

		// Add an opening parenthesis '(' if the number of '(' used is less than `n`
		if openN < n {
			stack = append(stack, "(")   // Add '(' to the stack
			dfs(openN+1, closeN)         // Recurse with one more '('
			stack = stack[:len(stack)-1] // Backtrack by removing the last added '('
		}

		// Add a closing parenthesis ')' if the number of ')' used is less than the number of '('
		if closeN < openN {
			stack = append(stack, ")")   // Add ')' to the stack
			dfs(openN, closeN+1)         // Recurse with one more ')'
			stack = stack[:len(stack)-1] // Backtrack by removing the last added ')'
		}
	}

	// Start the DFS process with 0 opening and 0 closing parentheses
	dfs(0, 0)

	// Return the result containing all valid combinations
	return res
}

func main() {
	// Test cases
	testCases := []struct {
		n        int
		expected []string
	}{
		{1, []string{"()"}},           // Smallest case: 1 pair of parentheses
		{2, []string{"(())", "()()"}}, // 2 pairs of parentheses
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}}, // 3 pairs of parentheses
		{0, []string{""}}, // Edge case: 0 pairs of parentheses
		{4, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}}, // 4 pairs of parentheses
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := generateParenthesisBT(testCase.n)
		fmt.Printf("Test Case %d: Input: %d\n", i+1, testCase.n)
		fmt.Printf("Expected: %v\n", testCase.expected)
		fmt.Printf("Got: %v\n", result)

		// Validate the result
		if reflect.DeepEqual(result, testCase.expected) {
			fmt.Printf("Test Case %d Passed\n\n", i+1)
		} else {
			fmt.Printf("Test Case %d Failed\n\n", i+1)
		}
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := generateParenthesisBF(testCase.n)
		fmt.Printf("Test Case %d: Input: %d\n", i+1, testCase.n)
		fmt.Printf("Expected: %v\n", testCase.expected)
		fmt.Printf("Got: %v\n", result)

		// Validate the result
		if reflect.DeepEqual(result, testCase.expected) {
			fmt.Printf("Test Case %d Passed\n\n", i+1)
		} else {
			fmt.Printf("Test Case %d Failed\n\n", i+1)
		}
	}
}
