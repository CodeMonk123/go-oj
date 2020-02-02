// 给你一棵二叉树，它的根为 root 。请你删除 1 条边，使二叉树分裂成两棵子树，且它们子树和的乘积尽可能大。

// 由于答案可能会很大，请你将结果对 10^9 + 7 取模后再返回。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package maxProductOfSplitBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxProd int64

func calcSum(subRoot *TreeNode) int {
	if subRoot == nil {
		return 0
	}
	leftSum := calcSum(subRoot.Left)
	rightSum := calcSum(subRoot.Right)
	subRoot.Val += (leftSum + rightSum)
	return subRoot.Val
}

func findMax(subRoot *TreeNode, sum int) {
	if subRoot.Left != nil {
		prod1 := int64(sum-subRoot.Left.Val) * int64(subRoot.Left.Val)
		if prod1 > maxProd {
			maxProd = prod1
		}
		findMax(subRoot.Left, sum)
	}
	if subRoot.Right != nil {
		prod2 := int64(sum-subRoot.Right.Val) * int64(subRoot.Right.Val)
		if prod2 > maxProd {
			maxProd = prod2
		}
		findMax(subRoot.Right, sum)
	}
}

func maxProduct(root *TreeNode) int {
	calcSum(root)
	findMax(root, root.Val)
	res := int(maxProd % 1000000007)
	maxProd = 0
	return res
}
