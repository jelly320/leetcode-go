package Queue

import (
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
	t.Logf("param_1:%v", param_1)
	param_2 := obj.Rear()
	t.Logf("param_2:%v", param_2)
	param_3 := obj.EnQueue(9)
	t.Logf("param_3:%v", param_3)
	param_4 := obj.DeQueue()
	t.Logf("param_4:%v", param_4)
	param_5 := obj.Front()
	t.Logf("param_5:%v", param_5)
	param_6 := obj.DeQueue()
	t.Logf("param_6:%v", param_6)
	param_7 := obj.DeQueue() //true - false
	t.Logf("param_7:%v", param_7)
	param_8 := obj.IsEmpty()
	param_9 := obj.DeQueue()
	param_10 := obj.EnQueue(6)
	param_11 := obj.EnQueue(4)
	t.Log(param_1, param_2, param_3, param_4, param_5, param_6, param_7, param_8, param_9, param_10, param_11)
	//[null,true,4,true,true,9,true,false,true,false,true,true]
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
