//https://leetcode.com/problems/valid-palindrome/description/
//A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
//Given a string s, return true if it is a palindrome, or false otherwise.

//Example 1:
//Input: s = "A man, a plan, a canal: Panama"
//Output: true
//Explanation : "amanaplanacanalpanama" is a palindrome.
//Example 2:
//Input: s = "race a car"
//Output: false
//Explanation: "raceacar" is not a palindrome.

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s) // Convert the string to lowercase
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// Move the left pointer to the next alphanumeric character
		for i < j && !isAlphanumeric(s[i]) {
			i++
		}
		// Move the right pointer to the previous alphanumeric character
		for i < j && !isAlphanumeric(s[j]) {
			j--
		}
		if s[i] != s[j] { // Compare the characters at the two pointers
			return false // If they are not equal, it's not a palindrome
		}
	}
	return true // If all characters matched, it's a palindrome
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') // Check if the character is alphanumeric
}

func main() {
	// Test cases
	testCases := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true}, // General case: valid palindrome with spaces and punctuation
		{"race a car", false},                    // General case: not a palindrome
		{"", true},                               // Empty string: valid palindrome
		{" ", true},                              // Single space: valid palindrome
		{"a", true},                              // Single character: valid palindrome
		{"aa", true},                             // Two identical characters: valid palindrome
		{"ab", false},                            // Two different characters: not a palindrome
		{"0P", false},                            // Mixed alphanumeric, not a palindrome
		{"12321", true},                          // Numeric palindrome
		{"12345", false},                         // Numeric non-palindrome
		{"Able was I, ere I saw Elba", true},     // Valid palindrome with mixed case and spaces
		{"No lemon, no melon", true},             // Valid palindrome with spaces and punctuation
		{"Was it a car or a cat I saw?", true},   // Valid palindrome with punctuation
		{"Madam, in Eden, I'm Adam", true},       // Valid palindrome with punctuation and mixed case
		{"!@#$%^&*()", true},                     // Only special characters: valid palindrome
		{"a.b,a", true},                          // Valid palindrome with punctuation
		{"a.b,c", false},                         // Not a palindrome with punctuation
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := isPalindrome(testCase.input)
		fmt.Printf("Test Case %d: Input: %q, Expected: %v, Got: %v\n", i+1, testCase.input, testCase.expected, result)
	}
}
