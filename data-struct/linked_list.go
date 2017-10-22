package main

import (
	"fmt"
	"errors"
)

type Node struct {
	value int
	next  *Node
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

func add(l *List, val int) {
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

func get(l *List, i int) (int, error) {
	if i >= l.length {
		return 0, errors.New("To big index value")
	}
	node := l.head
	for idx := 0; idx < i; idx++ {
		node = node.next
	}
	return node.value, nil
}

func remove(l *List, i int) error {
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
	var l List
	add(&l, 10)
	add(&l, 20)
	add(&l, 30)
	add(&l, 40)
	add(&l, 50)
	remove(&l, 3)
	val, _ :=  get(&l, 2)
	fmt.Println("3 element value is ", val)
	val, _ =  get(&l, 3)
	fmt.Println("3 element value is ", val)
}
