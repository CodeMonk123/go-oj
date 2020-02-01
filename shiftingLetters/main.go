package main

import "fmt"

func shift(c byte, i int) byte {
	return byte('a' + (int(c-'a')+i)%26)
}

func calcFinalShiftTimes(shifts []int) []int {
	result := make([]int, len(shifts))
	for i := len(shifts) - 1; i >= 0; i-- {
		if i == len(shifts)-1 {
			result[i] = shifts[i]
		} else {
			result[i] = result[i+1] + shifts[i]
		}
	}
	return result
}

func shiftingLetters(S string, shifts []int) string {
	res := []byte(S)
	finalShifts := calcFinalShiftTimes(shifts)
	for i, c := range res {
		res[i] = shift(c, finalShifts[i])
	}

	return string(res)
}

func main() {
	fmt.Println(shiftingLetters("abc", []int{3, 5, 9}))
}
