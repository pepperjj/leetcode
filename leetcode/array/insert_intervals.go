// https://leetcode.com/problems/insert-interval/description/
package main

import (
	"fmt"
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	for i, j := 0, 1; i < len(intervals); {
		for j < len(intervals) && intervals[i][1] >= intervals[j][0] {
			intervals[j][0] = min(intervals[i][0], intervals[j][0])
			intervals[j][1] = max(intervals[i][1], intervals[j][1])
			intervals = intervals[1:]
		}
		res = append(res, intervals[i])
		i++
		j++
	}
	return res
}

func main() {
	intervals := [][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	res := insert(intervals, []int{4, 8})
	for _, interval := range res {
		fmt.Println(interval)
	}
}
