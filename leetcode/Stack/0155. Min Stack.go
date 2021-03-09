package Stack

// https://leetcode-cn.com/problems/min-stack/

type MinStack struct {
	stack []int
	top   int
	min   []int
	//bottom  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		[]int{},
		0,
		0,
	}
}

// 将元素 x 推入栈中. 顺便判断大小,记录最小值
func (this *MinStack) Push(x int) {
	if this.top == 0 {
		this.min = x
	} else {
		if x < this.min {
			this.min = x
		}
	}
	this.stack = append(this.stack, x)
	this.top++
}

// 删除栈顶的元素. 刪除的时候也要处理this.min的值
func (this *MinStack) Pop() {
	if this.top != 0 {
		this.top--
	}
}

// 获取栈顶元素
func (this *MinStack) Top() int {
	if this.top != 0 {
		return this.stack[this.top-1]
	}
	return -1
}

// 检索栈中的最小元素
func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
