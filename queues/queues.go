package queues

import "fmt"

//This solution uses a linked list structure to implement a queue

type Queue struct {
	Head *QueueNode
	Tail *QueueNode
}

type QueueNode struct {
	Val  int
	Next *QueueNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func NewQueueNode(val int) *QueueNode {
	return &QueueNode{Val: val}
}

func (q *Queue) Enqueue(val int) {
	node := &QueueNode{Val: val}
	if q.Head == nil {
		q.Head = node
		q.Tail = node
		return
	}
	q.Tail.Next = node
	q.Tail = node
}

func (q *Queue) Dequeue() int {
	val := q.Head.Val
	q.Head = q.Head.Next
	isEmpty := q.IsEmpty()
	if isEmpty {
		q.Tail = nil
	}
	return val
}

func (q *Queue) Peek() (int, error) {
	isEmpty := q.IsEmpty()
	if isEmpty {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.Head.Val, nil
}

func (q *Queue) IsEmpty() bool {
	return q.Head == nil
}

func (q *Queue) Print() {
	current := q.Head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}
