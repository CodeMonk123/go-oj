// 你有 4 张写有 1 到 9 数字的牌。你需要判断是否能通过 *，/，+，-，(，) 的运算得到 24。

// 示例 1:

// 输入: [4, 1, 8, 7]
// 输出: True
// 解释: (8-4) * (7-1) = 24
// 示例 2:

// 输入: [1, 2, 1, 2]
// 输出: False
// 注意:

// 除法运算符 / 表示实数除法，而不是整数除法。例如 4 / (1 - 2/3) = 12 。
// 每个运算符对两个数进行运算。特别是我们不能用 - 作为一元运算符。例如，[1, 1, 1, 1] 作为输入时，表达式 -1 - 1 - 1 - 1 是不允许的。
// 你不能将数字连接在一起。例如，输入为 [1, 2, 1, 2] 时，不能写成 12 + 12 。

package main

import "fmt"

import "math"

//3个运算符号位置，4个运算符号，4^3=64种可能
//括号只控制运算顺序，有3!=6种运算顺序
//深度搜索这6 * 64 = 384种可能
const (
	Plus = iota
	Minus
	Mul
	Div
)

func step(a, b float64, op int) float64 {
	switch op {
	case Plus:
		return a + b
	case Minus:
		return a - b
	case Div:
		return a / b
	case Mul:
		return a * b
	}
	return -1
}

func calc(nums []float64, op []int) bool {
	//fmt.Println(nums, op)
	res1 := step(step(nums[0], nums[1], op[0]), step(nums[2], nums[3], op[2]), op[1])
	if math.Abs(res1-24.0) < 0.0001 {
		return true
	}
	res2 := step(step(step(nums[0], nums[1], op[0]), nums[2], op[1]), nums[3], op[2])
	if math.Abs(res2-24.0) < 0.0001 {
		return true
	}
	res3 := step(nums[0], step(nums[1], step(nums[2], nums[3], op[2]), op[1]), op[0])
	if math.Abs(res3-24.0) < 0.0001 {
		return true
	}
	res4 := step(step(nums[0], step(nums[1], nums[2], op[1]), op[0]), nums[3], op[2])
	if math.Abs(res4-24.0) < 0.0001 {
		return true
	}
	res5 := step(step(step(nums[1], nums[2], op[1]), nums[3], op[2]), nums[0], op[0])
	//fmt.Println(res1, res2, res3, res4, res5)
	return math.Abs(res5-24.0) < 0.0001

}

func dfs(nums []float64, op []int, rest []int) bool {
	for i, v := range rest {
		newNums := append([]float64{}, nums...)
		newNums = append(newNums, float64(v))
		if len(newNums) == 4 {
			return calc(newNums, op)
		} else {
			newRest := append([]int{}, rest[:i]...)
			newRest = append(newRest, rest[i+1:]...)
			for j := 0; j < 4; j++ {
				newOp := append([]int{}, op...)
				newOp = append(newOp, j)
				if dfs(newNums, newOp, newRest) {
					return true
				}
			}
		}
	}
	return false
}

func judgePoint24(nums []int) bool {
	return dfs([]float64{}, []int{}, nums)
}

func main() {
	fmt.Println(judgePoint24([]int{3, 3, 8, 8}))
	fmt.Println(judgePoint24([]int{1, 1, 7, 7}))
}
