package main

import "fmt"

/*
Given the root of a tree, you are asked to find the most frequent subtree sum. The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself). So what is the most frequent subtree sum value? If there is a tie, return all the values with the highest frequency in any order.

Examples 1
Input:

  5
 /  \
2   -3
return [2, -3, 4], since all the values happen only once, return all of them in any order.
Examples 2
Input:

  5
 /  \
2   -5
return [2], since 2 happens twice, however -5 only occur once.
Note: You may assume the sum of values in any subtree is in the range of 32-bit signed integer.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this *TreeNode) CalcSum(record *map[int]int) int {
	if this == nil {
		return 0
	} else {
		sum := this.Val + this.Left.CalcSum(record) + this.Right.CalcSum(record)
		(*record)[sum]++ //sum出现次数+1
		return sum
	}
}

func findFrequentTreeSum(root *TreeNode) []int {
	sumFrequency := make(map[int]int)
	root.CalcSum(&sumFrequency)

	maxFreq := 0
	for _, v := range sumFrequency {
		if v > maxFreq {
			maxFreq = v
		}
	}

	res := []int{}
	for k, v := range sumFrequency {
		if v == maxFreq {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	root := TreeNode{-3, nil, nil}
	root.Left = &TreeNode{5, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	fmt.Println(findFrequentTreeSum(&root))
}
