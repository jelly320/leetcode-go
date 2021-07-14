package Link

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// 把链表复制到数组
	cur := head
	lst := []int{}
	for cur != nil {
		lst = append(lst, cur.Val)
		cur = cur.Next
	}

	// 双"指针"
	j := len(lst) - 1

	for i := range lst {
		if lst[i] == lst[j] {
			j--
		} else {
			return false
		}
	}
	return true
}
