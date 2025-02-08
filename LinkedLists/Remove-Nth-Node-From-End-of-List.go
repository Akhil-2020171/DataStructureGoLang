package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first, second := dummy, dummy
	for i := 0; i < n + 1; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

// Explanation
// We can remove the nth node from the end of a singly linked list in one pass by using two pointers. 
// We create a dummy node that points to the head of the list. We then create two pointers, first and second, 
// that both point to the dummy node. We move the first pointer n + 1 steps ahead of the second pointer. 
// We then move both pointers one step at a time until the first pointer reaches the end of the list. 
// At this point, the second pointer will be pointing to the node before the nth node from the end. 
// We can then remove the nth node by updating the next pointer of the second node to skip over the nth node. 
// Finally, we return the next pointer of the dummy node, which will be the head of the modified list.