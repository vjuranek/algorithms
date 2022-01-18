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

func Delete(t *BTree, node *Node) {
	if node.left == nil {
		switchSubtree(t, node, node.right)
	} else if node.right == nil {
		switchSubtree(t, node, node.left)
	} else {
		minRight := Min(node.right)
		if minRight.parent != node {
			// efectively remove minRight from the tree, as
			// it hasn't anything on the left and now is replaced
			// with its right child
			switchSubtree(t, minRight, minRight.right)
			
			// handle right child of node to be deleted
			// minRight cannot have any left child, so no need
			// to set left child
			minRight.right = node.right
			// fix parent after switching right child
			minRight.right.parent = minRight
		}
		switchSubtree(t, node, minRight)
		// we have to fix only left link, by now right child of node to
		// be delete is minRight or the right link was already fixed in
		// branch above
		minRight.left = node.left
		minRight.left.parent = minRight
	}
}

// switchSubtree switches sub-tree with root node1 with sub-tree with root node2.
// Setting up child poiters is responsibility of the called (if it's needed).
func switchSubtree(t *BTree, node1 *Node, node2 *Node) {
	
	// replace node1 by node2 in the tree
	if node1.parent == nil {
		t.root = node2
	} else if node1 == node1.parent.left {
		node1.parent.left = node2
	} else {
		node1.parent.right = node2
	}

	// fix node2 parent
	if node2 != nil {
		node2.parent = node1.parent
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
