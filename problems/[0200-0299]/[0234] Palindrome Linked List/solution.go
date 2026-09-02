package main

//Example 1:
//
//
//Input: head = [1,2,2,2,2,1]
//Output: true
//Example 2:
//
//
//Input: head = [1,2]
//Output: false

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	var pre *ListNode
	cur := slow.Next
	for cur != nil {
		t := cur.Next
		cur.Next = pre
		pre = cur
		cur = t
	}
	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre, head = pre.Next, head.Next
	}
	return true
}

func isPalindrome(head *ListNode) bool {
	dummy := &ListNode{}
	current := head

	for current != nil {
		dummy.Val = current.Val
		current = current.Next
		dummy = &ListNode{
			Next: dummy,
		}
	}

	cur := dummy.Next

	x := head == cur
	return x
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	isPalindrome(list)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
