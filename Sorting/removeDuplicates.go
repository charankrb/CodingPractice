// https://leetcode.com/problems/remove-duplicates-from-sorted-array

package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	targetIndx := 1
	uniqueNum := nums[0]

	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] != uniqueNum {
			uniqueNum = nums[i]
			nums[targetIndx] = nums[i]
			targetIndx++
		}
	}
	return targetIndx
}

func main() {
	array := []int{0, 0, 0, 1, 1, 1, 3, 4, 4, 4, 8, 8, 8}
	actual := removeDuplicates(array)
	fmt.Println(actual)
}
