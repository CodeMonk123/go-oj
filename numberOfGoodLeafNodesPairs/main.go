package main

import "fmt"

// https://leetcode-cn.com/contest/weekly-contest-199/problems/number-of-good-leaf-nodes-pairs/
// 给你二叉树的根节点 root 和一个整数 distance 。
//
// 如果二叉树中两个 叶 节点之间的 最短路径长度 小于或者等于 distance ，那它们就可以构成一组 好叶子节点对 。
//
// 返回树中 好叶子节点对的数量 。
//
// 输入：root = [1,2,3,null,4], distance = 3
// 输出：1
// 解释：树的叶节点是 3 和 4 ，它们之间的最短路径的长度是 3 。这是唯一的好叶子节点对。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func visit(subRoot *TreeNode, distance int) (map[int]int, int) {
	if subRoot == nil {
		depthInfo := make(map[int]int)
		return depthInfo, 0
	}

	if subRoot.Left == nil && subRoot.Right == nil {
		depthInfo := make(map[int]int)
		depthInfo[1] = 1
		return depthInfo, 0
	}

	rightDepthInfo, rightCount := visit(subRoot.Right, distance)
	leftDepthInfo, leftCount := visit(subRoot.Left, distance)
	count :=  rightCount + leftCount

	for depth, numRight := range rightDepthInfo {
		for targetDepth := 1; targetDepth <= distance - depth; targetDepth ++ {
			if numLeft, ok := leftDepthInfo[targetDepth]; ok {
				count += numLeft * numRight
			}
		}
	}

	// merge depth info
	depthInfo := make(map[int]int)
	for depth, numRight := range rightDepthInfo {
			depthInfo[depth+1] = numRight
	}

	for depth, numLeft := range leftDepthInfo {
		if _, ok := depthInfo[depth+1]; !ok {
			depthInfo[depth+1] = numLeft
		} else {
			depthInfo[depth+1] += numLeft
		}
	}

	return depthInfo, count
}

func countPairs(root *TreeNode, distance int) int {
	if distance == 1 {
		return 0
	}

	rightDepthInfo, rightCount := visit(root.Right, distance)
	leftDepthInfo, leftCount := visit(root.Left, distance)

	res := rightCount + leftCount

	for depth, numRight := range rightDepthInfo {
		for targetDepth := 1; targetDepth <= distance - depth; targetDepth ++ {
			if numLeft, ok := leftDepthInfo[targetDepth]; ok {
				res += numLeft * numRight
			}
		}
	}

	return res
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right:  &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(countPairs(root, 2))
}