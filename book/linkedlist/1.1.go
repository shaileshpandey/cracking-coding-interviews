package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

func removeDups_Bruteforce(root *Node) {
	tmp1 := root

	for tmp1 != nil {
		tmp2 := tmp1
		for tmp2 != nil && tmp2.next != nil {
			if tmp1.data == tmp2.next.data {
				tmp2.next = tmp2.next.next
			} else {
				tmp2 = tmp2.next
			}
		}

		tmp1 = tmp1.next
	}
}

func removeDups_Map(root *Node) *Node {
	uniqueValue := map[int]bool{}
	tmp := root
	for tmp != nil {
		uniqueValue[tmp.data] = true
		tmp = tmp.next
	}
	var dummyNode *Node
	for k := range uniqueValue {
		dummyNode = insertNode(dummyNode, k)
	}

	return dummyNode
}

func insertNode(root *Node, element int) *Node {
	tmpNode := Node{data: element}
	if root == nil {
		root = &tmpNode
	} else {
		curr := root
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &tmpNode
	}

	return root
}

func Call1_1() {
	n, m := &Node{
		data: 1,
		next: &Node{
			data: 1,
			next: &Node{
				data: 1,
				next: &Node{
					data: 1,
					next: &Node{
						data: 2,
					},
				},
			},
		},
	}, &Node{
		data: 1,
		next: &Node{
			data: 1,
			next: &Node{
				data: 1,
				next: &Node{
					data: 1,
					next: &Node{
						data: 2,
					},
				},
			},
		},
	}

	removeDups_Bruteforce(n)

	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}

	r := removeDups_Map(m)

	for r != nil {
		fmt.Println(r.data)
		r = r.next
	}
}
