package main

import "github.com/rafaelcmd/go-algorithms-and-datastructures/sets/stacks"

func main() {
	var stack = stacks.NewStack()
	stack.Top = stacks.NewStackNode(3)
	stack.Push(4)
	stack.Push(7)
	stack.Print()
	stack.Pop()
	stack.Print()
}
