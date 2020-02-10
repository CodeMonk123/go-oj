// 实现 pow(x, n) ，即计算 x 的 n 次幂函数。

// 示例 1:

// 输入: 2.00000, 10
// 输出: 1024.00000
// 示例 2:

// 输入: 2.10000, 3
// 输出: 9.26100
// 示例 3:

// 输入: 2.00000, -2
// 输出: 0.25000
// 解释: 2-2 = 1/22 = 1/4 = 0.25
// 说明:

// -100.0 < x < 100.0
// n 是 32 位有符号整数，其数值范围是 [−2^31, 2^31 − 1] 。

package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n > 0 {
		res := 1.0
		temp := x
		for i := 0; i < 32; i++ {
			t := n & 1
			n >>= 1
			if t == 1 {
				res *= temp
			}
			temp *= temp
		}
		return res
	} else {
		m := -(int64(n))
		res := 1.0
		temp := x
		for i := 0; i < 33; i++ {
			t := m & 1
			m >>= 1
			if t == 1 {
				res /= temp
			}
			temp *= temp
		}
		return res
	}
}

func main() {
	fmt.Println(myPow(2.1, 3))
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(3, -3))
}
