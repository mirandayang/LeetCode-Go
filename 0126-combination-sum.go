package main

func combinationSum(k int, sum int) [][]int {
	result := [][]int{}
	for i := 0; i < 1<<9; i++ {
		counter := 0
		sumOfAll := 0
		set := []int{}
		for j := 0; j < 9; j++ {
			if i>>j&1 == 1 {
				counter++
				sumOfAll += j + 1
				set = append(set, j+1)
			}
		}
		if counter == k && sumOfAll == sum {
			result = append(result, set)
		}
	}
	return result
}

// func main() {
// 	fmt.Println(combinationSum(2, 10))
// }
