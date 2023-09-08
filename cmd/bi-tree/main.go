package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func newNode(key int) *Node {
	return &Node{Key: key}
}

func (n *Node) Insert(key int) {
	if n.Key == key {
		fmt.Printf("Key %d already exists, please provide a key that is not already implemented.\n", key)
		return
	}

	if key < n.Key {
		//move to left
		if n.Left == nil {
			n.Left = newNode(key)
		} else {
			n.Left.Insert(key)
		}
	} else if key > n.Key {
		// move to the right
		if n.Right == nil {
			n.Right = newNode(key)
		} else {
			n.Right.Insert(key)
		}
	}
}

func (n *Node) Search(key int) bool {
	if n.Key == key {
		return true
	} else if key < n.Key {
		n.Left.Search(key)
	} else {
		n.Right.Search(key)
	}
	return false
}

func (n *Node) TraverseTree() {
	if n != nil {
		fmt.Println(n.Key)
		n.Left.TraverseTree()
		n.Right.TraverseTree()
	}
	return
}

func main() {
	tree := newNode(10)
	tree.Insert(20)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(30)
	tree.Insert(6)

	tree.TraverseTree()

}
