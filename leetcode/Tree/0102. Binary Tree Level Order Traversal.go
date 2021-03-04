package leetcode

// https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xej9yc/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

   3
  / \
 9  2
   /  \
  15   7

[
  [3],
  [9,20],
  [15,7]
]
*/

/*大概思路:
1. 每个node是要做一次Preoder
2.
*/

//func levelOrder(root *TreeNode) [][]int {
//
//	result := &([]*TreeNode{})
//	getNode(bb, result)
//
//}

func Just(root *TreeNode) [][]*TreeNode {
	//init:= []*TreeNode{root}
	result := &([][]*TreeNode{})
	getNode([]*TreeNode{root}, result)
	return *result
}

func getNodeValue(rootList []*TreeNode) []int {
	result := []int{}
	for _, v := range rootList {
		if v != nil {
			result = append(result, v.Val)
		}
	}
	return result
}

func getNode(rootList []*TreeNode, output *[][]*TreeNode) {
	tmp := []*TreeNode{}
	if rootList != nil {
		for _, v := range rootList {
			if v != nil {
				tmp = append(tmp, v.Left)
				tmp = append(tmp, v.Right)
				//*output = append(*output, v.Left)
				//*output = append(*output, v.Right)
			}
		}
		*output = append(*output, tmp)
	}
	getNode(tmp, output)
}

//func getRootVal(root *TreeNode) int {
//	return root.Val
//}
//
//func getLeftVal(root *TreeNode) int {
//	return root.Left.Val
//}
//
//func getRightVal(root *TreeNode) int {
//	return root.Right.Val
//}
