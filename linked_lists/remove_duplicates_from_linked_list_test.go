package linked_lists

import "testing"

func BenchmarkRemoveDuplicatesFromLinkedList(b *testing.B) {
	linkedList := &LinkedList{Value: 1}
	linkedList.Next = &LinkedList{Value: 1}
	linkedList.Next.Next = &LinkedList{Value: 2}
	linkedList.Next.Next.Next = &LinkedList{Value: 3}
	linkedList.Next.Next.Next.Next = &LinkedList{Value: 3}
	linkedList.Next.Next.Next.Next.Next = &LinkedList{Value: 4}
	linkedList.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 5}
	linkedList.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 5}
	linkedList.Next.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 5}
	linkedList.Next.Next.Next.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 6}

	for i := 0; i < b.N; i++ {
		RemoveDuplicatesFromLinkedList(linkedList)
	}
}
