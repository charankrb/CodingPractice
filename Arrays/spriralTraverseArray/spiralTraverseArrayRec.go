package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	actual := SpiralTraverse(matrix)
	fmt.Println(actual)
}

func SpiralTraverse(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}
	result := []int{}
	traverseArray(
		array,
		0,
		len(array)-1,
		0,
		len(array[0])-1,
		&result)
	return result
}

func traverseArray(array [][]int, sR, eR, sC, eC int, result *[]int) {
	if sR > eR || sC > eC {
		return
	}
	for col := sC; col <= eC; col++ {
		*result = append(*result, array[sR][col])
	}
	for row := sR + 1; row <= eR; row++ {
		*result = append(*result, array[row][eC])
	}
	for col := eC - 1; col >= sC; col-- {
		if sR == eR {
			break
		}
		*result = append(*result, array[eR][col])
	}
	for row := eR - 1; row > sR; row-- {
		if sC == eC {
			break
		}
		*result = append(*result, array[row][sC])
	}
	traverseArray(array, sR+1, eR-1, sC+1, eC-1, result)
}
