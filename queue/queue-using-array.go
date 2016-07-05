package main

import (
	"fmt"
	"log"
)

const size = 10

const special_value = -1

type Queue struct {
	list  [size]int
	headindex  int
	tailindex  int
	count int
}

func (q *Queue) isEmpty() bool {
	if q.headindex == special_value { 
		return true
	}
	return false
}

func (q *Queue) isFull() bool {
	nextindex := (q.tailindex + 1) % size
	return nextindex == q.headindex
}

func (q *Queue) enqueue(data int) {
	if q.isFull() {
		log.Fatal("queue: Overflow")
	}
	q.tailindex = (q.tailindex + 1) % size
	q.list[q.tailindex] = data
	if q.headindex == special_value {
		q.headindex = q.tailindex
	}
}

func (q *Queue) offer(data int) bool {
	if q.isFull() {
		return false
	}
	q.tailindex = (q.tailindex + 1) % size
	q.list[q.tailindex] = data
	return true
}

func (q *Queue) dequeue() int {
	if q.isEmpty() {
		log.Fatal("queue: Underflow")
	}
	x := q.list[q.headindex]
	q.list[q.headindex] = 0
	if q.headindex == q.tailindex {
		q.headindex = special_value
	} else {
		q.headindex = (q.headindex + 1) % size
	}
	return x
}

func (q *Queue) peek() int {
	x := q.list[q.tailindex-1]
	return x
}

func main() {
	var q Queue
	q.headindex = -1
	q.tailindex = -1
	fmt.Println(q)
	for i := 1; i <= 10; i++ {
		q.enqueue(i)
	}
	fmt.Println(q)
	q.dequeue()
	q.dequeue()
	fmt.Println(q)
	q.dequeue()
	q.dequeue()
	q.enqueue(22)
	q.enqueue(23)
	q.enqueue(24)
	q.enqueue(25)
	q.offer(26)
	fmt.Println(q)
	fmt.Println(q.peek())
}
