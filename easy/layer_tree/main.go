package main

import "fmt"

func main() {
	var root1 *TreeNode
	fmt.Println("Test 1 (Empty Tree):", levelOrderTraversal(root1))

	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2 (Single Node):", levelOrderTraversal(root2))

	root3 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println("Test 3 (Two Levels):", levelOrderTraversal(root3))

	root4 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
	}
	fmt.Println("Test 4 (Three Levels):", levelOrderTraversal(root4))

	root5 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println("Test 5 (Unbalanced Tree):", levelOrderTraversal(root5))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		res = append(res, first.Val)
		if first.Left != nil {
			queue = append(queue, first.Left)
		}
		if first.Right != nil {
			queue = append(queue, first.Right)
		}
	}

	return res
}
