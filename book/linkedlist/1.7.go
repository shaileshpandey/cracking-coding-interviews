package linkedlist

import (
	"fmt"
	"reflect"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := getLength(headA)
	lenB := getLength(headB)

	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	}

	if lenB > lenA {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}
	for headA != nil {
		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func getLength(head *ListNode) int {
	count := 0
	for head != nil {
		count = count + 1
		head = head.Next
	}

	return count
}

func Call1_7() {
	tests := map[int]InOutArray{
		1: {
			One:   []int{4, 1, 8, 4, 5},
			Two:   []int{5, 6, 1, 8, 4, 5},
			Three: []int{8},
		},
		2: {
			One:   []int{1, 9, 1, 2, 4},
			Two:   []int{1, 9, 1, 2, 4},
			Three: []int{3},
		},
	}

	for key, v := range tests {
		if !reflect.DeepEqual(listToArray(getIntersectionNode(arrayToList(v.One), arrayToList(v.Two))), v.Three) {
			fmt.Println("1_7 failed for key %d", key)
		}
	}
}
