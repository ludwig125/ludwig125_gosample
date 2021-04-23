package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {

	var tr Tree
	nodes := []int{4, 3, 8, 2, 6, 1, 5, 12, 10, 7, 9, 11, 14, 13, 15}
	for _, node := range nodes {
		if err := tr.Insert(node, fmt.Sprintf("%d", node)); err != nil {
			t.Fatal(err)
		}
	}

	// n := Node{Key: 4, Value: "4"}
	// // n.Print()
	// // fmt.Println("n", n)
	// if err := n.Insert(3, "3"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(8, "8"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(2, "2"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(6, "6"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(1, "1"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(5, "5"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(12, "12"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(10, "10"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(7, "7"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(9, "9"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(11, "11"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(14, "14"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(13, "13"); err != nil {
	// 	t.Fatal(err)
	// }
	// if err := n.Insert(15, "15"); err != nil {
	// 	t.Fatal(err)
	// }

	// fmt.Println(n)
	// fmt.Println(n.Left)
	// fmt.Println(n.Right)
	n.Print()

	fmt.Println("-----------------")
	// fmt.Println(n.Find("5"))
	// fmt.Println(n.Find("8"))

	if err := n.Delete(4, nil); err != nil {
		t.Fatal(err)
	}
	n.Print()
}

type Node struct {
	Key   int
	Value string
	Left  *Node
	Right *Node
}

func (n *Node) Insert(key int, value string) error {
	if n == nil {
		return errors.New("cannot insert nil node")
	}

	if key == 12 {
		fmt.Println("koko1")
	}
	switch {
	case n.Key == key:
		if key == 12 {
			fmt.Println("koko1")
		}
		return errors.New("failed to insert. already be same key")
	case n.Key > key:
		if key == 12 {
			fmt.Println("koko2", n.Key, key, n.Key > key)
		}
		if n.Left == nil {
			n.Left = &Node{Key: key, Value: value}
			return nil
		}
		return n.Left.Insert(key, value)
	case n.Key < key:
		if key == 12 {
			fmt.Println("koko3")
		}
		if n.Right == nil {
			n.Right = &Node{Key: key, Value: value}
			return nil
		}
		return n.Right.Insert(key, value)
	}
	return nil
}

func (n *Node) Find(key int) (string, bool) {
	if n == nil {
		return "", false
	}
	// fmt.Println("find:", n.Key)
	switch {
	case n.Key == key:
		return n.Value, true
	case n.Key > key:
		if n.Left != nil {
			return n.Left.Find(key)
		}
		return "", false
	case n.Key < key:
		if n.Right != nil {
			return n.Right.Find(key)
		}
		return "", false
	}
	return "", false
}

func (n *Node) Delete(key int, parent *Node) error {
	if n == nil {
		return errors.New("cannnot delete key this node is nil")
	}
	fmt.Println("koko1", n.Key)
	if n.Key > key {
		fmt.Println("koko1-2", n.Key)
		if n.Left != nil {
			fmt.Println("koko1-3", n.Key)
			return n.Left.Delete(key, n)
		}
		// fmt.Println("koko1-4", n.Key)
		return errors.New("cannnot delete key left node is nil")
	}

	// fmt.Println("koko2", n.Key)
	if n.Key < key {
		// fmt.Println("koko2-2", n.Key)
		if n.Right != nil {
			// fmt.Println("koko2-3", n.Key)
			return n.Right.Delete(key, n)
		}
		return errors.New("cannnot delete key right node is nil")
	}

	// fmt.Println("koko3", n.Key)
	// n.Key == key
	// if n.Key == key {
	// fmt.Println("koko3-1", n.Key)
	if n.Left == nil && n.Right == nil {
		// fmt.Println("koko3-2", n.Key)
		// parent.Print()
		n.replaceNode(parent, nil)
		// fmt.Println("-----------------")
		// parent.Print()
		// fmt.Println("-----------------")
		return nil
	}

	// half leaf node
	if n.Left != nil && n.Right == nil {
		// fmt.Println("koko3-3", n.Key)
		n.replaceNode(parent, n.Left)
		return nil
	}

	if n.Left == nil && n.Right != nil {
		// fmt.Println("koko3-4", n.Key)
		n.replaceNode(parent, n.Right)
		return nil
	}

	// n.Left != nil && n.Right != nil
	// fmt.Println("koko3-5", n.Key)
	newNode, parent := n.Left.findMax(n)

	if err := n.Delete(newNode.Key, parent); err != nil {
		return err
	}
	// newNode.Left = n.Left
	// newNode.Right = n.Right
	// fmt.Println("findmax", newNode)
	// newNode.Print()
	n.Key = newNode.Key
	n.Value = newNode.Value
	fmt.Println("----------------------")
	// n.replaceNode(parent, newNode)
	return nil
}

func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("failed")
	}
	// fmt.Println("parent, replacement", parent, replacement)
	if n == parent.Left {
		parent.Left = replacement
	}
	if n == parent.Right {
		parent.Right = replacement
	}
	// parent = replacement
	// fmt.Println("parent, replacement", parent, replacement)
	return nil
}

// Node n以下で最大の数値を見つける
// あとでその最大の数値をその親Nodeから消す必要があるのでparentを返す
func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(parent)
}

// func (n *Node) Delete(key string) (*Node, error) {
// 	if n == nil {
// 		return nil, errors.New("cannnot delete key this node is nil")
// 	}
// 	fmt.Println("koko1", n.Key)
// 	if n.Key > key {
// 		fmt.Println("koko1-2", n.Key)
// 		if n.Left != nil {
// 			fmt.Println("koko1-3", n.Key)
// 			// return n.Left.Delete(key)
// 			newN, err := n.Left.Delete(key)
// 			fmt.Println("koko1-4 newN err", newN, err)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to delete: %v", err)
// 			}
// 			n.Print()
// 			n.Left = newN
// 			fmt.Println("-----------------")
// 			n.Print()
// 			return n, nil
// 		}
// 		fmt.Println("koko1-4", n.Key)
// 		return nil, errors.New("cannnot delete key left node is nil")
// 	}

// 	fmt.Println("koko2", n.Key)
// 	if n.Key < key {
// 		fmt.Println("koko2-2", n.Key)
// 		if n.Right != nil {
// 			fmt.Println("koko2-3", n.Key)
// 			newN, err := n.Right.Delete(key)
// 			fmt.Println("koko2-4 newN err", newN, err)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to delete: %v", err)
// 			}
// 			n.Print()
// 			n.Right = newN
// 			fmt.Println("-----------------")
// 			n.Print()
// 			return n, nil
// 		}
// 		return nil, errors.New("cannnot delete key right node is nil")
// 	}

// 	fmt.Println("koko3", n.Key)
// 	// n.Key == key
// 	// if n.Key == key {
// 	fmt.Println("koko3-1", n.Key)
// 	if n.Left == nil && n.Right == nil {
// 		fmt.Println("koko3-2", n.Key)
// 		return nil, nil
// 	}

// 	// half leaf node
// 	if n.Left != nil && n.Right == nil {
// 		fmt.Println("koko3-3", n.Key)
// 		return nil, nil
// 	}

// 	if n.Left == nil && n.Right != nil {
// 		fmt.Println("koko3-4", n.Key)
// 		return nil, nil
// 	}

// 	// n.Left != nil && n.Right != nil
// 	fmt.Println("koko3-5", n.Key)

// 	// }

// 	return nil, nil
// }

func (n *Node) Print() {
	n.print(0)
}

func (n *Node) print(depth int) {
	if n == nil {
		// return errors.New("cannot insert nil node")
		return
	}
	n.Left.print(depth + 1)
	fmt.Println(formatDepth(depth), n.Key)
	n.Right.print(depth + 1)
	return
}

func formatDepth(depth int) string {
	s := ""
	for i := 0; i < depth; i++ {
		s += "        "
	}
	return s
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(key int, value string) error {
	if t.Root == nil {
		t.Root = &Node{
			Key:   key,
			Value: value,
		}
	}
	return t.Root.Insert(key, value)
}

func (t *Tree) Find(key int) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(key)
}

func (t *Tree) Delete(key int) error {
	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}
	fakeParent := &Node{Right: t.Root}
	if err := t.Root.Delete(key, fakeParent); err != nil {
		return err
	}
	if fakeParent.Right == nil {
		t.Root = nil
	}
	return nil
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}
