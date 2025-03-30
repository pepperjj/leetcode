// https://leetcode.com/problems/number-of-islands/?envType=problem-list-v2&envId=oizxjoit&
package main

import "fmt"

func numIslands(grid [][]byte) int {
	res := 0
	row, col := len(grid), len(grid[0])
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if visited[i][j] || grid[i][j] == '0' {
				continue
			}
			dfs(row, col, grid, visited, i, j)
			res++
		}
	}
	return res
}

func dfs(row, col int, grid [][]byte, visited [][]bool, i, j int) {
	if i < 0 || j < 0 || i >= row || j >= col || grid[i][j] == '0' || visited[i][j] {
		return
	}
	visited[i][j] = true
	dfs(row, col, grid, visited, i+1, j)
	dfs(row, col, grid, visited, i, j+1)
	dfs(row, col, grid, visited, i-1, j)
	dfs(row, col, grid, visited, i, j-1)
}

func main() {
	input := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(input))
}
