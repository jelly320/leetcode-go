package leetcode

// https://leetcode-cn.com/problems/balanced-binary-tree/

// 平衡二叉树：每个节点的左右两个子树的高度差的绝对值不超过 1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*v0
执行用时：12 ms, 在所有 Go 提交中击败了50.53%的用户
内存消耗：5.7 MB, 在所有 Go 提交中击败了99.65%的用户
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDepth := balanceMaxDepth(root.Left)
	rightDepth := balanceMaxDepth(root.Right)
	//fmt.Printf("abc:%v,left:%v right:%v\n",abc(leftDepth-rightDepth),isBalanced(root.Left), isBalanced(root.Right))
	return abc(leftDepth-rightDepth) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func balanceMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(balanceMaxDepth(root.Left), balanceMaxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abc(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
