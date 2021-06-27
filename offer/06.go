package offer

/*
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reversePrint(head *ListNode) []int {
	temp := make([]int, 0)

	for head != nil {
		temp = append(temp, head.Val)
		head = head.Next
	}
	i := 0
	j := len(temp) - 1
	for i < j {
		temp[i], temp[j] = temp[j], temp[i]
		i++
		j--
	}

	return temp
}

func reversePrint1(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint1(head.Next), head.Val)
}
