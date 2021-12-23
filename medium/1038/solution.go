package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	right := &TreeNode{Val: 1}
	input := &TreeNode{Val: 0, Right: right}
	output := bstToGst(input)
	fmt.Println(output)
	return
}

/*
	right -> root -> left
	time: O(n)
	memory: O(n)
*/
var sum int

func bstToGst(root *TreeNode) *TreeNode {
	sum = 0
	travel(root)
	return root
}

func travel(root *TreeNode) {
	if root == nil {
		return
	}
	travel(root.Right)
	sum += root.Val
	root.Val = sum
	travel(root.Left)
	return
}
