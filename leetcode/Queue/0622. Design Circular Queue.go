package Queue

import (
	"fmt"
)

// https://leetcode-cn.com/problems/design-circular-queue/

type MyCircularQueue struct {
	queue []int
	head  int
	tail  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k, k),
		head:  -1,
		tail:  -1,
	}
}

// 入队
func (this *MyCircularQueue) EnQueue(value int) bool {
	// 成功则返回true,反之返回false
	// 先移动再赋值
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.head, this.tail = 0, 0
	} else if this.tail == len(this.queue)-1 {
		this.tail = 0
	} else {
		this.tail++
	}
	this.queue[this.tail] = value
	return true
}

// 出队
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	if this.head == len(this.queue)-1 {
		this.head = 0
	} else if this.head == this.tail {
		this.head, this.tail = -1, -1
	} else {
		this.head++
	}
	// this.queue = this.queue[this.head+1:] 不采取改变queue的方法,而是移动head tail

	return true
}

// 获取队首元素
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

// 获取队尾元素
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.head == -1 && this.tail == -1 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.head < this.tail && (this.tail-this.head) == len(this.queue)-1 {
		return true
	}
	if this.head > this.tail && this.head-this.tail == 1 {
		fmt.Printf(">. %v %v", this.head, this.tail)
		return true
	}
	return false
}

// 输出queue
func (this *MyCircularQueue) PrintQueue() {
	for _, v := range this.queue {
		fmt.Println(v)
	}
	fmt.Println(len(this.queue), cap(this.queue))
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
