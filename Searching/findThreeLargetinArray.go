// Find three larget numbers in an array
// write a finction that takes in array and return array of three larget numbers
//
// Input  : {71, 1, 17, -7, -17, -27, 12, 200, 8, 5, 70}
// Output : {70, 71, 200}

package main

import (
	"fmt"
	"math"
)

func FindThreeLargestNumbers(array []int) []int {
	threeLargest := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for _, num := range array {
		updateLargest(threeLargest, num)
	}
	return threeLargest
}

func updateLargest(three []int, num int) {
	if num > three[2] {
		shiftUpdate(three, num, 2)
	} else if num > three[1] {
		shiftUpdate(three, num, 1)
	} else if num > three[0] {
		shiftUpdate(three, num, 0)
	}
}

func shiftUpdate(three []int, num, pos int) {
	for i := 0; i <= pos; i++ {
		if i == pos {
			three[i] = num
		} else {
			three[i] = three[i+1]
		}
	}
}

func main() {
	array := []int{71, 1, 17, -7, -17, -27, 12, 200, 8, 5, 70}
	actual := FindThreeLargestNumbers(array)
	fmt.Println(actual)
}
