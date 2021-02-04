package main

// 给你一个字符串 s 和一个 长度相同 的整数数组 indices 。
//
// 请你重新排列字符串 s ，其中第 i 个字符需要移动到 indices[i] 指示的位置。
//
// 返回重新排列后的字符串。
//
// 输入：s = "codeleet", indices = [4,5,6,7,0,2,1,3]
// 输出："leetcode"
// 解释：如图所示，"codeleet" 重新排列后变为 "leetcode" 。

func partition(array []int, begin, end int, res []byte) int{
	pivot := array[begin]
	pivot2 := res[begin]
	partitonPoint := begin
	for i:=begin+1; i <= end; i++ {
		if array[i] < pivot {
			partitonPoint++
			if i != partitonPoint {
				temp1 := array[i]
				temp2 := res[i]
				array[i] = array[partitonPoint]
				res[i] = res[partitonPoint]
				array[partitonPoint] = temp1
				res[partitonPoint] = temp2
			}
		}
	}
	array[begin] = array[partitonPoint]
	array[partitonPoint] = pivot
	res[begin] = res[partitonPoint]
	res[partitonPoint] = pivot2
	return partitonPoint
}

func  quickSort(array []int, begin, end int, res []byte) {
	if begin < end {
		partitionPoint := partition(array, begin, end, res)
		quickSort(array, begin, partitionPoint, res)
		quickSort(array, partitionPoint+1, end, res)
	}
}


func restoreString(s string, indices []int) string {
	var input = []byte(s)
	quickSort(indices, 0, len(indices)-1, input)
	return string(input)
}

func main() {
	println(restoreString("codeleet", []int{4,5,6,7,0,2,1,3}))
}