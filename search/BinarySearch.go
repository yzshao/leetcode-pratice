package main

import (
	"fmt"
)

func main() {
	//var list []int
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i := 0; i < 15; i++ {
	//	list = append(list, int(r.Intn(1000)))
	//}
	//sort.Ints(list)
	list := []int{1, 29, 37, 40, 51, 63, 110}
	fmt.Println(binarySearch(list, 110))
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
