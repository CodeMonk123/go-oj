// 在有向图中, 我们从某个节点和每个转向处开始, 沿着图的有向边走。 如果我们到达的节点是终点 (即它没有连出的有向边), 我们停止。

// 现在, 如果我们最后能走到终点，那么我们的起始节点是最终安全的。 更具体地说, 存在一个自然数 K,  无论选择从哪里开始行走, 我们走了不到 K 步后必能停止在一个终点。

// 哪些节点最终是安全的？ 结果返回一个有序的数组。

// 该有向图有 N 个节点，标签为 0, 1, ..., N-1, 其中 N 是 graph 的节点数.  图以以下的形式给出: graph[i] 是节点 j 的一个列表，满足 (i, j) 是图的一条有向边。

// 示例：
// 输入：graph = [[1,2],[2,3],[5],[0],[5],[],[]]
// 输出：[2,4,5,6]

package main

import "fmt"

type Graph [][]int

const (
	WHITE = iota
	GREY
	BLACK
)

//	如果返回false，说明current是不安全的
func dfs(graph Graph, current int, color *map[int]int, isSafe *map[int]bool) bool {
	fmt.Println("Visited: ", current)
	neighbours := graph[current]
	(*color)[current] = GREY
	safe := true
	for _, v := range neighbours {
		if (*color)[v] == WHITE {
			if !dfs(graph, v, color, isSafe) {
				fmt.Printf("%v point to an unsafe Node: %v\n", current, v)
				safe = false
				(*isSafe)[current] = false
			}
		} else if (*color)[v] == GREY {
			fmt.Printf("Find a BE: %v->%v\n", current, v)
			safe = false
			(*isSafe)[current] = false
		} else {
			if !(*isSafe)[v] {
				safe = false
				(*isSafe)[current] = false
			}
		}
	}
	(*color)[current] = BLACK
	return safe
}

func eventualSafeNodes(graph [][]int) []int {
	isSafe := make(map[int]bool)
	color := make(map[int]int)
	for i := range graph {
		isSafe[i] = true
		color[i] = WHITE
	}

	for i := range graph {
		if color[i] == WHITE {
			dfs(Graph(graph), i, &color, &isSafe)
		}
	}

	res := []int{}

	for i := range graph {
		if isSafe[i] {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	fmt.Println(eventualSafeNodes([][]int{
		{1, 2},
		{2, 3},
		{5},
		{0},
		{5},
		{},
		{},
	}))
}
