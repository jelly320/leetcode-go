package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//功能：打印节点的值
//参数：nil
//返回值：nil
func (node *TreeNode) Print() {
	fmt.Printf("%d ", node.Val)
}

//功能：创建节点
//参数：节点的值
//返回值：nil
func CreateNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

func preorderTraversal(root *TreeNode) []int {
	//var result []int
	//递归
	if root != nil {
		result = append(result, root.Val)
		preorderTraversal(root.Left)
		preorderTraversal(root.Right)
	}
	return result
}

func main() {

	root := CreateNode(1)
	//root.Left = CreateNode(nil)
	root.Right = CreateNode(2)
	root.Right.Right = CreateNode(3)

	aa := preorderTraversal(root)
	for _, v := range aa {
		println(v)
	}
}
