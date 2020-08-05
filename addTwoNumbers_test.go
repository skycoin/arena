package main

import "testing"

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Test1",args{2,4},6},
		{"Test2",args{7,4},11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}