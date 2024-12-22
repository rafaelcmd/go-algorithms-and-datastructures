package main

import (
	"github.com/rafaelcmd/go-algorithms-and-datastructures/queues"
)

func main() {
	queue := queues.NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Print()

	queue.Dequeue()
	queue.Print()
}
