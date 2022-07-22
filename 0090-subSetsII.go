package main

import "fmt"

func subSetsWODup(nums []int) [][]int {

	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]]++
	}

	fmt.Println((numsMap))

	numsSets := [][]int{{}} // numsSets start with one set which is empty
	for num, count := range numsMap {
		numsSets = getNextLevelSets(numsSets, num, count)
	}

	return numsSets

}

func getNextLevelSets(numsSets [][]int, num int, count int) [][]int {

	currNumSet := []int{}
	combinedNumSets := [][]int{}

	for j := 0; j < count; j++ {

		currNumSet = append(currNumSet, num)

		for k := 0; k < len(numsSets); k++ {
			setWithNum := []int{}
			setWithNum = append(setWithNum, numsSets[k]...)
			setWithNum = append(setWithNum, currNumSet...)
			combinedNumSets = append(combinedNumSets, setWithNum)
			fmt.Println("Add Set:", setWithNum)
		}

	}
	fmt.Println(combinedNumSets)

	numsSets = append(numsSets, combinedNumSets...)
	return numsSets

}

// func main() {
// 	nums := []int{9, 0, 3, 5, 7}
// 	fmt.Println(subSetsWODup(nums))
// }
