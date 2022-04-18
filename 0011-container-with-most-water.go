package main

func maxArea(heights []int) int {
	maxArea := 0
	iStart := 0
	iEnd := len(heights) - 1

	for iStart < iEnd {
		// move the index for the lower one
		height := 0
		width := iEnd - iStart
		if heights[iStart] < heights[iEnd] {
			height = heights[iStart]
			iStart++
		} else {
			height = heights[iEnd]
			iEnd--
		}

		// calculate area
		area := width * height
		if area > maxArea {
			maxArea = area
		}

	}
	return maxArea
}

// func main() {
// 	var heights = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
// 	fmt.Printf("Max water area is %d", maxArea(heights))
// }
