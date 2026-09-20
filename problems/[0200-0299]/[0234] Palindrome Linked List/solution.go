package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}

func isPalindrome(head *ListNode) bool {
	middle := getMiddleNode(head)
	tail := reverseList(middle)

	for tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head, tail = head.Next, tail.Next
	}

	return true
}
