package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		temp := curr
		curr = curr.Next
		temp.Next = prev
		prev = temp
	}

	return prev
}
