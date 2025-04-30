//https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
// Evaluate Reverse Polish Notation
// You are given an array of strings tokens that represents an expression in a Reverse Polish Notation (RPN).
// Evaluate the expression and return an integer representing the value of that expression.
// Note that:
// The valid operators are '+', '-', '*', and '/'.
// Each operand may be an integer or another expression.
// example 1:
// Input: tokens = ["2","1","+","3","*"]
// Output: 9
// Explanation: ((2 + 1) * 3) = 9
// example 2:
// Input: tokens = ["4","13","5","/","+"]
// Output: 6
// Explanation: (4 + (13 / 5)) = 6

package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens { // Iterate through each token in the input array
		// Check if the token is an operator
		switch token {
		case "+": // If the token is "+", pop the top two elements from the stack, add them, and push the result back onto the stack
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			// Pop the top two elements from the stack
			stack = stack[:len(stack)-2]
			// Add them and push the result back onto the stack
			stack = append(stack, n1+n2)
		case "-": // If the token is "-", pop the top two elements from the stack, subtract them, and push the result back onto the stack
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			// Pop the top two elements from the stack
			stack = stack[:len(stack)-2]
			// Subtract them and push the result back onto the stack
			stack = append(stack, n1-n2)
		case "*":
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, n1*n2)
		case "/":
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, n1/n2)
		default: // If the token is not an operator, it must be a number, so convert it to an integer and push it onto the stack
			// if number, convert it to an integer and push it onto the stack
			//"2" -> 2
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func main() {
	// Test cases
	testCases := []struct {
		tokens   []string
		expected int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},                 // General case: ((2 + 1) * 3) = 9
		{[]string{"4", "13", "5", "/", "+"}, 6},                // General case: (4 + (13 / 5)) = 6
		{[]string{"10", "6", "9", "3", "/", "-", "*"}, 30},     // Complex case: ((10 * (6 - (9 / 3))) = 30
		{[]string{"2", "3", "+"}, 5},                           // Simple addition: 2 + 3 = 5
		{[]string{"5", "1", "-"}, 4},                           // Simple subtraction: 5 - 1 = 4
		{[]string{"3", "4", "*"}, 12},                          // Simple multiplication: 3 * 4 = 12
		{[]string{"8", "4", "/"}, 2},                           // Simple division: 8 / 4 = 2
		{[]string{"2", "3", "4", "*", "+"}, 14},                // Mixed operations: 2 + (3 * 4) = 14
		{[]string{"4", "2", "/", "3", "-"}, -1},                // Mixed operations: (4 / 2) - 3 = -1
		{[]string{"100", "200", "+", "2", "/", "5", "*"}, 750}, // Large numbers: (((100 + 200) / 2) * 5) = 750
		{[]string{"3", "-4", "*"}, -12},                        // Negative numbers: 3 * -4 = -12
		{[]string{"-3", "-4", "+"}, -7},                        // Negative addition: -3 + -4 = -7
		{[]string{"-10", "2", "/"}, -5},                        // Negative division: -10 / 2 = -5
		{[]string{"0", "1", "+"}, 1},                           // Zero addition: 0 + 1 = 1
		{[]string{"0", "1", "-"}, -1},                          // Zero subtraction: 0 - 1 = -1
		{[]string{"0", "1", "*"}, 0},                           // Zero multiplication: 0 * 1 = 0
		{[]string{"1"}, 1},                                     // Single number: 1
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := evalRPN(testCase.tokens)
		fmt.Printf("Test Case %d: Input: %v, Expected: %d, Got: %d\n",
			i+1, testCase.tokens, testCase.expected, result)
	}
}
