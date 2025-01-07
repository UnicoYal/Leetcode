package main

import "fmt"


type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	fmt.Println(mergeTwoLists(nil, nil))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var it = &ListNode{}
	head := it

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			it.Next = list2
			list2 = list2.Next
		} else {
			it.Next = list1
			list1 = list1.Next
		}

		it = it.Next
	}

	if list1 != nil {
		it.Next = list1
	} else {
		it.Next = list2
	}

	return head.Next
}
