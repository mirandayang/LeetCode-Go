package main

func subSets(nums []int) [][]int {
	sets := [][]int{}
	size := len(nums)
	for i := 0; i < 1<<size; i++ {
		set := []int{}
		for j := 0; j < size; j++ {
			if i>>j&1 == 1 {
				set = append(set, nums[j])
			}
		}
		sets = append(sets, set)
	}
	return sets
}

// func main() {
// 	nums := []int{1, 2, 3}
// 	fmt.Println(subSets(nums))
// }
