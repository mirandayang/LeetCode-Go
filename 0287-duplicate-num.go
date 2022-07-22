package main

func DupNum(nums []int) int {
	var i, j, k int
	for i != j || i == 0 {
		i = nums[i]
		j = nums[nums[j]]
	}
	for k != i {
		i = nums[i]
		k = nums[k]
	}
	return i
}

// func main() {
// 	nums := []int{1, 3, 4, 2, 2}
// 	fmt.Println(DupNum(nums))
// }
