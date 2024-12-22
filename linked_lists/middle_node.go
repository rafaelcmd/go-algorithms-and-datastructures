package linked_lists

// Time complexity: O(n)
// Space complexity: O(1)
func MiddleNode1(linkedList *LinkedList) *LinkedList {
	currentNode := linkedList
	totalNodes := 0
	visitedNodes := 0

	for currentNode != nil {
		totalNodes++
		currentNode = currentNode.Next
	}

	currentNode = linkedList

	for visitedNodes < totalNodes/2 {
		currentNode = currentNode.Next
		visitedNodes++
	}

	return currentNode
}

// Time complexity: O(n)
// Space complexity: O(1)
func MiddleNode2(linkedList *LinkedList) *LinkedList {
	currentNode := linkedList
	totalNodes := 0

	for currentNode != nil {
		totalNodes++
		currentNode = currentNode.Next
	}

	middleNode := linkedList

	for i := 0; i < totalNodes/2; i++ {
		middleNode = middleNode.Next
	}

	return middleNode
}

// Time complexity: O(n)
// Space complexity: O(1)
func MiddleNode3(linkedList *LinkedList) *LinkedList {
	slowNode := linkedList
	fastNode := linkedList

	for fastNode != nil && fastNode.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}

	return slowNode
}
