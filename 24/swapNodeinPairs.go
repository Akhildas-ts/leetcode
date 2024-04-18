package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	if head != nil && head.Next == nil {
		return head
	}

	if head == nil {

		return nil
	}

	current := head

	var prev *ListNode
	var newHead *ListNode

	for current != nil && current.Next != nil {

		firstNode := current
		secondNode := current.Next

		// Adjust the pointers to swap the nodes
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		// Update the head of the list if the swap occurs at the beginning
		if current == head {
			head = secondNode
		}

		if prev != nil {
			prev.Next = secondNode
		} else {
			newHead = secondNode
		}
		// Move to the next pair of nodes
		prev = firstNode
		current = firstNode.Next
	}

	return newHead

}
