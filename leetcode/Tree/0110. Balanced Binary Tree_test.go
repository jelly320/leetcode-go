package leetcode

import (
	"go-leetcode/structures"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args []int
		want bool
	}{
		// TODO: Add test cases.
		{name: "a", args: []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}, want: true},
		{name: "b", args: []int{1, 2, 2, 3, 3, structures.NULL, structures.NULL, 4, 4}, want: false},
		{name: "c", args: []int{}, want: true},
		{name: "d", args: []int{1}, want: true},
		{name: "error", args: []int{2, 1, 3}, want: true},
		{name: "error", args: []int{1, 2, 2, 3, structures.NULL, structures.NULL, 3, 4, structures.NULL, structures.NULL, 4}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structures.Ints2TreeNode(tt.args)
			if got := isBalanced(root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
