package structures

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 构造单链表
func CreateOneListNode(lst []int) *ListNode {
	n := len(lst)

	if n == 0 {
		return nil
	}

	root := &ListNode{
		Val: lst[0],
	}
	temp := root
	for i := 1; i < n; i++ {
		temp.Next = &ListNode{Val: lst[i]}
		temp = temp.Next
	}
	return root

}

// 临时用于输出ListNode结果
func PrintListNode(l *ListNode) {
	var res []int
	if l == nil {
		fmt.Printf("nil")
	}
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	fmt.Printf("%#v", res)
}
