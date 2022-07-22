package main

func pancakeSort(arr []int) []int {
	var length int = len(arr)
	var sortIndexArr []int
	for i := length - 1; i > 0; i-- {
		sortIndexArr = moveMaxnumToTheEnd(arr, sortIndexArr, length, i)
	}
	return sortIndexArr
}

func moveMaxnumToTheEnd(nums []int, sort []int, length int, pos int) []int {
	// get max number
	var max_index, maxnum int
	var i int = 0
	for i <= pos {
		if nums[i] > maxnum {
			max_index = i
			maxnum = nums[i]
		}
		i++
	}

	// move max number to the end of the array
	if max_index != pos {

		if max_index != 0 {
			sort = append(sort, max_index+1)
			reverse(nums, max_index)
		}

		sort = append(sort, pos+1)
		reverse(nums, pos)
	}

	return sort

}

func reverse(nums []int, k int) {
	i, j := 0, k
	for i < j {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		i++
		j--
	}
}

// func main() {
// 	nums := []int{3, 2, 4, 1}
// 	fmt.Println(pancakeSort(nums))
// }
