package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy

	p1, p2 := list1, list2

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			head.Next = p1
			p1 = p1.Next
		} else {
			head.Next = p2
			p2 = p2.Next
		}
		head = head.Next
	}

	if p1 != nil {
		head.Next = p1
	} else {
		head.Next = p2
	}

	return dummy.Next
}
