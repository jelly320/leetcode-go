package offer

import "testing"

func Test_minArray(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test0", args: args{numbers: []int{3, 4, 5, 1, 2}}, want: 1},
		{name: "test1", args: args{numbers: []int{2, 2, 2, 0, 1}}, want: 0},
		{name: "test2", args: args{numbers: []int{1, 0, 1, 1, 1}}, want: 0},
		{name: "test3", args: args{numbers: []int{1, 1, 1, 0, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.args.numbers); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
