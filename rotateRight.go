/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func rotateRight(head *ListNode, k int) *ListNode {
    // Check if the linked list is empty or if k is 0
    if head == nil || k == 0 {
        return head
    }

    // Calculate the length of the linked list and store the tail node
    length := 1
    tail := head
    for tail.Next != nil {
        tail = tail.Next
        length++
    }

    // Calculate the effective rotation count
    k = k % length

    // If the effective rotation count is 0, return the original head
    if k == 0 {
        return head
    }

    // Connect the tail node to the head to form a circular list
    tail.Next = head

    // Find the new tail and new head positions
    newTail := head
    for i := 0; i < length-k-1; i++ {
        newTail = newTail.Next
    }
    newHead := newTail.Next

    // Break the circular list at the new tail position
    newTail.Next = nil

    return newHead
}
