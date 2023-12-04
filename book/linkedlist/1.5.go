package linkedlist

import (
	"fmt"
	"reflect"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	carry := 0
	val := 0
	head := dummy
	for l1 != nil || l2 != nil {
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		val = val1 + val2 + carry
		carry = val / 10
		val = val % 10
		dummy.Next = &ListNode{Val: val}
		dummy = dummy.Next
	}

	if carry != 0 {
		dummy.Next = &ListNode{Val: carry}
	}

	return head.Next
}
func Call1_5() {
	tests := map[int]InOutArray{
		1: {
			One:   []int{2, 4, 3},
			Two:   []int{5, 6, 4},
			Three: []int{7, 0, 8},
		},
		2: {
			One:   []int{0},
			Two:   []int{0},
			Three: []int{0},
		},
		3: {
			One:   []int{9, 9, 9, 9, 9, 9, 9},
			Two:   []int{9, 9, 9, 9},
			Three: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for key, v := range tests {
		if !reflect.DeepEqual(listToArray(addTwoNumbers(arrayToList(v.One), arrayToList(v.Two))), v.Three) {
			fmt.Println("1_5 failed for key %d", key)
		}
	}
}
