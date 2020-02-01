package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(subRoot *TreeNode, val int) *TreeNode {
	if subRoot == nil {
		newRoot := new(TreeNode)
		newRoot.Val = val
		return newRoot
	} else if subRoot.Val > val {
		subRoot.Left = insert(subRoot.Left, val)
		return subRoot
	} else {
		subRoot.Right = insert(subRoot.Right, val)
		return subRoot
	}
}

func buildTree(nums []int) *TreeNode {
	var root *TreeNode
	for _, v := range nums {
		root = insert(root, v)
	}
	return root
}

func countNodes(subRoot *TreeNode) int {
	result := 0
	if subRoot == nil {
		return result
	} else {
		result += countNodes(subRoot.Left)
		result += countNodes(subRoot.Right)
	}
	return result + 1
}

func kthSmallest(root *TreeNode, k int) int {
	leftNum := countNodes(root.Left)
	if leftNum+1 == k {
		return root.Val
	} else if leftNum+1 < k {
		return kthSmallest(root.Right, k-(leftNum+1))
	} else {
		return kthSmallest(root.Left, k)
	}
}

func main() {
	root := buildTree([]int{5, 3, 6, 2, 4, 1})
	fmt.Println(kthSmallest(root, 3))
}
