package main

import (
	"errors"
	"fmt"
)

type StackNode struct { //TODO some packaging
	value int //TODO some generalization, template, generic type?
	next *StackNode
}

type Stack struct {
	maxDepth int
	depth int
	head *StackNode
}

func (s *Stack) push(val int) error {
	if s.depth + 1 > s.maxDepth {
		return errors.New("Stack overflow:-)")
	}

	s.head = &StackNode{val, s.head}
	s.depth++

	return nil
}

func (s *Stack) pop() (int, error) {
	if s.depth == 0 {
		return 0, errors.New("Stack is empty")
	}

	curr := s.head
	s.head = s.head.next
	s.depth--

	return curr.value, nil
}

func (s *Stack) empty() bool {
	return s.depth == 0
}

func main() {
	s := Stack{maxDepth: 100}
	s.push(100)
	s.push(200)
	s.push(300)
	fmt.Println("Empty?", s.empty())
	val, _ := s.pop()
	fmt.Println(val)
	val, _ = s.pop()
	fmt.Println(val)
	val, _ = s.pop()
	fmt.Println(val)
	fmt.Println("Empty?", s.empty())
}