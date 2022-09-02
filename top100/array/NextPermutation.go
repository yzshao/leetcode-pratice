package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 3, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}

/**
下一个排列

整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。
例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，
那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。
例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。
必须 原地 修改，只允许使用额外常数空间。
示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
链接：https://leetcode.cn/problems/next-permutation
*/
func nextPermutation(nums []int) {
	// 1,2,3,4  1,2,4,3  1,3,2,4  1,3,4,2
	// 1,2,3,4,5  1,2,3,5,4  1,2,4,3,5  1,2,4,5,3
	// 1,3,2,4,5  1,3,2,5,4   1,2,3,7,6,5,4   1,2,4,3,5,6,7
	n := len(nums)
	index := -1
	for i := n - 1; i > 0; i-- {
		if nums[i] <= nums[i-1] {
			continue
		}
		index = i - 1
		break
	}
	reverseIndex := 0
	if index != -1 {
		tempNum := nums[index]
		nextNum := nums[index+1]
		nextIndex := index + 1
		for i := index + 1; i < n; i++ {
			if nums[i] <= tempNum {
				continue
			}
			if nextNum >= nums[i] {
				nextNum = nums[i]
				nextIndex = i
				continue
			}
		}
		nums[nextIndex] = tempNum
		nums[index] = nextNum
		reverseIndex = index + 1
	}
	for i, j := reverseIndex, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
