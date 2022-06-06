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

func (ln *ListNode) Init(values []int) {
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
	ln.Next = head.Next
	ln.Val = head.Val
}

func printSingleNodeList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Print("\n")
}