package offer

import (
	"go-leetcode/structures"
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		args []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "test0", args: args{args: []int{1, 2, 3, 4, 5}, k: 2}, want: []int{4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			temp := structures.CreateOneListNode(tt.args.args)
			if got := getKthFromEnd(temp, tt.args.k); !reflect.DeepEqual(got, structures.CreateOneListNode(tt.want)) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
