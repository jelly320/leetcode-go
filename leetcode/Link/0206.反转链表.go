package Link

import "go-leetcode/structures"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	// 双指针
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next

	}
	return behind

}

//TODO: 和我的思路是一致的但为什么我写的跑不了，要体会下

//func reverseList(head *ListNode) *ListNode {
//	// 双指针
//	pre := &ListNode{Val: head.Val, Next: nil}
//	for head != nil {
//		cur := head
//		head = head.Next
//		pre = head
//		pre.Next = cur
//	}
//	return pre
//
//}
