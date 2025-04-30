// https://leetcode.com/problems/valid-parentheses/description/
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//An input string is valid if:
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Every close bracket has a corresponding open bracket of the same type.

// Example 1:
// Input: s = "{[()]}}"
// Output: false
// Example 2:
// Input: s = ""{[()]}[{}]({[]})""
// Output: true

package main

import "fmt"

func isValid(s string) bool { // Complexity O(n) and space complexity O(n)
	stack := []rune{}         // Initialize an empty stack to keep track of opening brackets
	mapping := map[rune]rune{ // Create a mapping of closing brackets to their corresponding opening brackets
		')': '(', // Closing parenthesis maps to opening parenthesis
		'}': '{', // Closing curly brace maps to opening curly brace
		']': '[', // Closing square bracket maps to opening square bracket
	}

	for _, char := range s { // Iterate through each character in the input string
		if _, ok := mapping[char]; ok { // If the character is a closing bracket
			if len(stack) == 0 || stack[len(stack)-1] != mapping[char] { // Check if the stack is empty or the top of the stack doesn't match the corresponding opening bracket
				return false // If so, return false as the string is not valid
			}
			stack = stack[:len(stack)-1] // Pop the top element from the stack (remove the last element)
		} else { // If the character is an opening bracket
			stack = append(stack, char) // Push it onto the stack
		}
	}
	return len(stack) == 0 // Return true if the stack is empty (all brackets matched), otherwise return false
}

func main() {
	// Test cases
	testCases := []struct {
		input    string
		expected bool
	}{
		{"{[()]}", true},                 // General case: valid nested brackets
		{"{[()]}}", false},               // Extra closing bracket
		{"{[()]}[{}]({[]})", true},       // Complex valid case with multiple bracket types
		{"", true},                       // Empty string: valid
		{"{[}", false},                   // Mismatched brackets
		{"{[(])}", false},                // Incorrectly nested brackets
		{"(((((((((())))))))))", true},   // Deeply nested valid parentheses
		{"(((((((((()))))))))))", false}, // Extra closing parenthesis
		{"{[()]}{[()]}{[()]}", true},     // Repeated valid patterns
		{"{[()]}{[()]}}", false},         // Valid pattern followed by an extra closing bracket
		{"{", false},                     // Single opening bracket
		{"}", false},                     // Single closing bracket
		{"{[()]}[{}]({[})", false},       // Complex invalid case
		{"{[()]}[{}]({[]}){}", true},     // Valid case with trailing valid brackets
		{"{[()]}[{}]({[]}){", false},     // Valid case with trailing unmatched opening bracket
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := isValid(testCase.input)
		fmt.Printf("Test Case %d: Input: %q, Expected: %v, Got: %v\n",
			i+1, testCase.input, testCase.expected, result)
	}
}
