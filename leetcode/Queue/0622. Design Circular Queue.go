package Queue

import "fmt"

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

	// 先判断再自增有问题
	//if ! this.IsFull() {
	//	this.queue[this.tail] = value
	//	if this.tail==len(this.queue)-1 {
	//		this.tail=0
	//	} else {
	//		this.tail++
	//	}
	//	return true
	//} else {
	//	return false
	//}

	// 先移动再赋值
	fmt.Printf("%v,%v\n", this.head, this.tail)
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
	return false
}

// 获取队首元素
func (this *MyCircularQueue) Front() int {
	return 1
}

// 获取队尾元素
func (this *MyCircularQueue) Rear() int {
	return 1
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.head == -1 && this.tail == -1 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	//fmt.Printf("start. %v %v\n", this.head, this.tail)
	if this.head < this.tail && (this.tail-this.head) == len(this.queue)-1 {
		//fmt.Printf("<. %v %v\n",this.head,this.tail)
		return true
	}
	if this.head > this.tail && this.head-this.tail == 1 {
		fmt.Printf(">. %v %v", this.head, this.tail)
		return true
	}
	return false
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
