package linkedlist

import (
	"fmt"
	"reflect"
)

func partition(head *ListNode, x int) *ListNode {
	left, right := &ListNode{}, &ListNode{}
	lleft, rright := left, right
	for head != nil {
		if head.Val < x {
			lleft.Next = head
			lleft = lleft.Next

		} else {
			rright.Next = head
			rright = rright.Next
		}

		head = head.Next
	}

	lleft.Next = right.Next
	rright.Next = nil
	return left.Next
}

func arrayToList(nums []int) *ListNode {
	head := &ListNode{}
	temp := head
	for _, v := range nums {
		temp.Next = &ListNode{Val: v}
		temp = temp.Next
	}
	displayList(head.Next)
	return head.Next
}

func listToArray(head *ListNode) []int {
	result := []int{}
	temp := head
	for temp != nil {
		result = append(result, temp.Val)
		temp = temp.Next
	}

	return result
}

func displayList(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
}

func Call1_4() {
	tests := map[int]InOutArray{
		3: {
			One: []int{1, 4, 3, 2, 5, 2},
			Two: []int{1, 2, 2, 4, 3, 5},
		},
		2: {
			One: []int{2, 1},
			Two: []int{1, 2},
		},
	}

	for key, v := range tests {
		if !reflect.DeepEqual(listToArray(partition(arrayToList(v.One), key)), v.Two) {
			fmt.Println("1_4 failed for key %d", key)
		}
	}
}

type InOutArray struct {
	One    []int
	Two    []int
	Three  []int
	Status bool
}
