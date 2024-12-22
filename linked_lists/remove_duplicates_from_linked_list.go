package linked_lists

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// Time complexity: O(n)
// Space complexity: O(1)
func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	currentNode := linkedList

	for currentNode != nil {
		nextDistinctNode := currentNode.Next

		for nextDistinctNode != nil && nextDistinctNode.Value == currentNode.Value {
			nextDistinctNode = nextDistinctNode.Next
		}

		currentNode.Next = nextDistinctNode
		currentNode = nextDistinctNode
	}

	return linkedList
}
