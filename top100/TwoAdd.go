package main

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
	}

	str, _ := json.Marshal(addTwoNumbers(l1, l2))
	fmt.Println(string(str))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	retNode := &ListNode{}
	// 都不为空
	c := 0
	cur := retNode
	for l1 != nil || l2 != nil {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		num := (c + val1 + val2) % 10
		c = (c + val1 + val2) / 10
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if c > 0 {
		cur.Next = &ListNode{Val: c}
	}
	return retNode.Next
}
