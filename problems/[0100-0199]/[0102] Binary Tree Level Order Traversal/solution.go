package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	var enrich func(root *TreeNode) int
	enrich = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left, right := enrich(root.Left), enrich(root.Right)

		result = append(result, []int{root.Left.Val, root.Right.Val})

		return 1 + max(left, right)
	}

	return result
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	levelOrder(root)
}
