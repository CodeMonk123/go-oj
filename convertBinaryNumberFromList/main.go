package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var result int
	for p := head; p != nil; p = p.Next {
		result <<= 1
		result = result | p.Val
	}

	return result
}

func main() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	fmt.Println(getDecimalValue(&head))
}
