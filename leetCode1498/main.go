package main

import (
	"fmt"
	"sort"
)

// 给你一个整数数组 nums 和一个整数 target 。
//
// 请你统计并返回 nums 中能满足其最小元素与最大元素的 和 小于或等于 target 的 非空 子序列的数目。
//
// 由于答案可能很大，请将结果对 10^9 + 7 取余后返回。
//
//  
//
// 示例 1：
//
// 输入：nums = [3,5,6,7], target = 9
// 输出：4
// 解释：有 4 个子序列满足该条件。
// [3] -> 最小元素 + 最大元素 <= target (3 + 3 <= 9)
// [3,5] -> (3 + 5 <= 9)
// [3,5,6] -> (3 + 6 <= 9)
// [3,6] -> (3 + 6 <= 9)
// 示例 2：
//
// 输入：nums = [3,3,6,8], target = 10
// 输出：6
// 解释：有 6 个子序列满足该条件。（nums 中可以有重复数字）
// [3] , [3] , [3,3], [3,6] , [3,6] , [3,3,6]
// 示例 3：
//
// 输入：nums = [2,3,3,4,6,7], target = 12
// 输出：61
// 解释：共有 63 个非空子序列，其中 2 个不满足条件（[6,7], [7]）
// 有效序列总数为（63 - 2 = 61）
// 示例 4：
//
// 输入：nums = [5,2,4,1,7,6,8], target = 16
// 输出：127
// 解释：所有非空子序列都满足条件 (2^7 - 1) = 127
//  
//
// 提示：
//
// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^6
// 1 <= target <= 10^6
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type SortableIntIndices []int

func (s SortableIntIndices) Len() int {
	return len([]int(s))
}

func (s SortableIntIndices) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortableIntIndices) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func pow2(n int) int {
	if n == 0 {
		return 1
	}
	res := 1
	step := 2
	for n > 0 {
		if n&1 == 1 {
			res = res * step % (1E9 + 7)
		}
		step = step * step % (1E9+7)
		n >>= 1
	}

	return res % (1E9 + 7)
}


func numSubseq(nums []int, target int) int {
	sort.Sort(SortableIntIndices(nums))
	res := 0
	i := 0
	j := len(nums) - 1

	for i <= j {
		if nums[i] + nums[j] > target {
			j--
		} else {
			res += pow2(j-i)
			i++
		}
 	}
	return res % 1000000007
}

func main() {
	fmt.Println(numSubseq([]int{5,2,4,1,7,6,8}, 16))
	fmt.Println(numSubseq([]int{2,3,3,4,6,7}, 12))
	fmt.Println(numSubseq([]int{27,21,14,2,15,1,19,8,12,24,21,8,12,10,11,30,15,18,28,14,26,9,2,24,23,11,7,12,9,17,30,9,28,2,14,22,19,19,27,6,15,12,29,2,30,11,20,30,21,20,2,22,6,14,13,19,21,10,18,30,2,20,28,22}, 31))
}
