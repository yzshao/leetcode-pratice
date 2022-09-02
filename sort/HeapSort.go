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
	//tree := []int{4,10,5,1,2,3}
	heapSort(list)
	fmt.Println(list)
}

func heapSort(nums []int) {
	buildHeap(nums, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		heapSwap(nums, i, 0)
		heapify(nums, i, 0)
	}
}

func buildHeap(tree []int, n int) {
	lastIndex := n - 1
	parentIndex := (lastIndex - 1) / 2
	for i := parentIndex; i >= 0; i-- {
		heapify(tree, n, i)
	}
}

func heapify(tree []int, n, index int) {
	if index >= n {
		return
	}
	left := 2*index + 1
	right := 2*index + 2
	max := index
	if left < n && tree[left] > tree[max] {
		max = left
	}
	if right < n && tree[right] > tree[max] {
		max = right
	}
	if max != index {
		heapSwap(tree, max, index)
		heapify(tree, n, max)
	}
}

func heapSwap(tree []int, i, j int) {
	tree[i], tree[j] = tree[j], tree[i]
}
