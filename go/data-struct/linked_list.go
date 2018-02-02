package main

import (
	"fmt"
	"errors"
)

type Node struct { //TODO some packaging
	value int //TODO some generalization, template, generic type?
	next  *Node
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

func (l *List) add(val int) {
	node := Node{value: val}
	if l.head == nil {
		l.head = &node
		l.tail = &node
	} else {
		l.tail.next = &node
		l.tail = &node
	}
	l.length++
}

func (l *List) get(i int) (int, error) {
	if i >= l.length {
		return 0, errors.New("To big index value")
	}
	node := l.head
	for idx := 0; idx < i; idx++ {
		node = node.next
	}
	return node.value, nil
}

func (l *List) remove(i int) error {
	if i >= l.length {
		return errors.New("To big index value")
	}

	node := l.head
	if i == 0 {
		l.head = node.next
	} else {
		for idx := 0; idx < i-1; idx++ {
			node = node.next
		}
	}
	if node.next == l.tail {
		l.tail = node
		node.next = nil
	} else {
		node.next = node.next.next
	}
	l.length--
	return nil
}

func main() {
	l := List{}
	l.add(10)
	l.add(20)
	l.add(30)
	l.add(40)
	l.add(50)
	l.remove(3)
	val, _ :=  l.get(2)
	fmt.Println("3 element value is ", val)
	val, _ =  l.get(3)
	fmt.Println("3 element value is ", val)
}
