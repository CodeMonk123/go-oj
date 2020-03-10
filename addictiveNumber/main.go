package main

import (
	"fmt"
	"strconv"
)

// 累加数是一个字符串，组成它的数字可以形成累加序列。
//
// 一个有效的累加序列必须至少包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。
//
// 给定一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是累加数。
//
// 说明: 累加序列里的数不会以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。
//
// 示例 1:
//
// 输入: "112358"
// 输出: true
// 解释: 累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
// 示例 2:
//
// 输入: "199100199"
// 输出: true
// 解释: 累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199

func sum(num1, num2 string) string {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	res := n1 + n2
	return fmt.Sprintf("%d", res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// raw 原始字符串
// [begin,mid) 给出了当前搜索的第一个子串
// [mid , end) 给出了当前搜索的第二个子串
func dfs(begin, mid, end int, raw string) bool {
	if end == len(raw) {
		return true
	}
	num1 := raw[begin:mid]
	num2 := raw[mid:end]
	num3 := sum(num1, num2)
	if num3 == raw[end:min(len(raw), end+len(num3))] {
		//fmt.Println("Find valid:", num1, "+", num2, "=", num3)
		return dfs(mid, end, min(len(raw), end+len(num3)), raw)
	}
	//fmt.Println("Return false: ",num1, num2)
	return false
}

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	begin := 0
	mid := 1
	end := 2
	for end < len(num) {
		if dfs(begin, mid, end, num) {
			return true
		} else {
			if end < len(num)-2 && num[mid] != '0' {
				end += 1
			} else if num[begin] != '0' {
				mid += 1
				end = mid + 1
			} else {
				return false
			}
		}
	}
	return false
}

func main() {
	fmt.Println(isAdditiveNumber("025813"))
	fmt.Println(isAdditiveNumber("101"))
	fmt.Println(isAdditiveNumber("112358"))
	fmt.Println(isAdditiveNumber("199100199299"))
}
