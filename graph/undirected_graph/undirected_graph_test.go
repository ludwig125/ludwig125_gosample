package main

import (
	"fmt"
	"sync"
	"testing"
)

// 無向グラフの例
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
	}

	g := New()
	for _, n := range nl {
		g.AddNode(n)
	}

	g.AddEdge(nl[0], nl[1])
	g.AddEdge(nl[0], nl[2])
	g.AddEdge(nl[1], nl[3])
	g.AddEdge(nl[1], nl[4])
	g.String()

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
	g.edges[*n2] = append(g.edges[*n2], n1)
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

// func (g *ItemGraph) String() {
// 	g.mu.RLock()
// 	s := ""
// 	fmt.Println(g.nodes)
// 	for i := 0; i < len(g.nodes); i++ {
// 		s += g.nodes[i].String() + " -> "
// 		near := g.edges[*g.nodes[i]]
// 		for j := 0; j < len(near); j++ {
// 			s += near[j].String() + " "
// 		}
// 		s += "\n"
// 	}
// 	fmt.Println(s)
// 	g.mu.RUnlock()
// }
