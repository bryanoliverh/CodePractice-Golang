/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    carrier := dummy

    for carrier.Next != nil && carrier.Next.Next != nil {
        first := carrier.Next
        second := first.Next

        // Swapping nodes
        first.Next = second.Next
        second.Next = first
        carrier.Next = second

        // Move the carrier pointer forward
        carrier = first
    }

    return dummy.Next
}
