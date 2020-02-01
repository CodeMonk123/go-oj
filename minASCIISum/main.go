package main

import "fmt"

/*
Given two strings s1, s2, find the lowest ASCII sum of deleted characters to make two strings equal.

Example 1:
Input: s1 = "sea", s2 = "eat"
Output: 231
Explanation: Deleting "s" from "sea" adds the ASCII value of "s" (115) to the sum.
Deleting "t" from "eat" adds 116 to the sum.
At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum possible to achieve this.
Example 2:
Input: s1 = "delete", s2 = "leet"
Output: 403
Explanation: Deleting "dee" from "delete" to turn the string into "let",
adds 100[d]+101[e]+101[e] to the sum.  Deleting "e" from "leet" adds 101[e] to the sum.
At the end, both strings are equal to "let", and the answer is 100+101+101+101 = 403.
If instead we turned both strings into "lee" or "eet", we would get answers of 433 or 417, which are higher.
Note:

0 < s1.length, s2.length <= 1000.
All elements of each string will have an ASCII value in [97, 122].
*/

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}

func minimumDeleteSum(s1 string, s2 string) int {
	all := 0
	for _, c := range []byte(s1) {
		all += int(c)
	}

	for _, c := range []byte(s2) {
		all += int(c)
	}

	l1, l2 := len([]byte(s1)), len([]byte(s2))
	maxSum := [][]int{}
	for i := 0; i < l1; i++ {
		maxSum = append(maxSum, make([]int, l2))
	}
	//fmt.Println(maxSum)
	if s1[0] == s2[0] {
		maxSum[0][0] = 2 * int(s1[0])
	} else {
		maxSum[0][0] = 0
	}

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			temp := maxSum[i][j]
			if i > 0 && j > 0 {
				if s1[i] == s2[j] {
					temp = max(maxSum[i-1][j], maxSum[i][j-1], 2*int(s1[i])+maxSum[i-1][j-1])
				} else {
					temp = max(maxSum[i-1][j], maxSum[i][j-1])
				}
			} else if i > 0 {
				if s1[i] == s2[j] {
					temp = max(maxSum[i-1][j], 2*int(s1[i]))
				} else {
					temp = maxSum[i-1][j]
				}
			} else if j > 0 {
				if s1[i] == s2[j] {
					temp = max(maxSum[i][j-1], 2*int(s1[i]))
				} else {
					temp = maxSum[i][j-1]
				}
			}
			maxSum[i][j] = temp
		}
	}
	return all - maxSum[l1-1][l2-1]
}

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
	fmt.Println(minimumDeleteSum("delete", "leet"))
	fmt.Println(minimumDeleteSum("a", "at"))
}
