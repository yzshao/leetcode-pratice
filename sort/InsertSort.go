package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 2, 1, 6, 0, 9, 10}
	insertSort(nums)
	fmt.Println(nums)
}

func insertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		preIndex := i
		curVal := nums[preIndex]
		for preIndex > 0 && curVal < nums[preIndex-1] {
			nums[preIndex] = nums[preIndex-1]
			preIndex--
		}
		nums[preIndex] = curVal
	}
	return
}
