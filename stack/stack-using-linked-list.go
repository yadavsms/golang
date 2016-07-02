package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) push(k int) {
	new_node := &Node{k, nil}
	new_node.next = s.head
	s.head = new_node
}

func (s *Stack) pop() int {
	current := s.head
	s.head = current.next
	return current.data
}

func (s *Stack) remove(n int) {
	for i := s.head; i.next != nil; i = i.next {
		if i.next.data == n {
			i.next = i.next.next
		}
	}
}

func (s *Stack) peek() int {
	return s.head.data
}

func (s *Stack) all() {
	for i := s.head; i != nil; i = i.next {
		fmt.Println(i)
	}
}

func main() {
	var s Stack
	s.push(23)
	s.push(34)
	s.push(35)
	s.push(37)
	s.push(43)
	s.push(47)
	s.all()
	fmt.Println("Poping", s.pop())
	fmt.Println("Poping", s.pop())
	s.all()
	fmt.Println("peek", s.peek())
}
