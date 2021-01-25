/*
 * https://leetcode.com/problems/count-good-triplets/
 *
 * Given an array of integers arr, and three integers a, b and c. You need to find the number of good triplets.
 *
 * A triplet (arr[i], arr[j], arr[k]) is good if the following conditions are true:
 *
 * 0 <= i < j < k < arr.length
 * |arr[i] - arr[j]| <= a
 * |arr[j] - arr[k]| <= b
 * |arr[i] - arr[k]| <= c
 * Where |x| denotes the absolute value of x.
 *
 * Return the number of good triplets.
 *
 * Example 1:
 *
 * Input: arr = [3,0,1,1,9,7], a = 7, b = 2, c = 3
 * Output: 4
 * Explanation: There are 4 good triplets: [(3,0,1), (3,0,1), (3,1,1), (0,1,1)].
 * Example 2:
 *
 * Input: arr = [1,1,2,2,3], a = 0, b = 0, c = 1
 * Output: 0
 * Explanation: No triplet satisfies all conditions.
 */

package main

import (
	"fmt"
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	if len(arr) < 3 {
		return 0
	}
	n := len(arr)
	goodTriplets := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			if abs(arr[i]-arr[j]) <= a {
				for k := j + 1; k < n; k++ {
					if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
						goodTriplets++
					}
				}
			}
		}
	}
	return goodTriplets
}

func main() {
	arr := []int{7, 3, 7, 3, 12, 1, 12, 2, 3}
	output := countGoodTriplets(arr, 5, 8, 1)
	fmt.Println(output)
}
