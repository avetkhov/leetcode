package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	var result []float64
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		var sum float64

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			sum += float64(node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, sum/float64(size))
	}

	return result
}
