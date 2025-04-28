// https://leetcode.com/problems/group-anagrams/description/
//Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
// For example, "anagram" is an anagram of "nagaram".
// Example 1:
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Explanation: One possible answer is [["bat"],["nat","tan"],["ate","eat","tea"]].
// Example 2:
// Input: strs = [""]
// Output: [[""]]
// Example 3:
// Input: strs = ["a"]
// Output: [["a"]]

package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	// Create a map to group anagrams. The key is the sorted string, and the value is a slice of strings that are anagrams.
	res := make(map[string][]string)

	// Iterate over each string in the input array
	for _, s := range strs {
		// Sort the string to create a key for the map
		sortedS := sortString(s)
		// Append the original string to the group corresponding to the sorted key
		res[sortedS] = append(res[sortedS], s)
		// Example after processing:
		// For "eat", sortedS = "aet", res = {"aet": ["eat"]}
		// For "tea", sortedS = "aet", res = {"aet": ["eat", "tea"]}
		// For "tan", sortedS = "ant", res = {"aet": ["eat", "tea"], "ant": ["tan"]}
		// For "ate", sortedS = "aet", res = {"aet": ["eat", "tea", "ate"], "ant": ["tan"]}
		// For "nat", sortedS = "ant", res = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"]}
		// For "bat", sortedS = "abt", res = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"], "abt": ["bat"]}
	}
	// Create a result slice to store the grouped anagrams
	var result [][]string
	// Iterate over the map and append each group of anagrams to the result
	for _, group := range res {
		result = append(result, group)
		// Example after processing:
		// result = [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
	}
	return result
}

func sortString(s string) string {
	// Create a slice of runes from the string
	runes := []rune(s)
	// Sort the slice of runes
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	// Convert the sorted slice of runes back to a string
	return string(runes)
}
func main() {
	// Test cases
	testCases := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"}, // General case with multiple groups
		{"", ""},                                 // Case with empty strings
		{"a"},                                    // Single string
		{"a", "b", "c"},                          // No anagrams, all distinct
		{"abc", "bca", "cab", "bac"},             // All strings are anagrams
		{"abc", "def", "ghi"},                    // No anagrams, all distinct groups
		{"listen", "silent", "enlist", "inlets"}, // Anagrams of the same word
		{"a", "aa", "aaa"},                       // Strings of different lengths
		{"abc", "cba", "bac", "xyz", "zyx"},      // Multiple groups with anagrams
		{"abcd", "dcba", "bcda", "dabc", "efgh"}, // Mixed anagrams and distinct strings
	}

	// Run the function on each test case and print the results
	for i, testCase := range testCases {
		fmt.Printf("Test Case %d: Input: %v\n", i+1, testCase)
		result := groupAnagrams(testCase)
		fmt.Println("Output:")
		for _, group := range result {
			fmt.Println(group)
		}
		fmt.Println()
	}
}
