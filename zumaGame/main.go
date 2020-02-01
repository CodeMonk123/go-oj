// Think about Zuma Game. You have a row of balls on the table, colored red(R), yellow(Y), blue(B), green(G), and white(W). You also have several balls in your hand.

// Each time, you may choose a ball in your hand, and insert it into the row (including the leftmost place and rightmost place). Then, if there is a group of 3 or more balls in the same color touching, remove these balls. Keep doing this until no more balls can be removed.

// Find the minimal balls you have to insert to remove all the balls on the table. If you cannot remove all the balls, output -1.

// Example 1:

// Input: board = "WRRBBW", hand = "RB"
// Output: -1
// Explanation: WRRBBW -> WRR[R]BBW -> WBBW -> WBB[B]W -> WW
// Example 2:

// Input: board = "WWRRBBWW", hand = "WRBRW"
// Output: 2
// Explanation: WWRRBBWW -> WWRR[R]BBWW -> WWBBWW -> WWBB[B]WW -> WWWW -> empty
// Example 3:

// Input: board = "G", hand = "GGGGG"
// Output: 2
// Explanation: G -> G[G] -> GG[G] -> empty
// Example 4:

// Input: board = "RBYYBBRRB", hand = "YRBGB"
// Output: 3
// Explanation: RBYYBBRRB -> RBYY[Y]BBRRB -> RBBBRRB -> RRRB -> B -> B[B] -> BB[B] -> empty

// Constraints:

// You may assume that the initial row of balls on the table won’t have any 3 or more consecutive balls with the same color.
// The number of balls on the table won't exceed 16, and the string represents these balls is called "board" in the input.
// The number of balls in your hand won't exceed 5, and the string represents these balls is called "hand" in the input.
// Both input strings will be non-empty and only contain characters 'R','Y','B','G','W'.

package main

import "fmt"

type group struct {
	Color byte
	Num   int
	Next  *group
}

func (head *group) String() string {
	if head == nil {
		return "nil\n"
	} else if head.Color == ' ' {
		return "[HEAD]->" + head.Next.String()
	}
	res := fmt.Sprintf("[%c, %d]->", head.Color, head.Num) + head.Next.String()
	return res

}

// 消除合并
func (head *group) erase() {
	for {
		//fmt.Println(head)
		change := false
		for p, q := head, head.Next; q != nil; {
			if q.Num >= 3 {
				p.Next = q.Next
				q = q.Next
				change = true
				if q == nil {
					break
				}
			}
			p, q = p.Next, q.Next
		}
		if !change {
			break
		} else {
			for p, q := head, head.Next; q != nil; {
				if p.Color == q.Color {
					p.Num += q.Num
					p.Next = q.Next
					q = q.Next
					if q == nil {
						break
					}
				}
				p, q = p.Next, q.Next
			}
		}
	}
}

//	拷贝
func (head *group) copy() *group {
	newHead := group{' ', -1, nil}
	tail := &newHead
	for p := head.Next; p != nil; p = p.Next {
		q := &group{p.Color, p.Num, nil}
		tail.Next = q
		tail = q
	}
	return &newHead
}

func generateGroups(board string) *group {
	balls := []byte(board)
	current := balls[0]
	first := group{current, 1, nil}
	tail := &first
	for i := 1; i < len(balls); i++ {
		if balls[i] == current {
			tail.Num++
		} else {
			newTail := group{balls[i], 1, nil}
			tail.Next = &newTail
			tail = &newTail
			current = balls[i]
		}
	}
	head := group{' ', -1, &first}
	return &head
}

//深度搜索
func dfs(head *group, rest []byte) int {
	if head.Next == nil {
		return 0 //完成
	}
	if len(rest) == 0 {
		return -1 //未完成
	}

	minStep := -1
	for i, ball := range rest {
		for count, g := 0, head.Next; g != nil; count, g = count+1, g.Next {
			if g.Color == ball {
				replica := head.copy()
				p := replica.Next
				for j := 0; j < count; j++ {
					p = p.Next
				}
				p.Num++
				// fmt.Printf("Replica: %v\n", replica)
				replica.erase()
				// fmt.Printf("After erase: %v\n", replica)
				newRest := append([]byte{}, rest[:i]...)
				newRest = append(newRest, rest[i+1:]...)
				temp := dfs(replica, newRest)
				if temp != -1 {
					if minStep == -1 {
						minStep = temp + 1
					} else if minStep > temp+1 {
						minStep = temp + 1
					}
				}
			}
		}
	}
	return minStep
}

func findMinStep(board string, hand string) int {
	head := generateGroups(board)
	//fmt.Println(head)
	return dfs(head, []byte(hand))
}

func main() {
	fmt.Println(findMinStep("RBYYBBRRB", "YRBGB"))
	fmt.Println(findMinStep("WRRBBW", "RB"))
	fmt.Println(findMinStep("WWRRBBWW", "WRBRW"))
	fmt.Println(findMinStep("G", "GGGGG"))
	//输入"RRWWRRBBRR" "WB" 时，leetCode的答案是2，但是个人认为是-1
}
