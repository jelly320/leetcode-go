package leetcode

import (
	"go-leetcode/structures"
	"testing"
)

func Test_minDepth(t *testing.T) {
	//type args struct {
	//	root *TreeNode
	//}
	tests := []struct {
		name string
		args []int
		want int
	}{
		// TODO: Add test cases.
		{name: "one", args: []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}, want: 2},
		{name: "two", args: []int{2, structures.NULL, 3, structures.NULL, 4, structures.NULL, 5, structures.NULL, 6}, want: 5},
		{name: "1", args: []int{1}, want: 1},
		{name: "0", args: []int{}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structures.Ints2TreeNode(tt.args)
			if got := minDepth(root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
