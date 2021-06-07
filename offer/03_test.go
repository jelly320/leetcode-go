package offer

import "testing"

func Test_findRepeatNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name:"test",args:args{nums: []int{2, 3, 1, 0, 2, 5, 3}},want:2},
		{name:"test1",args:args{nums: []int{0,1,2,3}},want:-1},
		{name:"test2",args:args{nums: []int{0,0,1,1}},want:0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRepeatNumberv1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name:"test",args:args{nums: []int{2, 3, 1, 0, 2, 5, 3}},want:2},
		{name:"test1",args:args{nums: []int{0,1,2,3}},want:-1},
		{name:"test2",args:args{nums: []int{0,0,1,1}},want:0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumberv1(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}