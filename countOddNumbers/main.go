package main

import "fmt"

func countOdds(low int, high int) int {
	count := high - low + 1
	if count % 2 == 0{
		return count / 2
	} else {
		if low % 2 == 0 {
			count -= 1
		} else {
			count += 1
		}
		return count / 2
	}
}

func main() {
	fmt.Println(countOdds(14,17))
}
