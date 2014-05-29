package main

import (
	"fmt"
)

func mergeSort(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	right := make([]int, 0)
	left := make([]int, 0)

	left = data[:len(data)/2]
	right = data[len(data)/2:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {

	result = make([]int, 0)
	iLeft, iRight := 0, 0

	for {
		if left[iLeft] < right[iRight] {
			result = append(result, left[iLeft])
			iLeft++
		} else {
			result = append(result, right[iRight])
			iRight++
		}

		if iLeft == len(left) {
			result = append(result, right[iRight:]...)
			break
		} else if iRight == len(right) {
			result = append(result, left[iLeft:]...)
			break
		}
	}

	return
}

func main() {
	//input := []int{84, 13, 73, 26, 32, 19, 91, 38}
	input := []int{3, 8, 2, 7, 4, 3, 3, 9, 82, 10, 100}
	fmt.Println("input : ", input)
	fmt.Println(mergeSort(input))
}
