package leetcode

import "fmt"

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	//leftDepth:=1
	//rightDepth:=1
	if root == nil {
		return 0
	}
	var result []int
	var depth int
	operator(root, &depth, &result)
	fmt.Println(result)
	//operator(root, &leftDepth, &rightDepth)
	//if leftDepth >= rightDepth {
	//	return leftDepth
	//} else {
	//	return rightDepth
	//}
	return 1

}

//func operator(root *TreeNode, left_depth *int,right_depth *int,result *[]int)  {
//	if root != nil {
//		*depth++
//		operator(root.Left,left_depth,result)
//		operator(root.Right,right_depth,result)
//	}
//	*result=append(*result,*depth)
//}

//func operator(root *TreeNode, leftDepth *int, rightDepth *int) {
//	if root != nil {
//		if root.Left != nil {
//			*leftDepth++
//			operator(root.Left, leftDepth, rightDepth)
//		}
//		if root.Right != nil {
//			*rightDepth++
//			operator(root.Right, leftDepth, rightDepth)
//		}
//	}
//}
