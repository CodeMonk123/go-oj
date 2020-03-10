package main

import (
	"fmt"
	"sort"
)

// 给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。
//
// 现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。
//
// 给定一个对数集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。
//
// 示例 :
//
// 输入: [[1,2], [2,3], [3,4]]
// 输出: 2
// 解释: 最长的数对链是 [1,2] -> [3,4]

type Pairs [][]int

func (p Pairs) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}

func (p Pairs) Swap(i, j int) {
	p[i][0], p[j][0] = p[j][0], p[i][0]
	p[i][1], p[j][1] = p[j][1], p[i][1]
}

func (p Pairs) Len() int {
	return len(p)
}

func findLongestChain(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}
	res := 1
	sort.Sort(Pairs(pairs))
	pivot := pairs[0][1]
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > pivot {
			res++
			pivot = pairs[i][1]
		}
	}
	return res
}

func main() {
	fmt.Println(findLongestChain([][]int{
		{4, 5},
		{2, 3},
		{1, 2},
	}))
}
