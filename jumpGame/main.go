// 给你一个整数数组 arr 和一个整数 d 。每一步你可以从下标 i 跳到：

// i + x ，其中 i + x < arr.length 且 0 < x <= d 。
// i - x ，其中 i - x >= 0 且 0 < x <= d 。
// 除此以外，你从下标 i 跳到下标 j 需要满足：arr[i] > arr[j] 且 arr[i] > arr[k] ，其中下标 k 是所有 i 到 j 之间的数字（更正式的，min(i, j) < k < max(i, j)）。

// 你可以选择数组的任意下标开始跳跃。请你返回你 最多 可以访问多少个下标。

// 请注意，任何时刻你都不能跳到数组的外面。

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func next(current int, arr []int, d int) []int {
	nextPos := []int{}
	left := -1
	for i := current - 1; i >= max(0, current-d); i-- {
		if arr[i] < arr[current] {
			left = i
		} else {
			break
		}
	}
	if left != -1 {
		for i := left; i < current; i++ {
			nextPos = append(nextPos, i)
		}
	}

	right := -1
	for i := current + 1; i <= min(len(arr)-1, current+d); i++ {
		if arr[i] < arr[current] {
			right = i
		} else {
			break
		}
	}

	if right != -1 {
		for i := current + 1; i <= right; i++ {
			nextPos = append(nextPos, i)
		}
	}
	return nextPos
}

func dfs(current int, arr []int, d int, known *map[int]int) int {
	if (*known)[current] != -1 {
		return (*known)[current]
	}
	nextPos := next(current, arr, d)
	//fmt.Printf("Current: %v, Next: %v\n", current, nextPos)
	if len(nextPos) == 0 {
		(*known)[current] = 1
		return 1
	}
	maxStep := 0
	for _, v := range nextPos {
		temp := dfs(v, arr, d, known)
		if temp > maxStep {
			maxStep = temp
		}
	}
	(*known)[current] = maxStep + 1
	return maxStep + 1
}

func maxJumps(arr []int, d int) int {
	maxStep := 1
	known := make(map[int]int)
	for i := range arr {
		known[i] = -1
	}
	for i := range arr {
		temp := dfs(i, arr, d, &known)
		if temp > maxStep {
			maxStep = temp
		}
	}
	return maxStep
}

func main() {
	fmt.Println(maxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2))
	fmt.Println(maxJumps([]int{7, 6, 5, 4, 3, 2, 1}, 1))
}
