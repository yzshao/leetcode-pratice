package main

import "fmt"

func main() {
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

/**
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。
链接：https://leetcode.cn/problems/container-with-most-water

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
*/
func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			num := minNum(height[i], height[j])
			temp := num * (j - i)
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

func maxArea2(height []int) int {
	i, j, res := 0, len(height)-1, 0
	for i < j {
		if height[i] < height[j] {
			res = maxNum(res, (j-i)*height[i])
			i++
		} else {
			res = maxNum(res, (j-i)*height[j])
			j--
		}
	}
	return res
}

func maxNum(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minNum(x, y int) int {
	if x < y {
		return x
	}
	return y
}
