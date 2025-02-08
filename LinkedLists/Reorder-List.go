package main

func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}
	
	// Find the middle of the list
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	// Reverse the second half of the list
	prev, curr := slow, slow.Next
	for curr.Next != nil {
		temp := curr.Next
		curr.Next = temp.Next
		temp.Next = prev.Next
		prev.Next = temp
	}
	
	// Reorder the list
	p1, p2 := head, prev.Next
	for p1 != prev {
		prev.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = prev.Next
	}
}