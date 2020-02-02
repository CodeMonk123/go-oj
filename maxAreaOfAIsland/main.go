// 给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

// 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

// 示例 1:

// [[0,0,1,0,0,0,0,1,0,0,0,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,1,1,0,1,0,0,0,0,0,0,0,0],
//  [0,1,0,0,1,1,0,0,1,0,1,0,0],
//  [0,1,0,0,1,1,0,0,1,1,1,0,0],
//  [0,0,0,0,0,0,0,0,0,0,1,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,0,0,0,0,0,0,1,1,0,0,0,0]]
// 对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

// 示例 2:

// [[0,0,0,0,0,0,0,0]]
// 对于上面这个给定的矩阵, 返回 0。

// 注意: 给定的矩阵grid 的长度和宽度都不超过 50。

package main

import "fmt"

func transform(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func dfs(grid [][]int, x, y int, visited *map[string]bool) int {
	area := 1
	height, width := len(grid), len(grid[0])
	for _, i := range []int{x - 1, x + 1} {
		j := y
		if i >= 0 && i < height && grid[i][j] == 1 && !(*visited)[transform(i, j)] {
			(*visited)[transform(i, j)] = true
			area += dfs(grid, i, j, visited)
		}
	}

	for _, j := range []int{y - 1, y + 1} {
		i := x
		if j >= 0 && j < width && grid[i][j] == 1 && !(*visited)[transform(i, j)] {
			(*visited)[transform(i, j)] = true
			area += dfs(grid, i, j, visited)
		}
	}

	return area
}

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	maxArea := 0
	height, width := len(grid), len(grid[0])
	visited := make(map[string]bool)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == 1 && !visited[transform(i, j)] {
				visited[transform(i, j)] = true
				area := dfs(grid, i, j, &visited)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func main() {
	grid := [][]int{
		{0, 0, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
	}

	fmt.Println(maxAreaOfIsland(grid))
}
