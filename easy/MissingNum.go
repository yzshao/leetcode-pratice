package main

import "fmt"

func main() {
	fmt.Println(missingNumber2([]int{3, 0, 1}))
}

func missingNumber(nums []int) int {
	n := len(nums)
	tempMap := make(map[int]bool)
	for _, num := range nums {
		tempMap[num] = true
	}
	for i := 0; i <= n; i++ {
		if !tempMap[i] {
			return i
		}
	}
	return 0
}

/**
0异或任何数=任何数
相同值异或=0，否则为1
*/
func missingNumber2(nums []int) (xor int) {
	for i, num := range nums {
		xor = xor ^ (i ^ num)
	}
	return xor ^ len(nums)
}
