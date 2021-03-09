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

/*v1 学习BFS解法后:
执行用时： beats 11.35%
内存消耗：beats 48.38%
*/
func levelOrder_v1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	currentNum, nextLevelNum := 1, 0
	var tmp []int
	var result [][]int
	var queue []*TreeNode

	//根节点首先入队
	queue = append(queue, root)

	for len(queue) != 0 {
		if currentNum > 0 {
			top := queue[0]
			if top.Left != nil {
				queue = append(queue, top.Left)
				nextLevelNum++
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
				nextLevelNum++
			}
			currentNum--
			tmp = append(tmp, top.Val)
			queue = queue[1:]
		}
		if currentNum == 0 {
			result = append(result, tmp)
			currentNum = nextLevelNum
			nextLevelNum = 0
			tmp = []int{}
		}

	}
	return result
}

/*v0:
执行用时： beats 100%
内存消耗： beats 5.44%
*/
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
