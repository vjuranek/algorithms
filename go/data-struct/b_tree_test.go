package btree

import (
	"testing"
)

func CreateTree() (*Node) {
	root := NewNode(5)
	n1 := NewNode(4)
	n2 := NewNode(7)
	root.left = n1
	root.right = n2

	n1.left = NewNode(1)
	n1.right = NewNode(3)
	n2.left = NewNode(8)

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
