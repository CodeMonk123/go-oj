package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	res += distributeCoins(root.Left)
	res += distributeCoins(root.Right)
	if root.Left != nil {
		if root.Left.Val > 1 {
			root.Val += root.Left.Val - 1
			res += root.Left.Val - 1
			root.Left.Val = 1
		} else if root.Left.Val < 1 {
			root.Val -= (1 - root.Left.Val)
			res += (1 - root.Left.Val)
			root.Left.Val = 1
		}
	}

	if root.Right != nil {
		if root.Right.Val > 1 {
			root.Val += root.Right.Val - 1
			res += root.Right.Val - 1
			root.Right.Val = 1
		} else if root.Right.Val < 1 {
			root.Val -= (1 - root.Right.Val)
			res += (1 - root.Right.Val)
			root.Right.Val = 1
		}
	}

	return res
}

func main() {
	root := new(TreeNode)
	root.Left = &TreeNode{3, nil, nil}
	root.Right = new(TreeNode)
	fmt.Println(distributeCoins(root))
}
