package leetcode

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

import (
	"fmt"
	"go-leetcode/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法一 递归
func preorderTraversal(root *TreeNode, tarnum int) []int {
	tarnum++
	fmt.Printf("===tag: start tarnum: %d===\n", tarnum)
	res := []int{}
	if root != nil {
		fmt.Printf("---tag: forstart---\n")
		res = append(res, root.Val)
		tmp := preorderTraversal(root.Left, tarnum) //if tmp= []int{} ,是不会执行for _,t range tmp的
		for _, t := range tmp {
			fmt.Printf("tag: for root.Left\n")
			res = append(res, t)
		}
		tmp = preorderTraversal(root.Right, tarnum)
		for _, t := range tmp {
			fmt.Printf("tag: for root.Right\n")
			res = append(res, t)
		}
		fmt.Printf("---tag:end for. res: %v---\n", res)
	}
	fmt.Printf("===tag: END tagnum:%v ===\n", tarnum)
	return res
}

// 解法二 递归
func preorderTraversal1(root *TreeNode) []int {
	var result []int
	preorder(root, &result)
	return result
}

func preorder(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		preorder(root.Left, output)
		preorder(root.Right, output)
	}
}

// 解法三 非递归，用栈模拟递归过程
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*TreeNode{}, []int{}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

func main() {
	print()
}
