package main

import (
	"fmt"
	"sort"
)

func main() {
	// [[] [1] [1 2] [1 2 2] [2] [2 2]]
	fmt.Println(subsets2([]int{1, 2, 2}))
	fmt.Println(subsets([]int{1, 2, 2}))
}

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	backtracking(nums, 0, path, &res)
	return res
}

func backtracking(nums []int, start int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backtracking(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}

func subsets2(nums []int) [][]int {
	var res [][]int
	var path []int
	sort.Ints(nums)
	backtracking2(nums, 0, path, &res)
	return res
}

func backtracking2(nums []int, start int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		backtracking2(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
