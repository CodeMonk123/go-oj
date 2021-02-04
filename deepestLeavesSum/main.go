package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var currentLayer []*TreeNode
	var nextLayer []*TreeNode

	var result int
	currentLayer = append(currentLayer, root)

	for {
		for _, node := range currentLayer {
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}
		if len(nextLayer) == 0 {
			break
		} else {
			currentLayer = nextLayer
			nextLayer = []*TreeNode{}
		}
	}

	for _, node := range currentLayer {
		result += node.Val
	}

	return result
}
