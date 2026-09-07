package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	Node *TreeNode
	Sum  int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	items := []Item{{Node: root, Sum: root.Val}}

	for len(items) > 0 {
		current := items[0]
		items = items[1:]

		if current.Node.Left == nil && current.Node.Right == nil && current.Sum == targetSum {
			return true
		}

		if current.Node.Left != nil {
			items = append(items, Item{Node: current.Node.Left, Sum: current.Sum + current.Node.Left.Val})
		}
		if current.Node.Right != nil {
			items = append(items, Item{Node: current.Node.Right, Sum: current.Sum + current.Node.Right.Val})
		}
	}

	return false
}
