package main

import "fmt"

func find(start int ,nums1 []int, nums2 []int, suffix1 map[int]int, suffix2 map[int]int) int {
	res := 0
	for {
		if start == len(nums1) {
			return res
		}
		if _, ok:=suffix1[start]; !ok {
			res += nums1[start]
			start ++
		} else {
			res1 := find(start+1, nums1, nums2, suffix1, suffix2)  + nums1[start]
			res2 := find(suffix1[start], nums2, nums1, suffix2, suffix1)  + nums1[start]
			if res1 > res2 {
				res +=res1
			} else {
				res +=res2
			}
			return res
		}

	}

}

func maxSum(nums1 []int, nums2 []int) int {
	index := make(map[int]int)
	suffix1 := make(map[int]int)
	suffix2 := make(map[int]int)

	for i, v := range nums1 {
		index[v] = i
	}

	for i, v := range nums2 {
		if j, ok := index[v]; ok {
			suffix1[j] = i + 1
			suffix2[i] = j + 1
		}
	}

	res1 := find(0, nums1, nums2, suffix1, suffix2)
	res2 := find(0, nums2, nums1, suffix2, suffix1)

	if res1 > res2 {
		return res1
	} else {
		return res2
	}
 }

func main() {

	nums1 := []int{1,4,5,8,9,11,19}
	nums2 := []int{2,3,4,11,12}
	fmt.Println(maxSum(nums1, nums2))
}
