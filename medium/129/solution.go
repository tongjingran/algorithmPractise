package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	left := TreeNode{2, nil, nil}
	right := TreeNode{3, nil, nil}
	input := struct {
		root *TreeNode
	}{&TreeNode{1, &left, &right}}
	output := 25
	if sumNumbers(input.root) != output {
		panic("wrong answer")
	}
	fmt.Println("pass")
	return
}

/*
	sum numbers of bitTree
	time: O(n)
	memory: O(n)
*/
func sumNumbers(root *TreeNode) int {
	return roll(root, 0)
}

func roll(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = 10*sum + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return roll(root.Left, sum) + roll(root.Right, sum)
}
