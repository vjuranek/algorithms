package main

import (
	"fmt"
	"errors"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type BTree struct {
	root *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil}
} 

func PrintPostOrder(tree *Node) (err error) {
	if tree == nil {
		return errors.New("Node is nil")
	}
	
	if tree.left != nil {
		PrintPostOrder(tree.left)
	}
	

	if tree.right != nil {
		PrintPostOrder(tree.right)
	}

	fmt.Print(tree.value)

	return
}

func main() {
	root := NewNode(5)
	n1 := NewNode(4)
	n2 := NewNode(7)
	root.left = n1
	root.right = n2

	n1.left = NewNode(1)
	n1.right = NewNode(3)
	n2.left = NewNode(8)

	PrintPostOrder(root)
}
