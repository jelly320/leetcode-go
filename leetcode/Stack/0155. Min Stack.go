package Stack

// https://leetcode-cn.com/problems/min-stack/

/* v0
执行用时：20 ms, 在所有 Go 提交中击败了83.44%的用户
内存消耗：10.1 MB, 在所有 Go 提交中击败了7.07%的用户
*/
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
		this.stack = [][]int{}
	} else {
		if x < this.GetMin() {
			min = x
		} else {
			min = this.GetMin()
		}
	}
	//这里没办法使用赋值的方式去覆盖掉原有的值,只能在pop中调整队列,因为不是实际删除元素如果不调整Top就容易得出错误结果
	//this.stack[this.top]=[]int{x,min}  panic: runtime error: index out of range [0] with length 0
	this.stack = append(this.stack, []int{x, min})
	this.top++
}

// 删除栈顶的元素. 刪除的时候也要处理this.min的值
func (this *MinStack) Pop() {
	if this.top != 0 {
		this.top--
		this.stack = this.stack[:this.top]
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
	return this.stack[this.top-1][1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
