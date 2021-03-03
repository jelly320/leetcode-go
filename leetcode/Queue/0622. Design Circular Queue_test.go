package Queue

import "testing"

func TestMyCircularQueue_IsFull(t *testing.T) {
	obj := Constructor(3)
	t.Log(obj.EnQueue(1))
	t.Log(obj.EnQueue(2))
	t.Log(obj.EnQueue(3))
	t.Log(obj.EnQueue(4))
	//param_5 := obj.IsEmpty()
	//param_6 := obj.IsFull()
	//t.Log(param_5,param_6)

}
