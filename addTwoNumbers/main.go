package main

import "fmt"

// 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
// 示例：
//
// 输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 8 -> 0 -> 7

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Length() int {
	if l == nil {
		return 0
	}
	return 1 + l.Next.Length()
}

func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	}
	return fmt.Sprintf("%d->%s", l.Val, l.Next)
}

func align(l *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		newNode := &ListNode{
			Val:  0,
			Next: l,
		}
		l = newNode
	}
	return l
}

func handle(l1, l2 *ListNode) (*ListNode, int) {
	if l1 == nil {
		return nil, 0
	}
	rest, overflow := handle(l1.Next, l2.Next)
	val := l1.Val + l2.Val + overflow
	newNode := &ListNode{
		Val:  val % 10,
		Next: rest,
	}
	return newNode, val / 10
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	length1, length2 := l1.Length(), l2.Length()
	if length1 > length2 {
		l2 = align(l2, length1-length2)
	} else {
		l1 = align(l1, length2-length1)
	}
	res, overflow := handle(l1, l2)
	if overflow != 0 {
		newNode := &ListNode{
			Val:  overflow,
			Next: res,
		}
		return newNode
	}
	return res
}

func main() {
	l1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Println(addTwoNumbers(l1, l2))
}
