package linkedlist

import "fmt"

func deleteMiddle(head *ListNode) *ListNode {
	temp := head
	list := []int{}
	for temp != nil {
		list = append(list, temp.Val)
		temp = temp.Next
	}

	var middleElementIndex int = len(list) / 2
	if middleElementIndex <= 0 {
		return nil
	}

	list = append(list[:middleElementIndex], list[middleElementIndex+1:]...)

	head = &ListNode{Val: list[0]}
	temp = head
	for i := 1; i < len(list); i++ {
		node := &ListNode{Val: list[i]}
		temp.Next = node
		temp = node
	}

	return head
}

func Call1_3() {
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

	n = deleteMiddle(n)
	fmt.Println(`------- Delete Middle Element -------`)
	for n != nil {

		fmt.Println(n.Val)
		n = n.Next
	}
}
