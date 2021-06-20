package offer

import "testing"

func Test_findNumberIn2DArray(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test0",args: args{matrix: [][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		},target: 5},want: true},

		{name: "test1",args: args{matrix: [][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		},target: 20},want: false},

		{name: "test2",args: args{matrix: [][]int{},target: 0},want: false},
		{name: "test3",args: args{matrix: [][]int{{}},target: 1},want: false},
		{name: "test4",args: args{matrix: [][]int{{-5}},target: -5},want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArray(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findNumberIn2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
