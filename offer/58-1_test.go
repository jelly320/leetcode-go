package offer

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "test0",args: args{s: "I am a student.    "},want: "student. a am I"},
		{name: "test1",args: args{s: "the sky is blue"},want: "blue is sky the"},
		{name: "test2",args: args{s: "  hello world!  "},want: "world! hello"},
		{name: "test3",args: args{s: " h"},want: "h"},
		{name: "test4",args: args{s: "1"},want: "1"},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords2(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}