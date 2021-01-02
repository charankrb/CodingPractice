/*
https://leetcode.com/problems/high-five/

Given a list of the scores of different students, items, where items[i] = [IDi, scorei]
represents one score from a student with IDi, calculate each student's top five average.

Return the answer as an array of pairs result, where result[j] = [IDj, topFiveAveragej]
represents the student with IDj and their top five average. Sort result by IDj in
increasing order.

A student's top five average is calculated by taking the sum of their top five scores
and dividing it by 5 using integer division.


Example 1:

Input: items = [[1,91],[1,92],[2,93],[2,97],[1,60],[2,77],[1,65],[1,87],[1,100],[2,100],[2,76]]
Output: [[1,87],[2,88]]
Explanation:
The student with ID = 1 got scores 91, 92, 60, 65, 87, and 100. Their top five average is (100 + 92 + 91 + 87 + 65) / 5 = 87.
The student with ID = 2 got scores 93, 97, 77, 100, and 76. Their top five average is (100 + 97 + 93 + 77 + 76) / 5 = 88.6, but with integer division their average converts to 88.
Example 2:

Input: items = [[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100],[1,100],[7,100]]
Output: [[1,100],[7,100]]
*/

package main

import (
	"fmt"
	"sort"
)

func highFive(items [][]int) [][]int {
	if len(items) < 5 {
		return [][]int{}
	}
	studentScores := map[int][]int{}
	result := [][]int{}
	for _, item := range items {
		studentScores[item[0]] = append(studentScores[item[0]], item[1])
	}
	for v, score := range studentScores {
		result = append(result, []int{v, average(topK(score, 5))})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}

func average(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum / len(arr)
}

func topK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[len(arr)-k : len(arr)]
}

func main() {

	output := highFive([][]int{
		{1, 84},
		{1, 72},
		{1, 47},
		{1, 43},
		{1, 78},
		{2, 79},
		{2, 4},
		{2, 23},
		{2, 88},
		{2, 79},
		{3, 75},
		{3, 80},
		{3, 38},
		{3, 73},
		{3, 4}})
	fmt.Println(output)
}
