package main

import "fmt"

type stack struct {
	top   int
	array []int
}

func (p *stack) push(e int) {
	p.top++
	if p.top >= len(p.array) {
		p.array = append(p.array, e)
	} else {
		p.array[p.top] = e
	}
}

func (p *stack) pop() {
	if p.top != -1 {
		p.top--
	}
}

func (p *stack) getTopElement() int {
	if p.top != -1 {
		return p.array[p.top]
	}
	return 0
}

func subarrayRanges(arr []int) int64 {
	var sum int64
	var left = make([]int, len(arr))
	var right = make([]int, len(arr))
	var iStack = &stack{top: -1, array: []int{}}

	for i, value := range arr {

		// pop util the top is lower than current value
		for iStack.top != -1 && arr[iStack.getTopElement()] > value {
			iStack.pop()
		}

		// build left array which contains the closest element index that is smaller
		if iStack.top == -1 {
			left[i] = -1
		} else {
			left[i] = iStack.getTopElement()
		}

		// push curent element
		iStack.push(i)

	}

	fmt.Println(left)

	// reset stack
	iStack.top = -1

	for i := len(arr) - 1; i >= 0; i-- {

		// pop util the top is lower than current value
		for iStack.top != -1 && arr[iStack.getTopElement()] >= arr[i] {
			iStack.pop()
		}

		// build right array which contains the closest element index that is smaller
		if iStack.top == -1 {
			right[i] = len(arr)
		} else {
			right[i] = iStack.getTopElement()
		}

		// push curent element
		iStack.push(i)

	}

	fmt.Println(right)

	for i, value := range arr {
		sum -= int64((i - left[i]) * (right[i] - i) * value)
	}

	// reset data
	iStack.top = -1
	left = make([]int, len(arr))
	right = make([]int, len(arr))

	for i, value := range arr {

		// pop util the top is lower than current value
		for iStack.top != -1 && arr[iStack.getTopElement()] < value {
			iStack.pop()
		}

		// build left array which contains the closest element index that is smaller
		if iStack.top == -1 {
			left[i] = -1
		} else {
			left[i] = iStack.getTopElement()
		}

		// push curent element
		iStack.push(i)

	}

	fmt.Println(left)

	// reset stack
	iStack.top = -1

	for i := len(arr) - 1; i >= 0; i-- {

		// pop util the top is lower than current value
		for iStack.top != -1 && arr[iStack.getTopElement()] <= arr[i] {
			iStack.pop()
		}

		// build right array which contains the closest element index that is smaller
		if iStack.top == -1 {
			right[i] = len(arr)
		} else {
			right[i] = iStack.getTopElement()
		}

		// push curent element
		iStack.push(i)

	}

	fmt.Println(right)

	for i, value := range arr {
		sum += int64((i - left[i]) * (right[i] - i) * value)
	}

	return sum
}

// func main() {
// 	nums := []int{4, -2, -3, 4, 1}
// 	fmt.Println(subarrayRanges(nums))
// }
