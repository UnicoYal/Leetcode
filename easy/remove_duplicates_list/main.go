package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	i := head
	j := head.Next

	for j != nil {
		if i.Val != j.Val {
			i.Next = j
			i = i.Next
		} else if i.Val == j.Val && j.Next == nil {
			i.Next = nil
		}

		j = j.Next
	}

	return head
}
