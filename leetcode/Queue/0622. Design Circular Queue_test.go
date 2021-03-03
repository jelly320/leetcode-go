package Queue

import (
	"fmt"
	"testing"
)

func TestMyCircularQueue_IsFull(t *testing.T) {
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

	obj := Constructor(2)

	param_1 := obj.EnQueue(4)
	param_2 := obj.Rear()
	param_3 := obj.EnQueue(9)
	param_4 := obj.DeQueue()
	param_5 := obj.Front()
	param_6 := obj.DeQueue()
	param_7 := obj.DeQueue()
	param_8 := obj.IsEmpty()
	param_9 := obj.DeQueue()
	param_10 := obj.EnQueue(6)
	param_11 := obj.EnQueue(4)

	//for i:=1;i<=11;i++ {
	//	tag:="param_"+string(i)
	//	fmt.Printf("%v", tag)
	//}

	//obj := Constructor(3)
	//t.Log(obj.IsEmpty())
	//// 入队
	//t.Log(obj.EnQueue(1))
	//t.Log(obj.EnQueue(2))
	//t.Log(obj.EnQueue(3))
	//t.Log(obj.EnQueue(4))
	//t.Log(obj.IsFull())
	//// 出队
	//t.Log(obj.DeQueue())
	//// 入队
	//t.Log(obj.EnQueue(4))
	//
	//t.Log(obj.DeQueue())
	//t.Log(obj.DeQueue())
	//t.Log(obj.DeQueue())
	//t.Log(obj.DeQueue())
	//t.Log(obj.IsEmpty())
	//obj.PrintQueue()

}
