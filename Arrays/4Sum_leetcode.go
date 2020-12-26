// neds improvement
package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	fourSum := [][]int{}
	if len(nums) < 4 {
		return fourSum
	}
	sort.Ints(nums)
	n := len(nums) - 1
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if (nums[left] == nums[left-1]) && (left < right) {
					left++
				}
				if (nums[right] == nums[right-1]) && (left < right) {
					right--
				}
				curSum := nums[left] + nums[right] + nums[i] + nums[j]
				if curSum > target {
					right--
				} else if curSum < target {
					left++
				} else {
					fourSum = append(fourSum, []int{nums[i], nums[j], nums[left], nums[right]})
					break
				}
			}
		}

	}
	return fourSum
}

func main() {
	output := fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0)
	fmt.Println(output)
}
