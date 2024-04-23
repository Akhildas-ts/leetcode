package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	// Calculate the length of the linked list
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	// Reset current to the head of the list
	current = head

	// If n is equal to the length of the list, remove the first node
	if n == length {
		head = head.Next
		return head
	}

	// Traverse the list again to find the node before the nth node from the end
	for i := 1; i < length-n; i++ {
		current = current.Next
	}

	// Remove the nth node from the end
	if current.Next != nil {
		current.Next = current.Next.Next
	}

	return head

}
