package btree

import (
	"errors"
	"fmt"
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

func Min(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func Max(node *Node) *Node {
	for node.right != nil {
		node = node.right
	}
	if node.left != nil {
		return node.left
	} else {
		return node
	}
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
