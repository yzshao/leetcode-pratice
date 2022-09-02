package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permute([]int{2, 2, 2}))
	fmt.Println(permute2([]int{2, 2, 2}))
}

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	var used = make([]bool, len(nums))
	backtrackingPermute(nums, path, &res, used)
	return res
}

func backtrackingPermute(nums []int, path []int, res *[][]int, used []bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backtrackingPermute(nums, path, res, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func permute2(nums []int) [][]int {
	var res [][]int
	var path []int
	var used = make([]bool, len(nums))
	sort.Ints(nums)
	backtrackingPermute2(nums, path, &res, used)
	return res
}

func backtrackingPermute2(nums []int, path []int, res *[][]int, used []bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		// 新添加的剪枝逻辑，固定相同的元素在排列中的相对位置
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backtrackingPermute2(nums, path, res, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}
