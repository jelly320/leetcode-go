package Link

import (
	"go-leetcode/structures"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		//head *ListNode
		head []int
	}
	tests := []struct {
		name string
		args args
		//want *ListNode
		want []int
	}{
		// TODO: Add test cases.
		{name: "test0", args: args{head: []int{1, 2, 3, 4, 5}}, want: []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := structures.CreateOneListNode(tt.args.head)
			y := structures.CreateOneListNode(tt.want)
			if got := reverseList(x); !reflect.DeepEqual(got, y) {
				t.Errorf("reverseList() = %v, want %v", got, y)
			}
		})
	}
}
