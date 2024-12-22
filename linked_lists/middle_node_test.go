package linked_lists

import "testing"

func BenchmarkMiddleNode1(b *testing.B) {
	linkedList := &LinkedList{Value: 1}
	linkedList.Next = &LinkedList{Value: 2}
	linkedList.Next.Next = &LinkedList{Value: 3}
	linkedList.Next.Next.Next = &LinkedList{Value: 4}
	linkedList.Next.Next.Next.Next = &LinkedList{Value: 5}
	linkedList.Next.Next.Next.Next.Next = &LinkedList{Value: 6}
	linkedList.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 7}

	for i := 0; i < b.N; i++ {
		MiddleNode1(linkedList)
	}
}

func BenchmarkMiddleNode2(b *testing.B) {
	linkedList := &LinkedList{Value: 1}
	linkedList.Next = &LinkedList{Value: 2}
	linkedList.Next.Next = &LinkedList{Value: 3}
	linkedList.Next.Next.Next = &LinkedList{Value: 4}
	linkedList.Next.Next.Next.Next = &LinkedList{Value: 5}
	linkedList.Next.Next.Next.Next.Next = &LinkedList{Value: 6}
	linkedList.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 7}

	for i := 0; i < b.N; i++ {
		MiddleNode2(linkedList)
	}
}

func BenchmarkMiddleNode3(b *testing.B) {
	linkedList := &LinkedList{Value: 1}
	linkedList.Next = &LinkedList{Value: 2}
	linkedList.Next.Next = &LinkedList{Value: 3}
	linkedList.Next.Next.Next = &LinkedList{Value: 4}
	linkedList.Next.Next.Next.Next = &LinkedList{Value: 5}
	linkedList.Next.Next.Next.Next.Next = &LinkedList{Value: 6}
	linkedList.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 7}

	for i := 0; i < b.N; i++ {
		MiddleNode3(linkedList)
	}
}
