package main

// 给你一个树，请你 按中序遍历 重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。
//
//  
//
// 示例 ：
//
// 输入：[5,3,6,2,4,null,8,1,null,null,null,7,9]
//
// 5
// / \
// 3    6
// / \    \
// 2   4    8
//  /        / \
// 1        7   9
//
// 输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
//
// 1
//   \
//    2
//     \
//      3
//       \
//        4
//         \
//          5
//           \
//            6
//             \
//              7
//               \
//                8
//                 \
// 9
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/increasing-order-search-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


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

func inOrderTraverse(subroot *TreeNode) []*TreeNode{
	if subroot == nil {
		return []*TreeNode{}
	}
	res := []*TreeNode{}
	left := inOrderTraverse(subroot.Left)
	res = append(res, left...)
	res = append(res, subroot)
	right := inOrderTraverse(subroot.Right)
	res = append(res, right...)
	return res
}

func increasingBST(root *TreeNode) *TreeNode {
	res := inOrderTraverse(root)
	newRoot := res[0]
	p := newRoot
	for _, q := range res[1:]{
		p.Right = q
		p.Left = nil
		q.Left = nil
		p = q
	}
	return newRoot
}
