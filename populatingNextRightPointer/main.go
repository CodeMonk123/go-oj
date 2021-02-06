package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 填充二叉树的每个节点的next指针
// 让next指向右侧的兄弟节点

func connect(root *Node) *Node {
	var nextLayerNode, currentNode *Node
	dummyRoot := &Node{
		Val:   0,
		Left:  root,
		Right: nil,
		Next:  nil,
	}

	currentNode = dummyRoot

	for {
		var tail *Node
		for currentNode != nil {
			if currentNode.Left != nil {
				if tail != nil {
					tail.Next = currentNode.Left
				} else {
					nextLayerNode = currentNode.Left
				}
				if currentNode.Right != nil {
					currentNode.Left.Next = currentNode.Right
					tail = currentNode.Right
				} else {
					tail = currentNode.Left
				}
			} else if currentNode.Right != nil {
				if tail != nil {
					tail.Next = currentNode.Right
				} else {
					nextLayerNode = currentNode.Right
				}
				tail = currentNode.Right
			}
			currentNode = currentNode.Next
		}

		if nextLayerNode == nil {
			break
		} else {
			currentNode = nextLayerNode
			nextLayerNode = nil
		}
	}

	return root
}

func main() {

}
