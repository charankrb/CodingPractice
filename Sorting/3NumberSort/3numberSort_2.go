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

func ThreeNumberSort(array []int, order []int) []int {
	firstNum := order[0]
	thirdNum := order[2]
	firstInd := 0

	for index, num := range array {
		if num == firstNum {
			array[index], array[firstInd] = array[firstInd], array[index]
			firstInd++
		}
	}

	thirdInd := len(array) - 1
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == thirdNum {
			array[i], array[thirdInd] = array[thirdInd], array[i]
			thirdInd--
		}
	}
	return array
}

func main() {
	array := []int{1, 0, 0, -1, -1, 0, 1, 1}
	order := []int{0, 1, -1}
	actual := ThreeNumberSort(array, order)
	fmt.Println(actual)
}
