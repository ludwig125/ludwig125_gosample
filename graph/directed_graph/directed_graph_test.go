package main

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
)

// 有向グラフの例
// ref: https://flaviocopes.com/golang-data-structure-graph/

func TestGraph(t *testing.T) {
	// tests := map[string]struct {
	// 	str1 string
	// 	str2 string
	// 	want bool
	// }{
	// 	"1": {
	// 		str1: "waterbottle",
	// 		str2: "erbottlewat",
	// 		want: true,
	// 	},
	// 	// "2": {
	// 	// 	str1: "Alfa Bravo Charlie Delta Echo Foxtrot Golf",
	// 	// 	str2: "lta Echo Foxtrot GolfAlfa Bravo Charlie De",
	// 	// 	want: true,
	// 	// },
	// 	// "3": {
	// 	// 	str1: "Alfa Bravo Charlie Delta Echo Foxtrot Golf",
	// 	// 	str2: " lta Echo Foxtrot GolfAlfa Bravo Charlie De",
	// 	// 	want: false,
	// 	// },
	// }
	// for name, tt := range tests {
	// 	t.Run(name, func(t *testing.T) {
	// 		l := singleLinkedList{}
	// 		fmt.Println(l)

	// 	})
	// }

	nl := []*Node{
		{value: Item("0")},
		{value: Item("1")},
		{value: Item("2")},
		{value: Item("3")},
		{value: Item("4")},
		{value: Item("5")},
	}

	g := New()
	// Nodeの登録
	for _, n := range nl {
		g.AddNode(n)
	}

	g.AddEdge(nl[0], nl[1])
	g.AddEdge(nl[0], nl[4])
	g.AddEdge(nl[0], nl[5])
	g.AddEdge(nl[1], nl[3])
	g.AddEdge(nl[1], nl[4])
	g.AddEdge(nl[2], nl[1])
	g.AddEdge(nl[3], nl[2])
	g.AddEdge(nl[3], nl[4])
	g.String()

	seen := make(map[*Node]bool)
	dfs(g, nl[0], seen, 0)
	bfs(g, nl[0])
}

type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	mu    sync.RWMutex
}

type Node struct {
	value Item
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type Item string

// New returns new ItemGraph.
func New() *ItemGraph {
	return &ItemGraph{}
}

// AddNode inserts a node
func (g *ItemGraph) AddNode(n *Node) {
	g.mu.Lock()
	g.nodes = append(g.nodes, n)
	g.mu.Unlock()
}

// AddEdge inserts a edge
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	g.mu.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	// g.edges[*n2] = append(g.edges[*n2], n1) // これがあると無向グラフになる
	g.mu.Unlock()
}

// String for inspection purposes
func (g *ItemGraph) String() {
	g.mu.RLock()
	for _, n := range g.nodes {
		// fmt.Println(n.String())
		fmt.Printf("%v -> %v\n", n.String(), g.edges[*n])
	}
	g.mu.RUnlock()
}

func dfs(g *ItemGraph, s *Node, seen map[*Node]bool, depth int) {
	// time.Sleep(400 * time.Millisecond)
	if _, ok := seen[s]; ok {
		return
	}
	seen[s] = true

	fmt.Printf("%sNode %v\n", fmtDepth(depth), s.String())
	for _, n := range g.edges[*s] {
		dfs(g, n, seen, depth+1)
	}
}

// 幅優先探索
// 参考 https://cybernetist.com/2019/03/09/breadth-first-search-using-go-standard-library/
func bfs(g *ItemGraph, s *Node) {
	g.mu.Lock()
	defer g.mu.Unlock()
	seen := make(map[*Node]bool)

	todo := list.New()
	todo.PushBack(s)

	for todo.Len() > 0 {
		n := todo.Front().Value.(*Node)
		fmt.Println("Node:", n)

		seen[n] = true
		for _, node := range g.edges[*n] {
			if _, ok := seen[node]; !ok {
				// fmt.Println("adjacent", node)
				seen[node] = true
				todo.PushBack(node)
			}
		}
		todo.Remove(todo.Front())
	}
	// fmt.Println("seen", seen)
}

// printしたときに見やすいようにDepthの数だけインデントする
func fmtDepth(num int) string {
	s := ""
	for i := 0; i < num; i++ {
		s += " "
	}
	return s
}
