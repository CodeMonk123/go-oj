package main

import (
	"fmt"
	"sort"
)

// 给你一个整数数组 nums 和一个正整数 k，请你判断是否可以把这个数组划分成一些由 k 个连续数字组成的集合。
// 如果可以，请返回 True；否则，返回 False。
//
//
//
// 示例 1：
//
// 输入：nums = [1,2,3,3,4,4,5,6], k = 4
// 输出：true
// 解释：数组可以分成 [1,2,3,4] 和 [3,4,5,6]。
// 示例 2：
//
// 输入：nums = [3,2,1,2,3,4,3,4,5,9,10,11], k = 3
// 输出：true
// 解释：数组可以分成 [1,2,3] , [2,3,4] , [3,4,5] 和 [9,10,11]。
// 示例 3：
//
// 输入：nums = [3,3,2,2,1,1], k = 3
// 输出：true
// 示例 4：
//
// 输入：nums = [1,2,3,4], k = 3
// 输出：false
// 解释：数组不能分成几个大小为 3 的子数组。
//
//
// 提示：
//
// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^9
// 1 <= k <= nums.length

type SortableArray []int

func (s SortableArray) Len() int {
	return len(s)
}

func (s SortableArray) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortableArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func getSortedKeys(m map[int]int) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Sort(SortableArray(keys))
	return keys
}

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
	numCount := make(map[int]int)
	for _, v := range nums {
		if _, ok := numCount[v]; !ok {
			numCount[v] = 1
		} else {
			numCount[v]++
		}
	}

	keys := getSortedKeys(numCount)
	// fmt.Println(keys)
	// fmt.Println(numCount)
	for i := 0; i < len(keys); {
		// fmt.Println("start key is", keys[i])
		currentKey := keys[i]
		temp := currentKey
		for j := i + 1; j < i+k; j++ {
			if keys[j] != temp+1 {
				// fmt.Println("Key check failed")
				return false
			} else {
				temp++
			}
		}

		nextKey := -1
		for j := i; j < i+k; j++ {
			v, ok := numCount[keys[j]]
			if !ok {
				// fmt.Println("map failed")
				return false
			}
			if v < 1 {
				// fmt.Println("count failed")
				return false
			}
			// fmt.Printf("count[%d]=%d\n", keys[j], v)
			numCount[keys[j]]--
			if v > 1 && nextKey == -1 {
				nextKey = j
				// fmt.Printf("set next key: %d\n", keys[j])
			}
		}

		if nextKey != -1 {
			i = nextKey
		} else {
			i += k
		}
	}
	return true

}

func main() {
	fmt.Println(isPossibleDivide([]int{1, 2, 3, 3, 4, 4, 5, 6}, 4))
}
