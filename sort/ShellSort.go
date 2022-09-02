package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var list []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 15; i++ {
		list = append(list, int(r.Intn(1000)))
	}
	//nums := []int{5,4,3,2,1,6,0,9,10}
	shellSort(list)
	fmt.Println(list)
}

func shellSort(nums []int) {
	size := len(nums)
	for gap := size / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < size; i++ {
			preIndex := i
			curVal := nums[i]
			for preIndex >= gap && curVal < nums[preIndex-gap] {
				nums[preIndex] = nums[preIndex-gap]
				preIndex = preIndex - gap
			}
			nums[preIndex] = curVal
		}
	}
	return
}
