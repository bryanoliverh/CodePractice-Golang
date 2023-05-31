package main

import "fmt"

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addNextNodeWithValue(head *ListNode, targetValue, newValue int) *ListNode {
	current := head

	for current != nil {
		if current.Val == targetValue {
			newNode := &ListNode{Val: newValue}
			newNode.Next = current.Next
			current.Next = newNode
			break
		}
		current = current.Next
	}

	return head
}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// Create the linked list [1, 2, 3, 4, 5]
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	fmt.Println("Original linked list:")
	printLinkedList(head)

	// Add 5 as the next node of any node with value 3
	newHead := addNextNodeWithValue(head, 3, 5)

	fmt.Println("Updated linked list:")
	printLinkedList(newHead)
}
