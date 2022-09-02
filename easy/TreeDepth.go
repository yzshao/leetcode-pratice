package main

// 深度优先遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 广度优先遍历
func maxDepthByQueue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	stk := make([]*TreeNode, 0)
	stk = append(stk, root)
	for len(stk) > 0 {
		size := len(stk)
		for size > 0 {
			node := stk[0]
			stk = stk[1:]
			if node != nil {
				if node.Left != nil {
					stk = append(stk, node.Left)
				}
				if node.Right != nil {
					stk = append(stk, node.Right)
				}
			}
			size--
		}
		ans++
	}
	return ans
}

// 深度优先遍历
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 1.左孩子和有孩子都为空的情况，说明到达了叶子节点，直接返回1即可
	if root.Left == nil && root.Right == nil {
		return 1
	}
	m1 := minDepth(root.Left)
	m2 := minDepth(root.Right)
	// 2.如果左孩子和由孩子其中一个为空，那么需要返回比较大的那个孩子的深度
	if (root.Left == nil && root.Right != nil) || root.Left != nil && root.Right == nil {
		return m1 + m2 + 1
	}
	// 3.最后一种情况，也就是左右孩子都不为空，返回最小深度+1即可
	return Min(m1, m2) + 1
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 广度优先遍历
func minDepthByQueue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	stk := make([]*TreeNode, 0)
	stk = append(stk, root)
	for len(stk) > 0 {
		size := len(stk)
		for size > 0 {
			node := stk[0]
			stk = stk[1:]
			// 找到叶子节点，那么直接返回当前深度就是最小深度
			if node.Left == nil && node.Right == nil {
				return ans + 1
			}
			if node.Left != nil {
				stk = append(stk, node.Left)
			}
			if node.Right != nil {
				stk = append(stk, node.Right)
			}
			size--
		}
		ans++
	}
	return ans
}
