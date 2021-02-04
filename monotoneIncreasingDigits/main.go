package main

import "fmt"

//给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
//
//（当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）
//
//示例 1:
//
//输入: N = 10
//输出: 9
//示例 2:
//
//输入: N = 1234
//输出: 1234
//示例 3:
//
//输入: N = 332
//输出: 299

func monotoneIncreasingDigits(N int) int {
	var digits []int
	temp := N
	for temp != 0 {
		digits = append(digits, temp % 10)
		temp /= 10
	}

	for i,j := 0, len(digits) - 1; i < j; i,j = i + 1, j- 1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	if digits == nil {
		return 0
	}

	mono := true
	for i := 0 ;i < len(digits); i++ {
		if !mono {
			digits[i] = 9
			continue
		}
		if i != len(digits) - 1 && digits[i] > digits[i+1] {
			if i == 0 {
				digits[i] -= 1
			} else {
				for j := i - 1; j >= 0; j-- {
					if digits[j] == digits[j+1] {
						digits[j+1] = 9
						if j == 0 {
							digits[j] -= 1
						}
					} else {
						digits[j+1] -= 1
						break
					}
				}
			}
			mono = false
		}
	}

	if mono {
		return N
	}

	var result int
	step := 1
	for i := len(digits) - 1; i >= 0; i-- {
		result += digits[i] * step
		step *= 10
	}

	return result
}

func main() {
	fmt.Println(monotoneIncreasingDigits(332))
	fmt.Println(monotoneIncreasingDigits(1122334))
	fmt.Println(monotoneIncreasingDigits(10000))
}

