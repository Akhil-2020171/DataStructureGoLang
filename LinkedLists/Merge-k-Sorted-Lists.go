package main

func mergeKLists(lists []*ListNode) *ListNode {
    size := len(lists)

    if size == 0 {
        return nil
    }

    if size == 1 {
        return lists[0]
    }

    var head *ListNode = lists[0]

    for i := 1; i < size; i++ {
        head = mergeTwoListsIterative(head, lists[i])
    }

    return head
}
