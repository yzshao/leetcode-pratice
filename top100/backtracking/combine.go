package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combine2([]int{1, 1, 2, 2}, 2))
}

func combine(nums []int, k int) [][]int {
	var res [][]int
	var path []int
	backtrackingCombine(nums, 0, path, &res, k)
	return res
}

func backtrackingCombine(nums []int, start int, path []int, res *[][]int, k int) {
	if len(path) == k {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrackingCombine(nums, i+1, path, res, k)
		path = path[:len(path)-1]
	}
}

func combine2(nums []int, target int) [][]int {
	var res [][]int
	var path []int
	var sum = 0
	sort.Ints(nums)
	backtrackingCombine2(nums, 0, path, &res, target, &sum)
	return res
}

func backtrackingCombine2(nums []int, start int, path []int, res *[][]int, target int, sum *int) {
	if *sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	if *sum > target {
		return
	}
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		*sum = *sum + nums[i]
		path = append(path, nums[i])
		backtrackingCombine2(nums, i+1, path, res, target, sum)
		path = path[:len(path)-1]
		*sum = *sum - nums[i]
	}
}
