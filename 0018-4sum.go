package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	var list = [][]int{}

	sort.Ints(nums)
	fmt.Println(nums)
	length := len(nums)

	for i := 0; i <= length-4; i++ {
		// Ingore the same value to avoid duplicate result set
		if i >= 1 && nums[i] == nums[i-1] {
			fmt.Println("continue")
			continue
		}

		for j := i + 1; j <= length-3; j++ {

			// Ingore the same value to avoid duplicate result set
			if j > i+1 && nums[j] == nums[j-1] {
				fmt.Println("continue")
				continue
			}

			iStart, iEnd := j+1, length-1

			for iStart < iEnd {
				fmt.Println("check:", nums[i], nums[j], nums[iStart], nums[iEnd])
				if iStart > j+1 && nums[iStart] == nums[iStart-1] {
					iStart++
					continue
				}
				if iEnd < length-1 && nums[iEnd] == nums[iEnd+1] {
					iEnd--
					continue
				}

				sum := nums[i] + nums[j] + nums[iStart] + nums[iEnd]
				if sum == target {
					fmt.Println("found!")
					list = append(list, []int{nums[i], nums[j], nums[iStart], nums[iEnd]})
					iStart++
					iEnd--
				} else if sum < target {
					iStart++
				} else if sum > target {
					iEnd--
				}
			}
		}
	}
	return list
}

// func main() {
// 	nums := []int{1, 0, -1, 0, -2, 2}
// 	fmt.Println(fourSum(nums, 0))
// }
