package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(levelOrderBottom(root))
}

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	preorder(root, &ret)
	return ret
}

func preorder(node *TreeNode, ret *[]int) {
	if node == nil {
		return
	}
	*ret = append(*ret, node.Val)
	preorder(node.Left, ret)
	preorder(node.Right, ret)
}

func preorderTraversalByStack(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			ret = append(ret, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return ret
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	inorder(root, &ret)
	return ret
}

func inorder(node *TreeNode, ret *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, ret)
	*ret = append(*ret, node.Val)
	inorder(node.Right, ret)
	return
}

func inorderTraversalByStack(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		root = node.Right
	}
	return ret
}

func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	postorder(root, &ret)
	return ret
}

func postorder(node *TreeNode, ret *[]int) {
	if node == nil {
		return
	}
	postorder(node.Left, ret)
	postorder(node.Right, ret)
	*ret = append(*ret, node.Val)
	return
}

func postorderTraversalByStack(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var preNode *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == preNode {
			ret = append(ret, root.Val)
			preNode = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return ret
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	stk := make([]*TreeNode, 0)
	stk = append(stk, root)
	for len(stk) > 0 {
		size := len(stk)
		ret := make([]int, 0)
		for size > 0 {
			node := stk[0]
			stk = stk[1:]
			if node != nil {
				ret = append(ret, node.Val)
				if node.Left != nil {
					stk = append(stk, node.Left)
				}
				if node.Right != nil {
					stk = append(stk, node.Right)
				}
			}
			size--
		}
		res = append(res, ret)
	}
	return res
}

// 二叉树的层序遍历（自底向上）
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	stk := make([][]*TreeNode, 0)
	tempStk := make([]*TreeNode, 0)
	tempStk = append(tempStk, root)
	for len(tempStk) > 0 {
		size := len(tempStk)
		queue := make([]*TreeNode, 0)
		for size > 0 {
			node := tempStk[0]
			tempStk = tempStk[1:]
			queue = append(queue, node)
			if node.Left != nil {
				tempStk = append(tempStk, node.Left)
			}
			if node.Right != nil {
				tempStk = append(tempStk, node.Right)
			}
			size--
		}
		stk = append(stk, queue)
	}
	res := make([][]int, 0)
	for i := len(stk) - 1; i >= 0; i-- {
		ret := make([]int, 0)
		for _, node := range stk[i] {
			ret = append(ret, node.Val)
		}
		res = append(res, ret)
	}
	return res
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
