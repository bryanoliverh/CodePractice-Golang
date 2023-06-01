func getIntersectionNode(headA, headB *ListNode) *ListNode {
    taverseA := headA
    taverseB := headB
    for taverseA!=taverseB {
        if taverseA!=nil {
            taverseA = taverseA.Next
        }else{
            taverseA = headA
        }
        if taverseB!=nil {
            taverseB = taverseB.Next
        }else{
            taverseB = headB
        }
    }
    return taverseA
}
