package leetcode

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
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

/*v0:*/
func levelOrder(root *TreeNode) [][]int {

	init := []*TreeNode{root}
	var result [][]int
	initvalue := getNodeValue(init)
	if initvalue != nil {
		result = append(result, initvalue)
	}
	Just(init, &result)
	return result
}

func Just(root []*TreeNode, output *[][]int) {
	//init:= []*TreeNode{root}
	//result := &([][]*TreeNode{})
	if root != nil {
		newNode := getNode(root)
		nodevalue := getNodeValue(newNode)
		if nodevalue != nil {
			*output = append(*output, nodevalue)
		}
		Just(newNode, output)
	}
}

func getNodeValue(rootList []*TreeNode) []int {
	var result []int
	if rootList != nil {
		for _, v := range rootList {
			if v != nil {
				result = append(result, v.Val)
			}
		}
		return result
	}
	return nil
}

func getNode(rootList []*TreeNode) []*TreeNode {
	var output []*TreeNode

	if rootList != nil {
		for _, v := range rootList {
			if v != nil {
				output = append(output, v.Left)
				output = append(output, v.Right)
			}
		}
		return output
	}
	return nil

}
