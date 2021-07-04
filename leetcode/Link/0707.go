package Link

/*
设计链表的实现。您可以选择使用单链表或双链表。
单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。
如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
*/

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{Val: -1, Next: nil} //这样的初始化方式意味着是存在一个虚拟的头结点的？
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	cur := this

	// for中的条件，不一定是只能判断i，get了
	for i := 0; cur != nil; i++ {
		if i == index {
			if cur.Val == -999 {
				return -1
			} else {
				return cur.Val
			}
		}
		cur = cur.Next
	}

	return -1

}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.Val == -999 {
		this.Val = val
	} else {
		tmp := &MyLinkedList{Val: this.Val, Next: this.Next}
		this.Val = val
		this.Next = tmp
	}

}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	cur := this
	for cur.Next != nil {
		cur = cur.Next
	}
	tmp := &MyLinkedList{Val: val, Next: nil}
	cur.Next = tmp
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this

	if index == 0 {
		this.AddAtHead(val)
		return
	}
	for i := 0; cur != nil; i++ {
		if i == index-1 {
			break
		}
		cur = cur.Next
	}
	tmp := &MyLinkedList{Val: val, Next: cur.Next}
	cur.Next = tmp
}

//func (this *MyLinkedList) DeleteAtHead() {
//
//
//}

//func (this *MyLinkedList) DeleteAtTail() {
//
//}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this

	for i := 0; cur != nil; i++ {
		if i == index-1 {
			break
		}
		cur = cur.Next
	}

	if cur != nil && cur.Next != nil {
		cur.Next = cur.Next.Next
	}
}
