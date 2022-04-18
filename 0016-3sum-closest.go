package main

import (
	"math"
	"sort"
)

func sumThreeClosest(nums []int, target int) int {

	var iStart, iEnd int
	var closestDiff int

	sort.Ints(nums)

	closestDiff = math.MaxInt32

	for i, num := range nums {

		// Ingore the same value to avoid duplicate result set
		if i >= 1 && num == nums[i-1] {
			continue
		}

		iStart = i + 1
		iEnd = len(nums) - 1

		for iStart < iEnd {

			// Calculate diff
			diff := target - (num + nums[iStart] + nums[iEnd])

			if diff == 0 {
				return target
			} else if diff > 0 {
				iStart++
			} else if diff < 0 {
				iEnd--
			}

			// Set current data set as the closest one if less diff
			if math.Abs(float64(diff)) < math.Abs(float64(closestDiff)) {
				closestDiff = diff
			}

		}
	}

	return target - closestDiff

}

// func main() {
// 	nums := []int{0, 0, 0}
// 	fmt.Println(sumThreeClosest(nums, 1))
// }
