package problems

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (p Problem) Run237() {
	fmt.Println("Problem 237: https://leetcode.com/problems/delete-node-in-a-linked-list/")
	numbers := []int{4, 5, 1, 9}
	head := &ListNode{numbers[0], nil}
	first := true
	curr := &ListNode{}
	for _, num := range numbers {
		if first {
			curr = head
			first = false
		} else {
			curr.Next = &ListNode{num, nil}
			curr = curr.Next
		}
	}
	deleteNode(head.Next)
	printList(head)
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Print("\n")
}
