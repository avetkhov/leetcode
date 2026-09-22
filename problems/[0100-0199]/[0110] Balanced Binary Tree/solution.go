package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return checkHeight(root) != -1
}

func checkHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := checkHeight(root.Left)
	rightDepth := checkHeight(root.Right)

	if leftDepth == -1 || rightDepth == -1 || abs(leftDepth-rightDepth) > 1 {
		return -1
	}

	return 1 + max(leftDepth, rightDepth)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
