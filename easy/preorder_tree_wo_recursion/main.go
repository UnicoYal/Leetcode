package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func preorderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	if root == nil { return res }
// 	stack := make([]*TreeNode, 0)
// 	stack = append(stack, root)

// 	for len(stack) > 0 {
// 		last := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		res = append(res, last.Val)

// 		if last.Right != nil { stack = append(stack, last.Right)}
// 		if last.Left != nil { stack = append(stack, last.Left)}
// 	}

// 	return res
// }

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)

	return res
}
