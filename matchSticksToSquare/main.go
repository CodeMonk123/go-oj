// 还记得童话《卖火柴的小女孩》吗？现在，你知道小女孩有多少根火柴，请找出一种能使用所有火柴拼成一个正方形的方法。不能折断火柴，可以把火柴连接起来，并且每根火柴都要用到。

// 输入为小女孩拥有火柴的数目，每根火柴用其长度表示。输出即为是否能用所有的火柴拼成正方形。

// 示例 1:

// 输入: [1,1,2,2,2]
// 输出: true

// 解释: 能拼成一个边长为2的正方形，每边两根火柴。
// 示例 2:

// 输入: [3,3,3,3,4]
// 输出: false

// 解释: 不能用所有火柴拼成一个正方形。
// 注意:

// 给定的火柴长度和在 0 到 10^9之间。
// 火柴数组的长度不超过15。

package main

import (
	"fmt"
	"sort"
)

type reverseSort []int

func (this reverseSort) Len() int {
	return len(this)
}

func (this reverseSort) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this reverseSort) Less(i, j int) bool {
	return this[i] > this[j]
}

func copyGroups(groups []int) []int {
	newGroups := append([]int{}, groups...)
	return newGroups
}

func dfs(current int, nums []int, groups []int, target int) bool {
	if current == len(nums) {
		for _, v := range groups {
			if v != target {
				return false
			}
		}
		return true
	}
	valid := false
	res := false
	for i := 0; i < 4; i++ {
		newGroups := copyGroups(groups)
		if newGroups[i]+nums[current] <= target {
			valid = true
			newGroups[i] = newGroups[i] + nums[current]
			res = res || dfs(current+1, nums, newGroups, target)
		}
	}
	if !valid {
		return false
	}
	return res
}

func makesquare(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%4 != 0 || sum == 0 {
		return false
	}

	groups := []int{0, 0, 0, 0}
	sort.Sort(reverseSort(nums))
	return dfs(0, nums, groups, sum/4)

}

func main() {
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2})) //true
	fmt.Println(makesquare([]int{3, 3, 3, 3, 4})) //false
}
