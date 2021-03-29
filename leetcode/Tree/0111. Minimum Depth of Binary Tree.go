package leetcode

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/* v0
执行用时：260 ms, 在所有 Go 提交中击败了65.90%的用户
内存消耗：21.8 MB, 在所有 Go 提交中击败了13.81%的用户
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left != nil {
		return minDepth(root.Left) + 1
	} else if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	return min(minDepth(root.Right), minDepth(root.Left)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
