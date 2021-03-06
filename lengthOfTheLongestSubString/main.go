package main

import "fmt"

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
//
// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
//
// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	subStringLength := make([]int, len(s))
	for i := range subStringLength {
		subStringLength[i] = 1
	}

	for i := 1; i < len(s); i++ {
		flag := true
		for j:=0; j < subStringLength[i-1]; j++ {
			if s[i-1-j] == s[i] {
				flag = false
				subStringLength[i] = j+1
				break
			}
		}
		if flag {
			subStringLength[i] = subStringLength[i-1] + 1
		}
	}
	return func(length []int) int {
		res := length[0]
		for _, v := range length {
			if v > res {
				res = v
			}
		}
		return res
	}(subStringLength)
}

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
