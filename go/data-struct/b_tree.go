package btree

import (
	"errors"
	"fmt"
)

type Node struct {
	value  int
	parent *Node
	left   *Node
	right  *Node
}

type BTree struct {
	root *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil, nil}
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
	return node
}

func Insert(t *BTree, node *Node) {
	x := t.root
	var y *Node = nil

	for x != nil {
		y = x
		if node.value < x.value {
			x = x.left
		} else {
			x = x.right
		}
	}

	node.parent = y
	if y == nil {
		t.root = node
	} else if node.value < y.value {
		y.left = node
	} else {
		y.right = node
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
