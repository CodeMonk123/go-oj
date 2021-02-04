package main

import (
	"fmt"
	"strings"
)

func canTransform(start string, end string) bool {
	s1 := []rune(start)
	s2 := []rune(end)
	var p, q int
	for p, q = 0, 0; p < len(start) && q < len(end); {
		if s1[p] == 'X' {
			p++
			continue
		}
		if s2[q] == 'X' {
			q++
			continue
		}

		if s1[p] != s2[q] {
			return false
		}

		if s1[p] == 'L' && p < q || s1[p] == 'R' && p > q {
			return false
		}
		p++
		q++
	}
	if p < len(start) {
		return !strings.ContainsAny(string(s1[p:]), "LR")
	}

	if q < len(end) {
		return !strings.ContainsAny(string(s2[q:]), "LR")
	}

	return true
}

func main() {
	fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))
	fmt.Println(canTransform("LXXLXRLXXL", "XLLXRXLXLX"))
}
