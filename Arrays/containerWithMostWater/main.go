// https://leetcode.com/problems/container-with-most-water
//
// Given n non-negative integers a1, a2, ..., an , where each represents
// a point at coordinate (i, ai). n vertical lines are drawn such that
// the two endpoints of the line i is at (i, ai) and (i, 0). Find two
// lines, which, together with the x-axis forms a container, such that
// the container contains the most water.
//
//  Input: height = [1,8,6,2,5,4,8,3,7]
//  Output: 49
//  Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
//  In this case, the max area of water (blue section) the container can contain is 49.

package main

import (
	"fmt"
)

func maxArea(height []int) int {
	i, j, water := 0, len(height)-1, 0
	for i < j {
		water = max(water, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return water
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	array := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	actual := maxArea(array)
	fmt.Println(actual)
}
