// https://leetcode.com/problems/valid-anagram/description/
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true
// Explanation: t is an anagram of s.
// example 2:
// Input: s = "rat", t = "car"
// Output: false
// Explanation: t is not an anagram of s.

package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) { // If lengths are different, they can't be anagrams
		return false
	}
	smap := [26]int{} // Create a frequency map for characters in s
	for i := range s {
		smap[s[i]-'a']++ // Increment the count for each character in s
		smap[t[i]-'a']-- // Decrement the count for each character in t
	}

	for _, count := range smap { // Check if all counts are zero
		if count != 0 { // If any count is not zero, they aren't anagrams
			return false
		}
	}
	return true // If all counts are zero, they are anagrams
	/*
		counts := make(map[rune]int) // Create a map to count occurrences of characters

		for _, char := range s { // Count characters in the first string
			counts[char]++
		}
		for _, char := range t { // Decrease counts for characters in the second string
			counts[char]--
			if counts[char] < 0 { // If any count goes negative, they aren't anagrams
				return false
			}
		}
		return true // If all counts are zero, they are anagrams
	*/
}

func main() {
	// Test cases
	testCases := []struct {
		s, t string
	}{
		{"anagram", "nagaram"}, // Anagrams
		{"rat", "car"},         // Not anagrams
		{"", ""},               // Empty strings (anagrams)
		{"a", "a"},             // Single character (anagrams)
		{"abc", "cba"},         // Different order (anagrams)
		{"abc", "def"},         // Different characters (not anagrams)
	}

	for _, testCase := range testCases {
		result := isAnagram(testCase.s, testCase.t)
		println("Input:", testCase.s, testCase.t, "Output:", result) // Print the result for each test case
	}
}
