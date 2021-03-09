package Stack

import (
	"fmt"
	"testing"
)

/*
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
*/

func TestStack(t *testing.T) {
	//obj1 := Constructor()
	//obj1.Push(1)
	//fmt.Printf("obj1 = %v\n", obj1)
	//obj1.Push(0)
	//fmt.Printf("obj1 = %v\n", obj1)
	//obj1.Push(10)
	//fmt.Printf("obj1 = %v\n", obj1)
	//obj1.Pop()
	//fmt.Printf("obj1 = %v\n", obj1)
	//param3 := obj1.Top()
	//fmt.Printf("param_3 = %v\n", param3)
	//param4 := obj1.GetMin()
	//fmt.Printf("param_4 = %v\n", param4)
	//
	//obj2 := Constructor()
	//obj2.Push(1)
	//fmt.Printf("obj2 = %v\n", obj2)
	//obj2.Push(0)
	//fmt.Printf("obj2 = %v\n", obj2)
	//obj2.Push(10)
	//fmt.Printf("obj2 = %v\n", obj2)
	//obj2.Pop()
	//fmt.Printf("obj2 = %v\n", obj2)
	//obj2.Pop()
	//fmt.Printf("obj2 = %v\n", obj2)
	//obj2.Pop()
	//fmt.Printf("obj2 = %v\n", obj2)
	//param5 := obj2.Top()
	//fmt.Printf("param_5 = %v\n", param5)
	//
	//obj3 := Constructor()
	//obj3.Push(-2)
	//fmt.Printf("obj3 = %v\n", obj3)
	//obj3.Push(0)
	//fmt.Printf("obj3 = %v\n", obj3)
	//obj3.Push(-3)
	//fmt.Printf("obj3 = %v\n", obj3)
	//param8 := obj3.GetMin()
	//fmt.Printf("param8 = %v\n", param8)
	//obj3.Pop()
	//fmt.Printf("obj3 = %v\n", obj3)
	//param6 := obj3.Top()
	//fmt.Printf("param6 = %v\n", param6)
	//param7 := obj3.GetMin()
	//fmt.Printf("param7 = %v\n", param7)
	//["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
	//[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
	obj5 := Constructor()
	fmt.Printf("obj5 = %v\n", obj5)
	fmt.Println(obj5.GetMin())

	obj5.Push(2147483646)
	obj5.Push(2147483646)
	obj5.Push(2147483647)
	fmt.Println(obj5.Top())
	obj5.Pop()
	fmt.Println(obj5.GetMin())
	obj5.Pop()
	fmt.Println(obj5.GetMin())
	obj5.Pop()
	obj5.Push(2147483647)
	fmt.Println(obj5.Top())
	fmt.Println(obj5.GetMin())
	obj5.Push(-2147483648)
	fmt.Println(obj5.Top())
	fmt.Println(obj5.GetMin())
	obj5.Pop()
	fmt.Println(obj5.GetMin())
}
