package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// sort by start[i]
	// if intervals[m][end] >= intervals[m+1][start]
	// res = append(res, [intervals[m][start], intervals[m+1][end]])
	// return result

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); {
		j := i + 1
		if j < len(intervals) && intervals[i][1] >= intervals[j][0] {
			intervals[i][1] = max(intervals[j][1], intervals[i][1])
			intervals = remove(intervals, j)
		} else {
			i++
		}
	}
	return intervals
}

func remove(intervals [][]int, index int) [][]int {
	return append(intervals[:index], intervals[index+1:]...)
}

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Print(merge(intervals))
}
