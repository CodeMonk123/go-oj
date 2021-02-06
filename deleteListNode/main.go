package main

import "fmt"

//给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
//返回删除后的链表的头节点。
//
//注意：此题对比原题有改动
//
//示例 1:
//
//输入: head = [4,5,1,9], val = 5
//输出: [4,1,9]
//解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
//示例 2:
//
//输入: head = [4,5,1,9], val = 1
//输出: [4,5,9]
//解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	} else {
		return fmt.Sprintf("%v -> %s ", l.Val, l.Next.String())
	}
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}
	p, q := head, head.Next
	for q != nil {
		if q.Val == val {
			p.Next = q.Next
			return head
		} else {
			p = q
			q = q.Next
		}
	}

	return head
}

func main() {
	head := ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	fmt.Println(deleteNode(&head, 1))
}
