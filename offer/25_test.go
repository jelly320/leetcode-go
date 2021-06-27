package offer

import (
	"go-leetcode/structures"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "test0", args: args{l1: []int{1, 2, 4}, l2: []int{1, 3, 4}}, want: []int{1, 1, 2, 3, 4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1 := structures.CreateOneListNode(tt.args.l1)
			p2 := structures.CreateOneListNode(tt.args.l2)
			if got := mergeTwoLists(p1, p2); !reflect.DeepEqual(got, structures.CreateOneListNode(tt.want)) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
