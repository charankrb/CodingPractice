//https://leetcode.com/problems/queens-that-can-attack-the-king/submissions/

package main

import (
	"fmt"
)

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	result := [][]int{}
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	q := [8][8]bool{}
	for _, queen := range queens {
		q[queen[0]][queen[1]] = true
	}
	for i := 0; i < 8; i++ {
		x := king[0]
		y := king[1]
		for x >= 0 && x < 8 && y >= 0 && y < 8 {
			if q[x][y] {
				result = append(result, []int{x, y})
				break
			}
			x += dir[i][0]
			y += dir[i][1]
		}
	}
	return result
}

func main() {
	output := queensAttacktheKing([][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}, []int{0, 0})
	fmt.Println(output)
}
