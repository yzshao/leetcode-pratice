package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 2, 1, 6, 0, 9, 10}
	selectSort(nums)
	fmt.Println(nums)
}

func selectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums)-1; j++ {
			if nums[i] > nums[j] {
				swap(&nums[i], &nums[j])
			}
		}
	}
	return
}
