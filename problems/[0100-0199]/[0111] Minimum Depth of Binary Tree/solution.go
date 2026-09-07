package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	depth := 0

	if root == nil {
		return depth
	}

	queue := []*TreeNode{root}
	depth = 1

	for len(queue) > 0 {
		for _, parent := range queue {
			queue = queue[1:]

			if parent.Left == nil && parent.Right == nil {
				return depth
			}
			if parent.Left != nil {
				queue = append(queue, parent.Left)
			}
			if parent.Right != nil {
				queue = append(queue, parent.Right)
			}
		}

		depth++
	}

	return depth
}
