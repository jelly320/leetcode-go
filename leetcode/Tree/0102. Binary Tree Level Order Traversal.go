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

func levelOrder(root *TreeNode) [][]int {

	init := []*TreeNode{root}
	var result [][]int
	result = append(result, getNodeValue(init))
	Just(init, &result)
	return result
}

func Just(root []*TreeNode, output *[][]int) {
	//init:= []*TreeNode{root}
	//result := &([][]*TreeNode{})
	if root != nil {
		newNode := getNode(root)
		*output = append(*output, getNodeValue(newNode))
		Just(newNode, output)
	}
}

func getNodeValue(rootList []*TreeNode) []int {
	var result []int
	for _, v := range rootList {
		if v != nil {
			result = append(result, v.Val)
		}
	}
	return result
}

func getNode(rootList []*TreeNode) []*TreeNode {
	var output []*TreeNode

	if rootList != nil {
		for _, v := range rootList {
			if v != nil {
				output = append(output, v.Left)
				output = append(output, v.Right)
			} else {
				return nil
			}
		}
	} else {
		return nil
	}
	return output
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
