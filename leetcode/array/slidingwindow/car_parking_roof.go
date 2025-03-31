package main

import (
	"fmt"
	"math"
	"sort"
)

func carParkingRoof(cars []int, k int) int {
	result := math.MaxInt64
	sort.Slice(cars, func(i, j int) bool {
		return cars[i] < cars[j]
	})
	start := 0
	end := start + k
	for end <= len(cars) {
		if cars[end-1]-cars[start]+1 < result {
			result = cars[end-1] - cars[start] + 1
		}
		start++
		end = start + k
	}

	return result
}

func main() {
	arr := []int{6, 2, 12, 7}
	k := 3

	fmt.Print(carParkingRoof(arr, k))
}
