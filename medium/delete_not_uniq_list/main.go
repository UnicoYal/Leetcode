package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	i := &ListNode{0, head}
	nh := i
	j := head

	for j != nil {
		for j.Next != nil && j.Val == j.Next.Val {
			j = j.Next
		}

		if i.Next == j {
			i = i.Next
		} else {
			i.Next = j.Next
		}

		j = j.Next
	}

	return nh.Next
}
