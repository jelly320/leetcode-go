package offer

/*

输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
    3
   / \
  9  20
    /  \
   15   7

     1
    / \
  2    3
 / \  / \
4  5 6   7
preorder = [1,2,4,5,3,6,7]
inorder = [4,2,5,1,6,3,7]
[1]
inorder = [4,2,5 | 1 | 6,3,7 ]
preorder = [1 | 2,4,5 | 3,6,7 ]


*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}
//
//func buildTree(preorder []int, inorder []int) *TreeNode {
//	var resTree TreeNode
//	//Root Node，为preorder的第一个值
//	resTree.Val=preorder[0]
//
//	//由于前提是无重复的Node，则在inorder中找出Root Node后，可以划分出 左子树|Root|右子树
//	for i,v:=range inorder{
//		if v==resTree.Val {
//			leftTree:=inorder[0:i]
//			rightTree:=inorder[i+1:len(inorder)]
//		}
//	}
//
//
//	return nil
//
//}
