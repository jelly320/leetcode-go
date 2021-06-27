package offer

import (
	"go-leetcode/structures"
)

/*
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = structures.ListNode

// 双指针，伪头节点
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	dum := &ListNode{Val: -1}

	cur := dum

	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next

		} else if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next

		}
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	//structures.PrintListNode(dum.Next)
	return dum.Next
}

// 递归
func mergeTwoLists_v1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var n *ListNode
	if l1.Val < l2.Val {
		n = l1
		n.Next = mergeTwoLists(l1.Next, l2)
	} else {
		n = l2
		n.Next = mergeTwoLists(l1, l2.Next)
	}
	return n
}
