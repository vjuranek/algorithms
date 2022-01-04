package btree

import (
	"fmt"
	"testing"
)

func CreateTree() *Node {
	root := NewNode(5)
	n1 := NewNode(3)
	n2 := NewNode(7)
	root.left = n1
	n1.parent = root
	root.right = n2
	n2.parent = root

	n1.left = NewNode(1)
	n1.left.parent = n1
	n1.right = NewNode(4)
	n1.right.parent = n1
	n2.right = NewNode(8)
	n2.right.parent = n2

	return root
}

func TestMin(t *testing.T) {
	root := CreateTree()
	min := Min(root)

	if min.value != 1 {
		t.Fatalf("Wrong min value, expected=1, actual=%d", min.value)
	}
}

func TestMax(t *testing.T) {
	root := CreateTree()
	max := Max(root)

	if max.value != 8 {
		t.Fatalf("Wrong max value, expected=8, actual=%d", max.value)
	}
}

func TestInsertEmptyTree(t *testing.T) {
	tree := &BTree{nil}
	n := NewNode(10)
	Insert(tree, n)

	if tree.root.value != 10 {
		t.Fatalf("Wrong root node value %d", tree.root.value)
	}

	if tree.root.left != nil || tree.root.right != nil {
		t.Fatalf("Root child is not nil")
	}
}

func TestInsertLeft(t *testing.T) {
	root := CreateTree()
	tree := &BTree{root}
	node := NewNode(0)
	Insert(tree, node)

	if tree.root.left.left.left != node {
		PrintPostOrder(tree.root)
		fmt.Print("\n")
		t.Fatalf("Wrong position of newly inserted node:")
	}
}

func TestInsertRight(t *testing.T) {
	root := CreateTree()
	tree := &BTree{root}
	node := NewNode(2)
	Insert(tree, node)

	if tree.root.left.left.right != node {
		PrintPostOrder(tree.root)
		fmt.Print("\n")
		t.Fatalf("Wrong position of newly inserted node:")
	}
}

func TestDeleteLeft(t *testing.T) {
	root := CreateTree()
	tree := &BTree{root}
	Delete(tree, tree.root.left)

	if tree.root.left.value != 4 {
		PrintPostOrder(tree.root)
		fmt.Print("\n")
		t.Fatalf("Wrong delete of node")
	}

	if tree.root.left.left.value != 1 {
		t.Fatalf("Wrong left child")
	}

	if tree.root.left.right != nil {
		t.Fatalf("Wrong right child")
	}
}

func TestDeleteRight(t *testing.T) {
	root := CreateTree()
	tree := &BTree{root}
	Delete(tree, tree.root.right)

	if tree.root.right.value != 8 {
		PrintPostOrder(tree.root)
		fmt.Print("\n")
		t.Fatalf("Wrong delete of node")
	}

	if tree.root.right.left != nil {
		t.Fatalf("Wrong left child")
	}

	if tree.root.right.right != nil {
		t.Fatalf("Wrong right child")
	}

}

func TestDeleteLeaf(t *testing.T) {
	root := CreateTree()
	tree := &BTree{root}
	Delete(tree, tree.root.left.right)

	if tree.root.left.value != 3 {
		PrintPostOrder(tree.root)
		fmt.Print("\n")
		t.Fatalf("Wrong parent value after delete")
	}

	if tree.root.left.right != nil {
		t.Fatalf("Node not deleted")
	}
}
