// https://leetcode.com/problems/top-k-frequent-elements/description/
// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Explanation: "1" and "2" are the two most frequent elements.
// Example 2:
// Input: nums = [1], k = 1
// Output: [1]
// Explanation: "1" is the only element in the array

package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int) // Create a map to store the frequency of each number
	for _, num := range nums {
		frequency[num]++ // Increment the frequency count for each number
	}

	// Create a slice of pairs (number, frequency) and sort it by frequency
	type pair struct {
		num  int
		freq int
	}
	pairs := make([]pair, 0, len(frequency))
	for num, freq := range frequency {
		pairs = append(pairs, pair{num, freq})
	}

	// Ensure k is valid
	if k <= 0 || k > len(frequency) {
		return []int{} // Return an empty result if k is invalid
	}
	// Sort the pairs by frequency in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	// Extract the top k frequent elements
	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].num)
	}
	return result
}

func topKFrequentBucketSort(nums []int, k int) []int {
	// Example: nums = [1, 1, 1, 2, 2, 3], k = 2
	// Step 1: Create a map to count the frequency of each number
	// count = {1: 3, 2: 2, 3: 1}
	count := make(map[int]int)

	// Step 2: Create a slice of slices to act as buckets, where the index represents the frequency
	// freq will have a size of len(nums)+1 (7 in this case)
	// freq = [[], [], [], [], [], [], []]
	freq := make([][]int, len(nums)+1)

	// Step 3: Populate the frequency map
	for _, num := range nums {
		count[num]++ // Increment the count for each number
		// After processing nums, count = {1: 3, 2: 2, 3: 1}
	}

	// Step 4: Populate the buckets based on frequency
	for num, cnt := range count {
		freq[cnt] = append(freq[cnt], num) // Append the number to the bucket corresponding to its frequency
		// After processing count, freq = [[], [3], [2], [1], [], [], []]
	}

	// Step 5: Result slice to store the top k frequent elements
	res := []int{}

	// Step 6: Iterate over the buckets in reverse order (from highest frequency to lowest)
	// Start from the highest frequency (len(freq)-1) and move down
	for i := len(freq) - 1; i > 0; i-- {
		for _, num := range freq[i] { // Iterate over all numbers in the current bucket
			res = append(res, num) // Add the number to the result
			// Example: Add 1 (from freq[3]), then add 2 (from freq[2])
			if len(res) == k { // If we have collected k elements, return the result
				return res
			}
		}
	}

	// Step 7: Return the result (in case k is not reached, though this shouldn't happen with valid input)
	return res
}

func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2},             // General case with multiple frequent elements
		{[]int{1}, 1},                            // Single element array
		{[]int{1, 2, 3, 4, 5}, 3},                // All elements have the same frequency
		{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 5}, // All elements have equal frequency, k equals the array size
		{[]int{1, 1, 1, 2, 2, 3, 3, 3, 3}, 1},    // One element is significantly more frequent
		{[]int{}, 0},                             // Empty array
		{[]int{1, 2, 3, 4, 5}, 0},                // k is 0
		{[]int{1, 1, 2, 2, 3, 3}, 6},             // k equals the number of unique elements
		{[]int{1, 1, 1, 1, 1}, 1},                // All elements are the same
		{[]int{-1, -1, -2, -2, -3, -3}, 2},       // Negative numbers
		{[]int{1, 2, 3, 1, 2, 3, 1, 2, 3}, 3},    // All elements have equal frequency
	}

	// Run the function on each test case and print the results
	for i, testCase := range testCases {
		fmt.Printf("Test Case %d: Input: nums = %v, k = %d -> Output: %v\n", i+1, testCase.nums, testCase.k, topKFrequent(testCase.nums, testCase.k))
	}

	for i, testCase := range testCases {
		fmt.Printf("Test Case %d: [Bucket Sort]Input: nums = %v, k = %d -> Output : %v\n", i+1, testCase.nums, testCase.k, topKFrequentBucketSort(testCase.nums, testCase.k))
	}
}
