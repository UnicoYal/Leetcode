package main

import "fmt"

func main() {
	list3 := ListNode{3, nil}
	list2 := ListNode{4, &list3}
	list1 := ListNode{2, &list2}

	llist3 := ListNode{4, nil}
	llist2 := ListNode{6, &llist3}
	llist1 := ListNode{5, &llist2}

	fmt.Println(addTwoNumbers(&list1, &llist1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	it1 := l1
	it2 := l2
	var res = &ListNode{}
	head := res
	var ost int

	for it1 != nil || it2 != nil {
		sum := 0
		if it1 != nil {
			sum += it1.Val
			it1 = it1.Next
		}

		if it2 != nil {
			sum += it2.Val
			it2 = it2.Next
		}

		sum += ost
		ost = calculateOst(sum)
		sum %= 10
		res.Next = &ListNode{Val: sum}
		res = res.Next
	}

	if ost == 1 {
		res.Next = &ListNode{1, nil}
	}

	return head.Next
}

func calculateOst(sum int) int {
	if sum > 9 {
		return 1
	} else {
		return 0
	}
}
