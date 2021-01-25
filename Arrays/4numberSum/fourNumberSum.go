// Write a Function that takes an array of distinct elements and target sum
// which calculates the quadruplets in the array that sum up to the target sum
// and returns all the possible quadruplets in a 2d array
// Ex : array = [7, 6, 4, -1, 1, 2] , targetSum = 16
// Output : [[7, 6, 4, -1], [7, 6, 1, 2]]

package main

import (
	"fmt"
)

func FourNumberSum(array []int, target int) [][]int {
	fourNumSum := [][]int{}
	twoNumSum := map[int][][]int{}
	for i := 1; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			currentSum := array[i] + array[j]
			currDiff := target - currentSum
			if pairs, found := twoNumSum[currDiff]; found {
				for _, pair := range pairs {
					newQuad := append(pair, array[i], array[j])
					fourNumSum = append(fourNumSum, newQuad)
				}
			}
		}
		for k := 0; k < i; k++ {
			currentSum := array[i] + array[k]
			twoNumSum[currentSum] = append(twoNumSum[currentSum], []int{array[i], array[k]})
		}
	}
	return fourNumSum
}

func main() {
	output := FourNumberSum([]int{7, 6, 4, -1, 1, 2}, 16)
	fmt.Println(output)
}
