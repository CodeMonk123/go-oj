// 在本问题中, 树指的是一个连通且无环的无向图。

// 输入一个图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N) 的树及一条附加的边构成。附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。

// 结果图是一个以边组成的二维数组。每一个边的元素是一对[u, v] ，满足 u < v，表示连接顶点u 和v的无向图的边。

// 返回一条可以删去的边，使得结果图是一个有着N个节点的树。如果有多个答案，则返回二维数组中最后出现的边。答案边 [u, v] 应满足相同的格式 u < v。

// 示例 1：

// 输入: [[1,2], [1,3], [2,3]]
// 输出: [2,3]
// 解释: 给定的无向图为:
//   1
//  / \
// 2 - 3
// 示例 2：

// 输入: [[1,2], [2,3], [3,4], [1,4], [1,5]]
// 输出: [1,4]
// 解释: 给定的无向图为:
// 5 - 1 - 2
//     |   |
//     4 - 3
// 注意:

// 输入的二维数组大小在 3 到 1000。
// 二维数组中的整数在1到N之间，其中N是输入数组的大小。

package main

import "fmt"

type Graph map[int][]int

func generateGraph(edges *[][]int) Graph {
	graph := make(map[int][]int)
	for _, edge := range *edges {
		n0 := edge[0]
		n1 := edge[1]
		graph[n0] = append(graph[n0], n1)
		graph[n1] = append(graph[n1], n0)
	}
	fmt.Println(graph)
	return graph
}

func findCircle(node int, precursor *map[int]int, visited *map[int]bool, graph *Graph) [][]int {
	(*visited)[node] = true
	neighbours := (*graph)[node]
	circle := [][]int{}
	for _, v := range neighbours {
		if v != (*precursor)[node] {
			if !(*visited)[v] {
				(*precursor)[v] = node
				temp := findCircle(v, precursor, visited, graph)
				if temp != nil {
					circle = temp
					break
				}
			} else {
				fmt.Printf("find a back edge:%d->%d\n", node, v)
				circle = append(circle, []int{node, v})
				p := node
				for p != v {
					circle = append(circle, []int{p, (*precursor)[p]})
					p = (*precursor)[p]
				}
				break
			}
		}
	}
	if len(circle) == 0 {
		return nil
	}
	return circle
}

func findRedundantConnection(edges [][]int) []int {
	graph := generateGraph(&edges)
	visited := make(map[int]bool)
	precursor := make(map[int]int)
	precursor[1] = -1
	circle := findCircle(1, &precursor, &visited, &graph)
	for i := range circle {
		if circle[i][0] > circle[i][1] {
			circle[i] = []int{circle[i][1], circle[i][0]}
		}
	}
	fmt.Println(circle)
	for i := len(edges) - 1; i >= 0; i-- {
		for _, e := range circle {
			if edges[i][0] == e[0] && edges[i][1] == e[1] {
				return e
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 3}, {3, 4}, {1, 5}, {3, 5}, {2, 3}}))
}
