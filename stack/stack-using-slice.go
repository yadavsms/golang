package main

import (
	"fmt"
)

type Stack []int

func (s Stack) isEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func (s Stack) peek() int {
	return s[len(s)-1]
}

func (s *Stack) push(n int) {
	*s = append(*s, n)
}

func (s *Stack) pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func main() {
	var s Stack
	fmt.Println(s.isEmpty())
	s.push(1)
	s.push(3)
	s.push(4)
	fmt.Println(s.isEmpty())
	fmt.Println(s)
	fmt.Println(s.peek())
	fmt.Println(s.pop())
	fmt.Println(s)
	fmt.Println(s.peek())
}
