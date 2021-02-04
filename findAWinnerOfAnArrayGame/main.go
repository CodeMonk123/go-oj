package main

import "fmt"

func getWinner(arr []int, k int) int {
	pivot := arr[0]
	count := 0
	for j := 1; j < len(arr); j++ {
		if arr[j] < pivot {
			count++
			if count == k {
				return pivot
			}
		} else {
			break
		}
	}
	if count == len(arr) - 1 {
		return pivot
	}

	max := pivot
	for i := 1; i < len(arr); i++ {
		if arr[i] < max {
			continue
		} else {
			max = arr[i]
		}
		if k == 1 {
			return max
		}
		count = 0
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < max {
				count ++
				if count == k -1 {
					return max
				}
			} else {
				break
			}
		}
		if count == len(arr) - i - 1 {
			return max
		}
	}
	return -1
}

func main() {
	fmt.Println(getWinner([]int{1,25,35,42,68,70}, 1))
	fmt.Println(getWinner([]int{1,11,22,33,44,55,66,77,88,99},1E8))
}

