// Write a function that takes in a non empty array of integers
// and returns array where each element in an array is equal to
// product of every other element in an array

package main

import (
	"fmt"
)

func main() {

	input := []int{3, 8, 4, 2, 1, 6}
	actual := ArrayOfProducts(input)
	fmt.Println(actual)
	actual = ArrayOfProductsBrut(input)
	fmt.Println(actual)
}

// O(n) Time | O(n) Space
func ArrayOfProducts(array []int) []int {

	products := make([]int, len(array))
	leftRunningProducts := 1
	rightRunningProducts := 1

	for i := range array {
		products[i] = 1
	}

	// Calculate product of every number from left to right
	for i := range array {
		products[i] = leftRunningProducts
		leftRunningProducts *= array[i]
	}

	// Calculate product of every number from right to left
	for i := len(array) - 1; i >= 0; i-- {
		products[i] *= rightRunningProducts
		rightRunningProducts *= array[i]
	}
	return products
}

// O(n^2) Time | O(n) Space
func ArrayOfProductsBrut(array []int) []int {
	var product []int
	for i := 0; i < len(array); i++ {
		runningProduct := 1
		for j := 0; j < len(array); j++ {
			if i != j {
				runningProduct *= array[j]
			}
		}
		product = append(product, runningProduct)
	}
	return product
}
