package problems

import "fmt"

func (p Problem) Run160() {
	fmt.Println("Problem 160: https://leetcode.com/problems/intersection-of-two-linked-lists/")
	printSingleNodeList(getIntersectionNode(nil, nil))
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
   	return a 
}