package main

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "With negative int",
			args: args{
				a: -3,
				b: -4,
			},
			want: -7,
		},
		{
			name: "With float",
			args: args{
				a: 1.2,
				b: 2,
			},
			want: 3.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
