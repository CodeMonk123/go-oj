// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

package main

import "fmt"

func merge(nums []int, start, mid, end int, count *int) {
	p0 := start
	p1 := mid + 1
	q0 := mid
	q1 := end
	n := 0
	newNums := make([]int, end-start+1)
	i, j := p0, p1
	for i <= q0 && j <= q1 {
		if nums[i] <= nums[j] {
			newNums[n] = nums[i]
			i++
			n++
		} else {
			newNums[n] = nums[j]
			*count += (q0 + 1 - i)
			j++
			n++
		}
	}

	if i > q0 {
		for ; j <= q1; j++ {
			newNums[n] = nums[j]
			n++
		}
	}

	if j > q1 {
		for ; i <= q0; i++ {
			newNums[n] = nums[i]
			n++
		}
	}

	for i, v := range newNums {
		nums[i+start] = v
	}
}

func mergesort(nums []int, start, end int, count *int) {
	if start < end {
		mid := (start + end) / 2
		mergesort(nums, start, mid, count)
		mergesort(nums, mid+1, end, count)
		merge(nums, start, mid, end, count)
	}
}

func reversePairs(nums []int) int {
	res := 0
	mergesort(nums, 0, len(nums)-1, &res)
	fmt.Println(nums)
	return res
}

func main() {
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
}
