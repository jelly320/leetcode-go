package Link

import (
	"go-leetcode/structures"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	aa := structures.CreateOneListNode([]int{1, 2, 3, 4, 5})
	cur := aa

	cur = cur.Next
	t.Logf("cur.Val: %v", cur.Val)
	t.Logf("aa.Val:%v", aa.Val)

}

//func Test_oddEvenList(t *testing.T) {
//	type args struct {
//		head *ListNode
//	}
//	tests := []struct {
//		name string
//		args args
//		want *ListNode
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
