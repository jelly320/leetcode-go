package offer

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

type ListNode struct {
	Val int
	Next *ListNode

}

func createListNode(lst []int) *ListNode {
	var res ListNode
	res.Val = lst[0]
	for i:=1;i<len(lst);i++ {

	}
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

}