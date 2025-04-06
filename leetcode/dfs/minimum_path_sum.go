// https://leetcode.com/problems/minimum-path-sum/description/
package main

import (
	"fmt"
	"math"
)

// need to add memo
func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	minSum, _ := pathDfs(0, 0, row, col, grid, []int{})
	return minSum
}

func pathDfs(i, j, row, col int, grid [][]int, path []int) (int, []int) {
	// Out of bounds
	if i >= row || j >= col {
		return math.MaxInt32, nil
	}

	// Add current value to path
	path = append(path, grid[i][j])

	// Base case: bottom-right
	if i == row-1 && j == col-1 {
		sum := sumPath(path)
		fmt.Printf("Valid path: %v -> sum: %d\n", path, sum)
		return sum, append([]int{}, path...) // Return a copy of the path
	}

	// Recursive calls
	downSum, downPath := pathDfs(i, j+1, row, col, grid, path)
	rightSum, rightPath := pathDfs(i+1, j, row, col, grid, path)

	if downSum < rightSum {
		return downSum, downPath
	}
	return rightSum, rightPath
}

func sumPath(path []int) int {
	sum := 0
	for _, val := range path {
		sum += val
	}
	return sum
}

func main() {
	// 1 1 4 2 1
	// 1 1 5 2 1
	// 1 1 5 1 1
	// 1 3 5 2 1
	// 1 3 5 1 1
	// 1 3 1 1 1
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
