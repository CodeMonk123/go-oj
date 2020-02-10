// 给定一个未排序的整数数组，找到最长递增子序列的个数。

// 示例 1:

// 输入: [1,3,5,4,7]
// 输出: 2
// 解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
// 示例 2:

// 输入: [2,2,2,2,2]
// 输出: 5
// 解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
// 注意: 给定的数组长度不超过 2000 并且结果一定是32位有符号整数。

package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := make([]int, len(nums))
	count := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		length[i], count[i] = 1, 1
	}
	res := 0
	max := 1
	for i := 1; i < len(nums); i++ {
		temp := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if length[j]+1 > temp {
					temp = length[j] + 1
					count[i] = count[j]
				} else if length[j]+1 == temp {
					count[i] += count[j]
				}
			}
		}
		length[i] = temp
		if temp > max {
			max = temp
			res = count[i]
		} else if temp == max {
			res += count[i]
		}
	}
	fmt.Println(length)
	fmt.Println(count)
	if max == 1 {
		return len(nums)
	}
	return res
}

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}))
	fmt.Println(findNumberOfLIS([]int{3, 1, 2}))
}
