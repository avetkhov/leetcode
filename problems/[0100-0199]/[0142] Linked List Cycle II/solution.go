package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			p1, p2 := head, slow

			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next
			}

			return p1
		}
	}

	return nil
}
