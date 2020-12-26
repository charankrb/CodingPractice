// https://leetcode.com/problems/longest-consecutive-sequence/solution/
// Given an unsorted array of integers nums, return the length of the
// longest consecutive elements sequence.
// EX: Input: nums = [100,4,200,1,3,2]
// Output: [1,4],4
// Explanation: The longest consecutive elements sequence is [1,4] which is [1, 2, 3, 4] . Therefore its length is 4.

package main

import (
	"fmt"
)

func LargestRange(array []int) ([]int, int) {
	longest := []int{}
	longestLength := 0
	visited := map[int]bool{}
	for _, num := range array {
		visited[num] = true
	}
	for _, num := range array {
		if !visited[num] {
			continue
		}
		visited[num] = false
		currentLength, left, right := 1, num-1, num+1
		for visited[left] {
			visited[left] = false
			currentLength++
			left--
		}
		for visited[right] {
			visited[right] = false
			currentLength++
			right++
		}
		if currentLength > longestLength {
			longestLength = currentLength
			longest = []int{left + 1, right - 1}
		}
	}
	return longest, longestLength
}

func main() {
	output, longestLength := LargestRange([]int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6})
	fmt.Println(output, longestLength)
}
