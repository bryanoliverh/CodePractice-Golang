func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // Create a dummy node as the head of the list
    dummy := &ListNode{Next: head}

    // Initialize two pointers
    slow := dummy
    fast := dummy

    // Move the fast pointer n nodes ahead
    for i := 1; i <= n+1; i++ {
        fast = fast.Next
    }

    // Move both pointers until the fast pointer reaches the end of the list
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }

    // Remove the nth node by adjusting the pointers
    slow.Next = slow.Next.Next

    // Return the updated head of the list
    return dummy.Next
}
