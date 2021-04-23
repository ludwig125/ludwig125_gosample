package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 参考: https://flaviocopes.com/golang-data-structure-binary-search-tree/

func TestBinarySearchTree(t *testing.T) {
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

	bt := New()
	// bt.Insert(8, Item("8"))
	// bt.Insert(4, Item("4"))
	// bt.Insert(10, Item("10"))
	// bt.Insert(2, Item("2"))
	// bt.Insert(6, Item("6"))
	// bt.Insert(1, Item("1"))
	// bt.Insert(3, Item("3"))
	// bt.Insert(5, Item("5"))
	// bt.Insert(7, Item("7"))
	// bt.Insert(9, Item("9"))

	bt.Insert(4, Item("4"))
	bt.Insert(3, Item("3"))
	bt.Insert(8, Item("8"))
	bt.Insert(2, Item("2"))
	bt.Insert(6, Item("6"))
	bt.Insert(1, Item("1"))
	bt.Insert(5, Item("5"))
	bt.Insert(12, Item("12"))
	bt.Insert(10, Item("10"))
	bt.Insert(7, Item("7"))
	bt.Insert(9, Item("9"))
	bt.Insert(11, Item("11"))
	bt.Insert(14, Item("14"))
	bt.Insert(13, Item("13"))
	bt.Insert(15, Item("15"))
	bt.String()

	min := bt.Min()
	fmt.Println("min", *min)

	max := bt.Max()
	fmt.Println("max", *max)

	fmt.Println("search", bt.Search(3))

	bt.Remove(8)
	bt.String()

	bt.Remove(9)
	bt.String()
}

type Item string

// Node a single node that composes the tree
type Node struct {
	key   int
	value Item
	left  *Node //left
	right *Node //right
}

// ItemBinarySearchTree the binary search tree of Items
type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func New() *ItemBinarySearchTree {
	return &ItemBinarySearchTree{}
}

func (t *ItemBinarySearchTree) String() {
	// t.lock.Lock()
	// defer t.lock.Unlock()
	printTree(t.root, 0)

	kd := make([]keyDepth, 0)
	makeTreePicture(&kd, t.root, 0)
	printTreeTranspose(kd)
}

func debug(n *Node, str string) {
	time.Sleep(100 * time.Millisecond)
	if n == nil {
		return
	}

	debug(n.left, str+"_left")
	fmt.Println(n.key, str)
	debug(n.right, str+"_right")
}

func printTree(n *Node, depth int) {
	if n == nil {
		return
	}
	format := ""
	for i := 0; i < depth; i++ {
		format += "        "
	}
	format += "---[ "
	depth++
	printTree(n.left, depth)
	fmt.Println(format, n.key)
	printTree(n.right, depth)
}

type keyDepth struct {
	key   int
	depth int
}

func makeTreePicture(kd *[]keyDepth, n *Node, depth int) {
	if n == nil {
		return
	}
	depth++
	makeTreePicture(kd, n.left, depth)
	*kd = append(*kd, keyDepth{key: n.key, depth: depth})
	makeTreePicture(kd, n.right, depth)
}

func printTreeTranspose(kd []keyDepth) {
	maxDepth := getMaxDepth(kd)
	for i := 1; i <= maxDepth; i++ {
		for _, v := range kd {
			if v.depth == i {
				fmt.Printf("%2d", v.key)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Printf("\n\n")
	}
}

func getMaxDepth(kd []keyDepth) int {
	max := 1
	for _, v := range kd {
		if max < v.depth {
			max = v.depth
		}
	}
	return max
}

func (t *ItemBinarySearchTree) Insert(key int, value Item) {
	t.lock.Lock()
	defer t.lock.Unlock()
	n := &Node{key, value, nil, nil}
	if t.root == nil {
		t.root = n
		return
	}
	insertNode(t.root, n)
}

func insertNode(node, newNode *Node) {
	if node.key > newNode.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (t *ItemBinarySearchTree) InOrderTraversal(f func(Item)) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	inOrderTraversal(t.root, f)
}

func inOrderTraversal(n *Node, f func(Item)) {
	if n != nil {
		inOrderTraversal(n.left, f)
		f(n.value)
		inOrderTraversal(n.right, f)
	}
}

func (t *ItemBinarySearchTree) PreOrderTraversal(f func(Item)) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	preOrderTraversal(t.root, f)
}

func preOrderTraversal(n *Node, f func(Item)) {
	if n != nil {
		f(n.value)
		preOrderTraversal(n.left, f)
		preOrderTraversal(n.right, f)
	}
}

func (t *ItemBinarySearchTree) PostOrderTraversal(f func(Item)) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	postOrderTraversal(t.root, f)
}

func postOrderTraversal(n *Node, f func(Item)) {
	if n != nil {
		postOrderTraversal(n.left, f)
		postOrderTraversal(n.right, f)
		f(n.value)
	}
}

func (t *ItemBinarySearchTree) Min() *Item {
	t.lock.RLock()
	defer t.lock.RUnlock()
	n := t.root
	if n == nil {
		return nil
	}
	for {
		if n.left == nil {
			return &n.value
		}
		n = n.left
	}
}

func (t *ItemBinarySearchTree) Max() *Item {
	t.lock.RLock()
	defer t.lock.RUnlock()
	n := t.root
	if n == nil {
		return nil
	}
	for {
		if n.right == nil {
			return &n.value
		}
		n = n.right
	}
}

func (t *ItemBinarySearchTree) Search(key int) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return search(t.root, key)
}

func search(n *Node, key int) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		return search(n.left, key)
	}
	if key > n.key {
		return search(n.right, key)
	}

	// ↓ uun.key == key
	return true
}

func (t *ItemBinarySearchTree) Remove(key int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	remove(t.root, key)
}

func remove(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = remove(n.left, key)
		return n
	}
	if key > n.key {
		n.right = remove(n.right, key)
		return n
	}

	// 以下key ==n.keyの場合

	// nの先に何も子がいない場合
	if n.left == nil && n.right == nil {
		n = nil
		return nil
	}

	// nのrightに子がいない場合
	if n.right == nil {
		n = n.left
		return n
	}

	// nのleftに子がいない場合
	if n.left == nil {
		n = n.right
		return n
	}

	smallestRight := n.right
	for {
		if smallestRight.left == nil {
			break
		}
		smallestRight = smallestRight.left
	}
	fmt.Println("smallestRight", smallestRight)
	// nodeの右側のnodeたちの中で最も小さな数字のnodeを見つけて、そのnodeで自分自身を置き換える
	n.key = smallestRight.key
	n.value = smallestRight.value
	// 置き換えた後の元のNodeは消す
	n.right = remove(n.right, n.key)
	return n
}
