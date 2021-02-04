package main

import "fmt"

func find(append int, depth int, sum int, num int) int {
	if append == 1 {
		sum += 1 << depth
		if sum > num {
			return 0
		}
		return find(0, depth + 1, sum, num)
	} else {
		if 1 << depth > num {
			return 1
		}
		return find(1, depth+1, sum, num) + find(0, depth+1,sum,num)
	}


}

// 超时
func findIntegers1(num int) int {
	return find(1,  0, 0, num) + find(0,0,0,num)
}

// Accepted

func findIntegers(num int) int {
	f := make([]int, 32)
	f[0] = 1
	f[1] = 2
	for i := 2; i < 32; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	i := 30
	res := 0
	prevBit := 0
	for i >= 0 {
		if (num & (1 << i)) != 0 {
			res += f[i]
			if prevBit == 1 {
				res--
				break
			}
			prevBit = 1
		} else {
			prevBit = 0
		}
		i--
	}
	return res + 1
}


func main() {
	fmt.Println(findIntegers1(949344261))
	fmt.Println(findIntegers(949344261))
}

