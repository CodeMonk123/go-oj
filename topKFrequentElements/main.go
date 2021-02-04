package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	key       int
	frequence int
}

type Pairs []Pair

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i].frequence > p[j].frequence
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v] += 1
	}

	var paris []Pair

	for key, v := range counts {
		paris = append(paris, Pair{
			key:       key,
			frequence: v,
		})
	}

	sort.Sort(Pairs(paris))

	var result []int

	for i := 0; i < k; i++ {
		result = append(result, paris[i].key)
	}

	return result

}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}
