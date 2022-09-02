package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var preNode *ListNode
	cur := head
	for cur != nil {
		nextNode := cur.Next
		cur.Next = preNode
		preNode = cur
		cur = nextNode
	}
	return preNode
}
