package main

import "fmt"

func sortColors(colors []int) []int {
	zero, one := 0, 0
	two := len(colors) - 1
	for one <= two {
		if colors[one] == 0 {
			swap(colors, zero, one)
			zero++
			one++
		} else if colors[one] == 1 {
			one++
		} else if colors[one] == 2 {
			swap(colors, one, two)
			two--
		}
		fmt.Println(colors)
	}
	return colors
}

func swap(nums []int, i int, j int) []int {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
	return nums
}

// func main() {
// 	colors := []int{2, 0, 1}
// 	fmt.Println(sortColors(colors))
// }
