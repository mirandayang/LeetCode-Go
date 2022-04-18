package main

import "fmt"

//import "fmt"

func twoSum(nums []int, sum int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		fmt.Println("loop at line", i, "with value", num)
		_, ok := m[sum-num]
		if ok == true {
			return []int{i, m[sum-num]}
		} else {
			m[num] = i
		}
	}
	return nil
}

// func main() {
// 	results := twoSum([]int{3, 6, 9, 10}, 9)
// 	if results != nil {
// 		fmt.Println(results)
// 	} else {
// 		fmt.Println("not found!")
// 	}
// }
