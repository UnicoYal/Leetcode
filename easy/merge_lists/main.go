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
	it1 := list1
	it2 := list2
	var head *ListNode

	if it1 == nil && it2 == nil {
		return it1
	} else if it1 == nil {
		return it2
	} else if it2 == nil {
		return it1
	} else {
		if it1.Val > it2.Val {
			head = it2
			it2 = it2.Next
		} else {
			head = it1
			it1 = it1.Next
		}
	}

	it := head

	for it1 != nil && it2 != nil {
		if it1.Val > it2.Val {
			it.Next = it2
			it2 = it2.Next
		} else {
			it.Next = it1
			it1 = it1.Next
		}
		it = it.Next
	}

	if it1 == nil {
		it.Next = it2
	} else {
		it.Next = it1
	}

	return head
}
