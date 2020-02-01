/*
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example:

X X X X
X O O X
X X O X
X O X X
After running your function, the board should be:

X X X X
X X X X
X X X X
X O X X
Explanation:

Surrounded regions shouldn’t be on the border,
which means that any 'O' on the border of the board are not flipped to 'X'.
Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'.
Two cells are connected if they are adjacent cells connected horizontally or vertically.

*/

// 正向搜索超时
package main

import "fmt"

// type position struct {
// 	x, y int
// }

// func contains(positions []position, p position) bool {
// 	for _, pos := range positions {
// 		if p == pos {
// 			return true
// 		}
// 	}
// 	return false
// }

// func dfs(board [][]byte, visited *[]position, x, y int) bool {
// 	height, width := len(board), len(board[0])
// 	if x == 0 || y == 0 || x == height-1 || y == width-1 {
// 		return false
// 	} else {
// 		for _, pos := range []position{position{x, y + 1}, position{x, y - 1}, position{x + 1, y}, position{x - 1, y}} {
// 			m, n := pos.x, pos.y
// 			if m < height && n < width && board[m][n] == 'O' && !contains(*visited, pos) {
// 				// fmt.Println("Find O in ", m, n, "Visited:", *visited)
// 				*visited = append(*visited, pos)
// 				if !dfs(board, visited, m, n) {
// 					return false
// 				}
// 			}
// 		}
// 		return true
// 	}
// }

// func flip(board [][]byte, positions []position) {
// 	for _, pos := range positions {
// 		board[pos.x][pos.y] = 'X'
// 	}
// }

// func solve(board [][]byte) {
// 	if len(board) <= 1 || len(board[0]) <= 1 {
// 		return
// 	}
// 	height, width := len(board), len(board[0])
// 	for i := 1; i < height; i++ {
// 		for j := 1; j < width; j++ {
// 			if board[i][j] == 'O' {
// 				visited := []position{position{i, j}}
// 				if dfs(board, &visited, i, j) {
// 					// fmt.Println("Dfs return true, visited", visited)
// 					flip(board, visited)
// 				}
// 			}
// 		}
// 	}
// }
type position struct {
	x, y int
}

func dfs(board [][]byte, x, y int) {
	height, width := len(board), len(board[0])
	for _, pos := range []position{position{x - 1, y}, position{x + 1, y}, position{x, y + 1}, position{x, y - 1}} {
		m, n := pos.x, pos.y
		if m < 0 || m >= height || n < 0 || n >= width {
			continue
		} else {
			if board[m][n] == 'O' {
				board[m][n] = '-'
				dfs(board, m, n)
			}
		}
	}
}

func flip(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '-' {
				board[i][j] = 'O'
			}
		}
	}
}

func solve(board [][]byte) {
	if len(board) <= 1 || len(board[0]) <= 1 {
		return
	} else {
		height, width := len(board), len(board[0])
		for j := 0; j < width; j++ {
			if board[0][j] == 'O' {
				board[0][j] = '-'
				dfs(board, 0, j)
			}
			if board[height-1][j] == 'O' {
				board[height-1][j] = '-'
				dfs(board, height-1, j)
			}
		}
		for i := 0; i < height; i++ {
			if board[i][0] == 'O' {
				board[i][0] = '-'
				dfs(board, i, 0)
			}
			if board[i][width-1] == 'O' {
				board[i][width-1] = '-'
				dfs(board, i, width-1)
			}
		}
		flip(board)
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'O', 'X', 'O'},
		{'X', 'O', 'X', 'O', 'X'},
		{'X', 'X', 'O', 'O', 'X'},
		{'X', 'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X', 'X'},
	}
	solve(board)
	for _, line := range board {
		for _, pos := range line {
			fmt.Printf("%c, ", pos)
		}
		fmt.Println()
	}
}
