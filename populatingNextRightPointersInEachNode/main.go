package main

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	currentLayer := root
	for {
		// fmt.Printf("Current layer start: %d\n", currentLayer.Val)
		var nextLayer *Node
		var parent *Node
		nextLayer = nil
		parent = nil
		for p := currentLayer; p != nil; p = p.Next {
			if p.Left != nil {
				parent = p
				nextLayer = p.Left
				break
			} else if p.Right != nil {
				parent = p
				nextLayer = p.Right
				break
			}
		}
		// current 全是叶子结点
		if nextLayer == nil {
			return root
		}
		// fmt.Printf("Next layer start: %d\n", nextLayer.Val)

		tail := nextLayer
		if parent.Left != nil {
			if parent.Right != nil {
				tail.Next = parent.Right
				// fmt.Printf("Link %d -> %d\n", tail.Val, tail.Next.Val)
				tail = parent.Right
			}
		}

		parent = parent.Next
		for parent != nil {
			if parent.Left != nil {
				tail.Next = parent.Left
				// fmt.Printf("Link %d -> %d\n", tail.Val, tail.Next.Val)
				tail = parent.Left
				if parent.Right != nil {
					tail.Next = parent.Right
					// fmt.Printf("Link %d -> %d\n", tail.Val, tail.Next.Val)
					tail = parent.Right
				}
			} else if parent.Right != nil {
				tail.Next = parent.Right
				// fmt.Printf("Link %d -> %d\n", tail.Val, tail.Next.Val)
				tail = parent.Right

			}
			parent = parent.Next
		}

		currentLayer = nextLayer
	}
}

func main() {
	root := &Node{
		Val:   7,
		Left:  &Node{
			Val:   -10,
			Left:  &Node{
				Val:   -4,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   3,
				Left:  nil,
				Right: &Node{
					Val:   -1,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Next:  nil,
			},
			Next:  nil,
		},
		Right: &Node{
			Val:   2,
			Left:  &Node{
				Val:   8,
				Left:  &Node{
					Val:   1,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Right: &Node{
					Val:   1,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Next:  nil,
			},
			Right: nil,
			Next:  nil,
		},
		Next:  nil,
	}

	connect(root)
}
