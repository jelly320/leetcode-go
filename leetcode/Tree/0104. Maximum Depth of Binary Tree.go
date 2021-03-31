package leetcode

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
	if root == nil {
		return 0
	}
	return max104(maxDepth(root.Right), maxDepth(root.Left)) + 1
}

func max104(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
