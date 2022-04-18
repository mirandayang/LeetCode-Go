package main

import (
	"fmt"
	"sort"
)

func sumThreeIntZero(nums []int) [][]int {

	var list = [][]int{}
	var iStart, iEnd int

	sort.Ints(nums)

	if len(nums) <= 2 || nums[0] > 0 {
		return nil
	}

	for i, num := range nums {

		// Ingore the same value to avoid duplicate result set
		if i >= 1 && num == nums[i-1] {
			fmt.Println("continue")
			continue
		}

		iStart = i + 1
		iEnd = len(nums) - 1

		for iStart < iEnd {
			fmt.Println("from index:", iStart, "value:", nums[iStart], ", end index: ", iEnd, "value:", nums[iEnd])
			if iStart > i+1 && nums[iStart] == nums[iStart-1] {
				iStart++
				continue
			}
			if iEnd < len(nums)-1 && nums[iEnd] == nums[iEnd+1] {
				iEnd--
				continue
			}

			sum := num + nums[iStart] + nums[iEnd]
			if sum == 0 {
				list = append(list, []int{num, nums[iStart], nums[iEnd]})
				iStart++
				iEnd--
			} else if sum < 0 {
				iStart++
				fmt.Println("start index increased")
			} else if sum > 0 {
				iEnd--
				fmt.Println("end index decreased")
			}
		}
	}
	return list
}

// func main() {
// 	nums := []int{-2, 0, 0, 2, 2}
// 	fmt.Println(sumThreeIntZero(nums))
// }
