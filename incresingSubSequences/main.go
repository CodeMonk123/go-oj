// Given an integer array, your task is to find all the different possible
//increasing subsequences of the given array, and the length of an increasing subsequence should be at least 2.

// Example:

// Input: [4, 6, 7, 7]
// Output: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]

// Note:

// The length of the given array will not exceed 15.
// The range of integer in the given array is [-100,100].
// The given array may contain duplicates, and two equal integers
// should also be considered as a special case of increasing sequence.

package main

import "fmt"

func dfs(index int, current []int, res *[][]int, nums []int) {
	current = append(current, nums[index])
	//fmt.Println("Current: ", current, "Index: ", index)
	if len(current) >= 2 {
		newSeq := append([]int{}, current...)
		*res = append(*res, newSeq)
	}
	for i := index + 1; i < len(nums); i++ {
		if nums[i] >= nums[index] {
			dfs(i, current, res, nums)
		}
	}
}

func unique(s [][]int) [][]int {
	res := [][]int{}
	exist := make(map[string]bool)
	for _, seq := range s {
		key := fmt.Sprintf("%v", seq)
		if !exist[key] {
			res = append(res, seq)
			exist[key] = true
		}
	}
	return res
}

func findSubsequences(nums []int) [][]int {
	temp := [][]int{}
	for i := range nums {
		dfs(i, []int{}, &temp, nums)
	}
	//fmt.Println(temp)
	result := unique(temp)
	return result
}

func main() {
	fmt.Println(findSubsequences([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
}
