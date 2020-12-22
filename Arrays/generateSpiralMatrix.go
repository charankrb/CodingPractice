// Generate a n * n 2D spiral matrix
// Ex: Input 3
//     Output =  [
//		  [1 2 3]
//		  [8 9 4]
//		  [7 6 5]
//		]

package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, 3)
	}
	fmt.Println(len(result))
	sC, eC := 0, n-1
	sR, eR := 0, n-1
	num := 1
	for sC <= eC && sR <= eR {
		for col := sC; col <= eC; col++ {
			result[sR][col] = num
			num++
		}
		for row := sR + 1; row <= eR; row++ {
			result[row][eC] = num
			num++
		}
		for col := eC - 1; col >= sC; col-- {
			if sR == eR {
				break
			}
			result[eR][col] = num
			num++
		}
		for row := eR - 1; row > sR; row-- {
			if sC == eC {
				break
			}
			result[row][sC] = num
			num++
		}
		sC++
		sR++
		eC--
		eR--
	}
	return result
}

func main() {
	actual := generateMatrix(3)
	fmt.Println(actual)
}
