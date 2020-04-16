package main

import "fmt"

// 给你一个由若干 0 和 1 组成的二维网格 grid，请你找出边界全部由 1 组成的最大 正方形 子网格，并返回该子网格中的元素数量。如果不存在，则返回 0。
//
//
//
// 示例 1：
//
// 输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
// 输出：9
// 示例 2：
//
// 输入：grid = [[1,1,0,0]]
// 输出：1
//
//
// 提示：
//
// 1 <= grid.length <= 100
// 1 <= grid[0].length <= 100
// grid[i][j] 为 0 或 1

func largest1BorderedSquare(grid [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	height, width := len(grid), len(grid[0])
	left := make([][]int, height)
	up := make([][]int, height)
	for i := range left {
		left[i] = make([]int, width)
		up[i] = make([]int, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == 1 {
				if j == 0 {
					left[i][j] = 1
				} else {
					left[i][j] = 1 + left[i][j-1]
				}
				if i == 0 {
					up[i][j] = 1
				} else {
					up[i][j] = 1 + up[i-1][j]
				}
			}
		}
	}
	// fmt.Println(left)
	// fmt.Println(up)
	res := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			l := min(left[i][j], up[i][j])
			for k := l; k > 0; k-- {
				if left[i+1-k][j] >= k && up[i][j+1-k] >= k {
					res = max(res, k)
				}
			}
		}
	}

	return res * res

}

func main() {
	grid := [][]int{
		{0, 1, 1, 1, 1, 0},
		{1, 1, 0, 1, 1, 0},
		{1, 1, 0, 1, 0, 1},
		{1, 1, 0, 1, 1, 1},
		{1, 1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1}}

	fmt.Println(largest1BorderedSquare(grid))
}
