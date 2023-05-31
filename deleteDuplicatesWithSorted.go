func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    // Sort the linked list
    head = mergeSort(head)
    
    // Remove duplicates
    current := head
    for current != nil && current.Next != nil {
        if current.Val == current.Next.Val {
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }
    
    return head
}

func mergeSort(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    // Find the middle node
    var prev *ListNode
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    // Split the linked list into two halves
    prev.Next = nil
    
    // Recursively sort the two halves
    left := mergeSort(head)
    right := mergeSort(slow)
    
    // Merge the sorted halves
    return merge(left, right)
}

func merge(left *ListNode, right *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    
    for left != nil && right != nil {
        if left.Val < right.Val {
            current.Next = left
            left = left.Next
        } else {
            current.Next = right
            right = right.Next
        }
        current = current.Next
    }
    
    if left != nil {
        current.Next = left
    }
    if right != nil {
        current.Next = right
    }
    
    return dummy.Next
}
