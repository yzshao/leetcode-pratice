package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 2, 1, 6, 0, 9, 10}
	quickSort(nums)
	fmt.Println(nums)
}

func quickSort(nums []int) {
	quick(nums, 0, len(nums)-1)
}

func quick(nums []int, start, end int) {
	if start >= end {
		//起始坐标大于等于结束坐标，无法分区，直接返回
		return
	}
	//进行分区，拿到分区位置
	parti := partition(nums, start, end)
	//左边继续快速排序
	quick(nums, start, parti-1)
	//右边继续快速排序
	quick(nums, parti+1, end)
}

func partition(nums []int, start, end int) int {
	//起始坐标对应的值作为分区值
	partiVal := nums[start]
	//遍历所有元素
	for start < end {
		//end指针值 >= parti指针值， 右移
		for start < end && partiVal <= nums[end] {
			end--
		}
		//填补start位置空值
		//end指针值 < partiVal end指针值移到start位置
		//end 位置值空
		nums[start] = nums[end]
		//start指针值 <= parti指针值， 左移
		for start < end && partiVal >= nums[start] {
			start++
		}
		//填补end位置空值
		//start指针值 > partiVal start指针值移到end位置
		//start 位置值空
		nums[end] = nums[start]
	}
	//partiVal 填补 start位置的空值
	nums[start] = partiVal
	return start
}
