package main

import (
	"container/heap"
	"fmt"
)

type heapArray []int

func (h heapArray) Len() int {
	return len(h)
}

func (h heapArray) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h heapArray) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heapArray) Push(x interface{}) {
	val := x.(int)
	h = append(h, val)
}

func (h heapArray) Pop() interface{} {
	return h[:len(h)-1]
}

func main() {
	a := []int{2, 3, 5, 6, 1, 3, 7, 9, 5, 3}
	heap.Init(heapArray(a))

	for len(a) > 0 {
		fmt.Println(a[0])
		a = heap.Pop(heapArray(a)).(heapArray)
	}
}
