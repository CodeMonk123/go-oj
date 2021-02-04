package main

import "fmt"

func myAbs(x int ) int{
	if x < 0 {
		return -x
	}
	return x
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			for k := j+1; k < len(arr); k++ {
				if myAbs(arr[i] - arr[j]) <= a && myAbs(arr[j] - arr[k]) <= b && myAbs(arr[i] - arr[k]) <= c {
					res ++
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(countGoodTriplets([]int{3,0,1,1,9,7}, 7, 2, 3))
}
