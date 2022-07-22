package main

func sumOfMinArrays(nums []int) int {
	var stack []int
	var left = make([]int, len(nums))
	var right = make([]int, len(nums))
	var sum int
	var j int
	const MOD int = 1000000007

	// build array of first left index that is larger
	for i := 0; i < len(nums); i++ {
		//stack = addStack(nums, stack, i)
		for j = len(stack) - 1; j >= 0; j-- {
			if nums[stack[j]] >= nums[i] {
				stack = stack[:j]
			}
		}

		stack = append(stack, i)

		if len(stack) <= 1 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-2]
		}
	}

	stack = []int{}

	// build array of first right index that is larger
	for i := len(nums) - 1; i >= 0; i-- {
		//	stack = addStack(nums, stack, i)
		for j = len(stack) - 1; j >= 0; j-- {
			if nums[stack[j]] > nums[i] {
				stack = stack[:j]
			}
		}

		stack = append(stack, i)

		if len(stack) <= 1 {
			right[i] = len(nums)
		} else {
			right[i] = stack[len(stack)-2]
		}
	}

	for i := 0; i < len(nums); i++ {
		sum = (sum + (i-left[i])*(right[i]-i)*nums[i]) % MOD
	}

	return sum
}

// func main() {
// 	nums := []int{71, 55, 82, 55}
// 	fmt.Println(sumOfMinArrays(nums))
// }
