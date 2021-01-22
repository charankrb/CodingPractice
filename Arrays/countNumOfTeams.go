// https://leetcode.com/problems/count-number-of-teams/
// There are n soldiers standing in a line. Each soldier is assigned a unique rating value.
//
// You have to form a team of 3 soldiers amongst them under the following rules:
// Choose 3 soldiers with index (i, j, k) with rating (rating[i], rating[j], rating[k]).
// A team is valid if: (rating[i] < rating[j] < rating[k]) or (rating[i] > rating[j] > rating[k]) where (0 <= i < j < k < n).
// Return the number of teams you can form given the conditions. (soldiers can be part of multiple teams).
//
// Input: rating = [2,5,3,4,1]
// Output: 3
// Explanation: We can form three teams given the conditions. (2,3,4), (5,4,1), (5,3,1).

package main

import (
	"fmt"
)

func numTeams(rating []int) int {
	ans, n := 0, len(rating)
	for i := 1; i < n-1; i++ {
		leftSmall, leftLarge, rightSmall, rightLarge := 0, 0, 0, 0
		for j := i - 1; j >= 0; j-- {
			if rating[j] < rating[i] {
				leftSmall++
			} else {
				leftLarge++
			}
		}
		for j := i + 1; j < n; j++ {
			if rating[j] < rating[i] {
				rightSmall++
			} else {
				rightLarge++
			}
		}
		ans += leftSmall*rightLarge + leftLarge*rightSmall
		//fmt.Println(ans)
	}

	return ans
}

func main() {
	output := numTeams([]int{2, 5, 3, 4, 1})
	fmt.Println(output)
}
