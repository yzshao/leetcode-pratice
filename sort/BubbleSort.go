package main

import "fmt"

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, ":", y)

	nums := []int{5, 4, 3, 2, 1, 6, 0, 9, 10}
	bubbleSort(nums)
	fmt.Println(nums)
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				swap(&nums[j], &nums[j+1])
			}
		}
	}
	return
}

func swap(x, y *int) {
	*x = *x ^ *y
	*y = *x ^ *y
	*x = *x ^ *y
}
