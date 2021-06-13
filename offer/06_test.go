package offer

import (
	"reflect"
	"testing"
)

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

/*tag:构造单链表*/
/*tag:构造ListNode*/
func InserListNode(ints []int) *ListNode {
	if len(ints) == 0 {
		return nil
	}
	value := &ListNode{
		Val: ints[0],
	}

	temp := value

	for i := 1; i < len(ints); i++ {
		temp.Next = &ListNode{Val: ints[i]}
		temp = temp.Next
	}
	return value

}
func Test_reversePrint(t *testing.T) {
	type args struct {
		//head *ListNode
		arg []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "test0", args: args{arg: []int{1, 3, 2}}, want: []int{2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aa := InserListNode(tt.args.arg)
			if got := reversePrint1(aa); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
