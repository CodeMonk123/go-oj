package main

import "fmt"

// 给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
//
// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
//
// 示例 1:
//
// 输入: nums: [1, 1, 1, 1, 1], S: 3
// 输出: 5
// 解释:
//
// -1+1+1+1+1 = 3
// +1-1+1+1+1 = 3
// +1+1-1+1+1 = 3
// +1+1+1-1+1 = 3
// +1+1+1+1-1 = 3
//
// 一共有5种方法让最终目标和为3。
// 注意:
//
// 数组非空，且长度不会超过20。
// 初始的数组的和不会超过1000。
// 保证返回的最终结果能被32位整数存下。

func dfs(currentSum, next, target int, nums, sums []int) int {
	if next == len(nums)-1 {
		count := 0
		if currentSum+nums[next] == target {
			count++
		}
		if currentSum-nums[next] == target {
			count++
		}
		return count
	}
	//  剪枝
	if currentSum+sums[next] < target || currentSum-sums[next] > target {
		return 0
	}
	return dfs(currentSum+nums[next], next+1, target, nums, sums) + dfs(currentSum-nums[next], next+1, target, nums, sums)
}

func findTargetSumWays(nums []int, S int) int {
	sums := make([]int, len(nums))
	n := len(nums) - 1
	sums[n] = nums[n]
	for i := n - 1; i >= 0; i-- {
		sums[i] = sums[i+1] + nums[i]
	}
	return dfs(0, 0, S, nums, sums)
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
