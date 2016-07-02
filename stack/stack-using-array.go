package main

import (
	"fmt"
	"log"
)

const size = 10

type Array [size]int

var top int

func (s *Array) push(n int) {
	if s.isFull() {
		log.Fatal("Stack overflow")
	}
	(*s)[top] = n
	top++
}

func (s Array) isempty() bool {
	if top == 0 {
		return true
	}
	return false
}
func (s Array) isFull() bool {
	if top == size {
		return true
	}
	return false
}

func (s Array) peek() int {
	if s.isempty() {
		log.Fatal("stack is empty")
	}
	return s[top-1]
}

func (s *Array) pop() int {
	if top == 0 {
		log.Fatal("Stack underflow")

	}
	x := (*s)[top-1]
	copy((*s)[top-1:], (*s)[top:])
	top--
	return x

}

func main() {
	var s Array
	fmt.Println(s.isempty())
	s.push(23)
	s.push(45)
	s.push(45)
	s.push(45)
	s.push(45)
	s.push(45)
	fmt.Println(s.isFull())
	s.push(45)
	s.push(45)
	fmt.Println(s.isFull())
	fmt.Println(s.peek())
	s.pop()
	fmt.Println(s)
}
