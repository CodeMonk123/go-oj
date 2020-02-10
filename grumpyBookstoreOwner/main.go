// 今天，书店老板有一家店打算试营业 customers.length 分钟。每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分钟结束后离开。

// 在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。

// 书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 X 分钟不生气，但却只能使用一次。

// 请你返回这一天营业下来，最多有多少客户能够感到满意的数量。

// 示例：

// 输入：customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
// 输出：16
// 解释：
// 书店老板在最后 3 分钟保持冷静。
// 感到满意的最大客户数量 = 1 + 1 + 1 + 1 + 7 + 5 = 16.

// 提示：

// 1 <= X <= customers.length == grumpy.length <= 20000
// 0 <= customers[i] <= 1000
// 0 <= grumpy[i] <= 1

package main

import (
	"fmt"
	"math"
)

func sum(a ...int) int {
	res := 0
	for _, v := range a {
		res += v
	}

	return res
}

func max(a []int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}

func maxSatisfied(customers []int, grumpy []int, X int) int {
	if X >= len(customers) {
		return sum(customers...) //所有顾客都满意
	}
	all := 0
	for i, v := range customers {
		if !(i < len(grumpy) && grumpy[i] == 1) {
			all += v
		}
	}
	// all 是不考虑X的情况下，满意的顾客数
	contribution := make([]int, len(customers))
	temp := 0
	for i := 0; i < X; i++ {
		if i < len(grumpy) && grumpy[i] == 1 {
			temp += customers[i]
		}
	}
	contribution[0] = temp
	for i := 1; i < len(contribution); i++ {
		if i-1 < len(grumpy) && grumpy[i-1] == 1 {
			temp -= customers[i-1]
		}
		end := i + X - 1
		if end < len(grumpy) && grumpy[end] == 1 {
			temp += customers[end]
		}
		contribution[i] = temp
	}
	fmt.Println(contribution)
	return all + max(contribution)
}

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
	fmt.Println(maxSatisfied([]int{5, 8}, []int{0, 1}, 1))
}
