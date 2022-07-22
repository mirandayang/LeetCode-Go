package main

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	j, k := m-1, n-1

	for i := m + n - 1; i >= 0; i-- {
		if j >= 0 && k >= 0 {
			if nums1[j] >= nums2[k] {
				nums1[i] = nums1[j]
				j--
			} else {
				nums1[i] = nums2[k]
				k--
			}

		} else if j < 0 {
			nums1[i] = nums2[k]
			k--
		} else if k < 0 {
			nums1[i] = nums1[j]
			j--
		}
	}
}

// func main() {
// 	num1 := []int{1, 2, 3, 0, 0, 0}
// 	num2 := []int{4, 5, 6}
// 	mergeSortedArray(num1, 3, num2, 3)
// 	fmt.Println(num1)
// }
