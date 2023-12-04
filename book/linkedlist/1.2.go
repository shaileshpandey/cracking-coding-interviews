package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	temp := head
	items := []int{}

	for temp != nil {
		count = count + 1
		items = append(items, temp.Val)
		temp = temp.Next
	}

	items = append(items[:count-n], items[count-n+1:]...)

	if len(items) == 0 {
		return nil
	}

	head = &ListNode{Val: items[0]}
	t := head
	for i := 1; i < len(items); i++ {
		newNode := &ListNode{Val: items[i]}
		t.Next = newNode
		t = t.Next
	}

	return head
}

func Call1_2() {
	n := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	n = removeNthFromEnd(n, 3)
	fmt.Println(`------- Removing 3rd Element from last -------`)
	for n != nil {

		fmt.Println(n.Val)
		n = n.Next
	}
}
