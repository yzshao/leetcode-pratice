package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}))
}

/**
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。

链接：https://leetcode.cn/problems/3sum
*/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	// 先让数组有序
	sort.Ints(nums)
	// 使用三个指针完成数据遍历以及数值计算
	n := len(nums)
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		x := i + 1
		y := n - 1
		target := -nums[i]
		for x < y {
			if nums[x]+nums[y] == target {
				res = append(res, []int{nums[i], nums[x], nums[y]})
				for x < y && nums[x+1] == nums[x] {
					x++
				}
				for x < y && nums[y-1] == nums[y] {
					y--
				}
				x++
				y--
			} else if nums[x]+nums[y] > target {
				y--
			} else {
				x++
			}
		}
	}
	return res
}
