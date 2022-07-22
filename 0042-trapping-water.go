package main

import (
	"fmt"
)

func trappingWater(heights []int) int {

	size := len(heights)
	left_max := make([]int, size)
	right_max := make([]int, size)

	left_max[0] = 0
	for i := 1; i < size-1; i++ {
		left_max[i] = left_max[i-1]
		if heights[i-1] > left_max[i] {
			left_max[i] = heights[i-1]
		}
	}
	fmt.Println(left_max)

	right_max[size-1] = 0
	for i := size - 2; i > 0; i-- {
		right_max[i] = right_max[i+1]
		if heights[i+1] > right_max[i] {
			right_max[i] = heights[i+1]
		}
	}
	fmt.Println(right_max)

	var area, increasedArea int
	for i := 1; i < size-1; i++ {
		if left_max[i] >= right_max[i] {
			increasedArea = right_max[i] - heights[i]
		} else {
			increasedArea = left_max[i] - heights[i]
		}
		if increasedArea > 0 {
			area += increasedArea
		}
	}

	return area

}

// func main() {
// 	heights := []int{4, 2, 0, 3, 2, 5}
// 	fmt.Println(trappingWater(heights))
// }
