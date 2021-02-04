package main

import (
	"fmt"
)

// FIXME: 超时

// https://leetcode-cn.com/contest/weekly-contest-199/problems/string-compression-ii/
// //
// // 行程长度编码 是一种常用的字符串压缩方法，它将连续的相同字符（重复 2 次或更多次）替换为字符和表示字符计数的数字（行程长度）。例如，用此方法压缩字符串 "aabccc" ，将 "aa" 替换为 "a2" ，"ccc" 替换为` "c3" 。因此压缩后的字符串变为 "a2bc3" 。
// //
// // 注意，本问题中，压缩时没有在单个字符后附加计数 '1' 。
// //
// // 给你一个字符串 s 和一个整数 k 。你需要从字符串 s 中删除最多 k 个字符，以使 s 的行程长度编码长度最小。
// //
// // 请你返回删除最多 k 个字符后，s 行程长度编码的最小长度 。
// //
// //
// //
// // 示例 1：
// //
// // 输入：s = "aaabcccd", k = 2
// // 输出：4
// // 解释：在不删除任何内容的情况下，压缩后的字符串是 "a3bc3d" ，长度为 6 。最优的方案是删除 'b' 和 'd'，这样一来，压缩后的字符串为 "a3c3" ，长度是 4 。
// // 示例 2：
// //
// // 输入：s = "aabbaa", k = 2
// // 输出：2
// // 解释：如果删去两个 'b' 字符，那么压缩后的字符串是长度为 2 的 "a4" 。
// // 示例 3：
// //
// // 输入：s = "aaaaaaaaaaa", k = 0
// // 输出：3
// // 解释：由于 k 等于 0 ，不能删去任何字符。压缩后的字符串是 "a11" ，长度为 3 。
// //
// //
// // 提示：
// //
// // 1 <= s.length <= 100
// // 0 <= k <= s.length
// // s 仅包含小写英文字母

type Block struct {
	ch byte
	count int
}

type CompressedString []Block

func (c CompressedString) String() string {
	var res string
	for _, block:= range []Block(c) {
		res += fmt.Sprintf("%c", block.ch)
		if block.count > 1 {
			res += fmt.Sprintf("%d", block.count)
		}
	}
	return res
}

func (c CompressedString) CopyBlocks() []Block {
	var blocks []Block
	for _, block:= range []Block(c) {
		blocks = append(blocks, Block{
			ch:    block.ch,
			count: block.count,
		})
	}
	return blocks
}

func (c CompressedString) Length() int {
	length := 0
	for _, block := range []Block(c) {
		length += 1
		if block.count > 1 && block.count < 10 {
			length += 1
		} else if block.count >= 10 {
			val := block.count
			for val != 0 {
				length += 1
				val /= 10
			}
		}
	}
	return length
}

func (c CompressedString) DeleteBlock(index int) []Block{
	var newBlocks []Block
	if index == 0{
		newBlocks = append(newBlocks, []Block(c)[1:]...)
		return newBlocks
	}
	newBlocks = append(newBlocks, []Block(c)[0:index]...)
	if len([]Block(c)) > index + 1 {
		if []Block(c)[index + 1].ch == newBlocks[index-1].ch {
			newBlocks[index-1].count += []Block(c)[index + 1].count
			if len([]Block(c)) > index + 2 {
				newBlocks = append(newBlocks, []Block(c)[index+2:]...)
			}
		} else {
			newBlocks = append(newBlocks, []Block(c)[index+1:]...)
		}
	}
	return newBlocks
}


func convertToBlocks(s string )[]Block {
	var blocks []Block
	currentChar:=s[0]
	count := 0
	for _, ch := range []byte(s) {
		if ch == currentChar {
			count ++
		} else {
			block := Block{
				ch:    currentChar,
				count: count,
			}
			currentChar = ch
			count =  1
			blocks = append(blocks, block)
		}
	}
	block := Block{
		ch:    currentChar,
		count: count,
	}
	blocks = append(blocks, block)
	return blocks
}



func findMinLength(blocks []Block, k int) int {
	// fmt.Printf("Call find min length: %s, %d\n", CompressedString(blocks), k)
	if k == 0 {
		// fmt.Printf("Best for: %s is %d\n", CompressedString(blocks), CompressedString(blocks).Length())
		return CompressedString(blocks).Length()
	}
	minLength := 1000
	for i := range blocks {
		replicaBlocks := (CompressedString(blocks)).CopyBlocks()
		if replicaBlocks[i].count > k {
			replicaBlocks[i].count -= k
			if CompressedString(replicaBlocks).Length() < minLength {
				// fmt.Printf("Better chocie: %s, %d\n", CompressedString(replicaBlocks), CompressedString(replicaBlocks).Length())
				minLength = CompressedString(replicaBlocks).Length()
			}

		} else {
			newK := k - replicaBlocks[i].count
			newBlocks := CompressedString(replicaBlocks).DeleteBlock(i)
			length := findMinLength(newBlocks, newK)
			if length < minLength {
				minLength = length
			}
		}
	}
	// fmt.Printf("Best for: %s, %d is %d\n", CompressedString(blocks), k, minLength)
	return minLength
}

func getLengthOfOptimalCompression(s string, k int) int {
	blocks := convertToBlocks(s)
	// fmt.Println(CompressedString(blocks))
	return findMinLength(blocks, k)
}

func main() {
	fmt.Println(getLengthOfOptimalCompression("abcdefghijklmnopqrstuvwxyzawuidh", 16))
}