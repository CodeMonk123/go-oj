package main

import (
	"fmt"
)

// 你现在手里有一份大小为 N x N 的「地图」（网格） grid，上面的每个「区域」（单元格）都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地，请你找出一个海洋区域，这个海洋区域到离它最近的陆地区域的距离是最大的。
// 我们这里说的距离是「曼哈顿距离」（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个区域之间的距离是 |x0 - x1| + |y0 - y1| 。
// 如果我们的地图上只有陆地或者海洋，请返回 -1。
// 示例 1：
// 输入：[[1,0,1],[0,0,0],[1,0,1]]
// 输出：2
// 解释：
// 海洋区域 (1, 1) 和所有陆地区域之间的距离都达到最大，最大距离为 2。
// 示例 2：
// 输入：[[1,0,0],[0,0,0],[0,0,0]]
// 输出：4
// 解释：
// 海洋区域 (2, 2) 和所有陆地区域之间的距离都达到最大，最大距离为 4。
// 提示：
//
// 1 <= grid.length == grid[0].length <= 100
// g
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/as-far-from-land-as-possible
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 从陆地出发反向进行广度优先搜索
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func update(grid [][]int, x,y int, distance [][]int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				newDis := abs(x-i) + abs(y-j)
				if newDis < distance[i][j] || distance[i][j] == -1{
					distance[i][j] = newDis
				}
			}
		}
	}
}

func maxDistance(grid [][]int) int {
	height, width := len(grid), len(grid[0])
	var distance [][]int
	for i := 0; i < height; i++ {
		row := make([]int, width)
		for i := range row {
			row[i] = -1
		}
		distance = append(distance, row)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				update(grid, i,j, distance)
			}
		}
	}

	return func(matrix [][]int) int{
		max := -100
		for i := range matrix {
			for j:= range matrix[i] {
				if matrix[i][j] > max {
					max = matrix[i][j]
				}
			}
		}
		return max
	}(distance)

}

func main() {
	grid := [][]int{{1,0,0},{0,0,0},{0,0,0}}
	fmt.Println(maxDistance(grid))
}
