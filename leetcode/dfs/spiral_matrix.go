// https://leetcode.com/problems/spiral-matrix/description/
package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	directions := [][]int{
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
		{-1, 0}, // up
	}

	rows, cols := len(matrix), len(matrix[0])
	res := make([]int, 0)
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	// for each direction
	// if the pos never visit & i, j < row, col, ++
	// if reach change to the next direction
	c := 0
	i, j := 0, 0
	d := directions[c]
	for len(res) < rows*cols {
		// dIndex
		// i, j = 0, i + d[0], j + d[1]; i, j < row, col
		// c + 1
		for i < rows && j < cols && i >= 0 && j >= 0 {
			if visited[i][j] {
				break
			}
			res = append(res, matrix[i][j])
			visited[i][j] = true
			i = i + d[0]
			j = j + d[1]
		}
		c = c + 1
		i = i - d[0]
		j = j - d[1]
		dIndex := c % 4
		d = directions[dIndex]
		i = i + d[0]
		j = j + d[1]
	}

	return res
}

func main() {
	matrix := [][]int{
		{1, 2},
		{4, 5},
		{7, 8},
	}
	fmt.Println(spiralOrder(matrix))
}
