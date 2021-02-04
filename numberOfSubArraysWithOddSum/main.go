package main

import "fmt"

func numOfSubarrays(arr []int) int {
	var subArrayWithOddSumCount []int
	for range arr {
		subArrayWithOddSumCount = append(subArrayWithOddSumCount, 0)
	}

	if arr[0] % 2 != 0 {
		subArrayWithOddSumCount[0] = 1
	}

	for i:=1; i < len(arr); i++ {
		total := i+1
		if arr[i] % 2 == 0 {
			subArrayWithOddSumCount[i] = subArrayWithOddSumCount[i-1]
		} else {
			subArrayWithOddSumCount[i] = total - subArrayWithOddSumCount[i-1]
		}
	}

	sum := 0
	for _, v := range subArrayWithOddSumCount {
		sum += v
	}
	return sum

}

func main() {
	fmt.Println(numOfSubarrays([]int{7}))
}
