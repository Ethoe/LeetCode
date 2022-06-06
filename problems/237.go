package problems

import "fmt"

func (p Problem) Run237() {
	fmt.Println("Problem 237: https://leetcode.com/problems/delete-node-in-a-linked-list/")
	numbers := []int{4, 5, 1, 9}
	head := new(ListNode)
	head.Init(numbers)

	deleteNode(head.Next)
	printSingleNodeList(head)
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
