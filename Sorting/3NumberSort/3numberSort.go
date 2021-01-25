// function will take array of integerts and other array of distint integers
// first array will contain only elements that are present in the second array
// sort the first array in the order of second array
//
//    array := []int{1, 0, 0, -1, -1, 0, 1, 1}
//	  order := []int{0, 1, -1}
//
//    expected := []int{0, 0, 0, 1, 1, 1, -1, -1}

package main

import "fmt"

func ThreeNumberSort(nums []int, order []int) []int {
	bucketNums := []int{0, 0, 0}

	for _, num := range nums {
		index := indexOrder(order, num)
		bucketNums[index] += 1
	}
	for i := 0; i < 3; i++ {
		value := order[i]
		count := bucketNums[i]

		numElementsBefore := sumSublist(bucketNums, i)
		for n := 0; n < count; n++ {
			curInd := numElementsBefore + n
			nums[curInd] = value
		}
	}

	return nums
}

func indexOrder(order []int, target int) int {
	for index, num := range order {
		if num == target {
			return index
		}
	}
	return -1
}

func sumSublist(nums []int, index int) int {
	sum := 0
	for i, value := range nums {
		if i >= index {
			return sum
		}
		sum += value
	}
	return sum
}

func main() {
	array := []int{1, 0, 0, -1, -1, 0, 1, 1}
	order := []int{0, 1, -1}
	actual := ThreeNumberSort(array, order)
	fmt.Println(actual)
}
