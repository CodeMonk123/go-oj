package main

import "fmt"

func countTriplets(arr []int) int {
	temp := make([]int, len(arr))
	temp[0] = arr[0]
	for i:=1; i < len(arr); i++ {
		temp[i] = arr[i] ^ temp[i-1]
	}

	count := 0
	for i:=0; i < len(arr); i++ {
		for j:=i+1; j < len(arr); j++ {
			for k := j; k  < len(arr); k++ {
				a := temp[j-1]
				if i > 0 {
					a ^= temp[i-1]
				}
				b := temp[k]^temp[j-1]
				if a == b {
					count ++
				}
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countTriplets([]int{2,3,1,6,7}))
}
