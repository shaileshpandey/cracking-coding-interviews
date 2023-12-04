package linkedlist

import (
	"fmt"
)

func isPalindrome(head *ListNode) bool {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		nxt := slow.Next
		slow.Next = prev
		prev = slow
		slow = nxt
	}

	for prev != nil && head != nil {
		if prev.Val != head.Val {
			return false
		}

		prev = prev.Next
		head = head.Next
	}

	return true
}

func Call1_6() {
	tests := map[int]InOutArray{
		1: {
			One:    []int{1, 2, 2, 1},
			Status: true,
		},
		2: {
			One:    []int{1, 2},
			Status: false,
		},
	}

	for key, v := range tests {
		if isPalindrome(arrayToList(v.One)) != v.Status {
			fmt.Println("1_6 failed for key %d", key)
		}
	}
}
