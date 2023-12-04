package linkedlist

import "fmt"

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	var cycleFound bool
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			cycleFound = true
			break
		}
	}

	if !cycleFound {
		return nil
	}

	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

func Call1_8() {
	// Please Test this https://leetcode.com/problems/linked-list-cycle-ii/
	fmt.Println("Please test it https://leetcode.com/problems/linked-list-cycle-ii/")
}
