package Link

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// TODO: evenHead后续的Next都是怎么接上去的？
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

//func oddEvenList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	// 奇数链表
//	odd := head
//	// 偶数链表
//	evenHead := head.Next
//	even:=evenHead
//
//	for {
//		// 更新奇数节点，指向偶数节点的下一个节点
//		if even.Next != nil {
//			odd.Next = even.Next
//			odd = odd.Next
//		} else {
//			break
//		}
//
//		//
//		if odd.Next != nil {
//			even.Next = odd.Next
//			even = even.Next
//		} else {
//			even.Next=nil
//			break
//		}
//	}
//
//	odd.Next=evenHead
//	return odd
//}
