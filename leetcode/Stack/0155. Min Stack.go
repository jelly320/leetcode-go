package Stack

// https://leetcode-cn.com/problems/min-stack/

type MinStack struct {
	stack [][]int
	top   int
	//bottom  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		[][]int{},
		0,
	}
}

// 将元素 x 推入栈中. 顺便判断大小,记录最小值
func (this *MinStack) Push(x int) {
	var min int
	if this.top == 0 {
		min = x
	} else {
		//fmt.Printf("getmin():%v\n",this.GetMin())
		if x < this.GetMin() {
			//fmt.Printf("tag <,min: %v; \n",min)
			min = x
		} else {
			min = this.GetMin()
		}
	}
	//fmt.Printf("min: %v; \n",min)
	this.stack = append(this.stack, []int{x, min})
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
		return this.stack[this.top-1][0]
	}
	return -1
}

// 检索栈中的最小元素
func (this *MinStack) GetMin() int {
	var result int
	if this.top > 0 {
		return this.stack[this.top-1][1]
	} else {
		return result
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
