package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var list []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 15000; i++ {
		list = append(list, int(r.Intn(100000)))
	}
	//nums := []int{5,4,3,2,1,6,0,9,10}
	fmt.Println(mergeSort(list))
}

func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[0:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	ret := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}
	for len(left) > 0 {
		ret = append(ret, left[0])
		left = left[1:]
	}
	for len(right) > 0 {
		ret = append(ret, right[0])
		right = right[1:]
	}
	return ret
}
