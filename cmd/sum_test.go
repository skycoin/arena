package main

import (
	"log"
	"testing"
)

func Test_Sum(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    float64
	}{
		{
			name: "success-float",
			args: args{
				a: "12.8900",
				b: "8",
			},
			want: 20.890000343322754,
		},
		{
			name: "success-int",
			args: args{
				a: "121",
				b: "3",
			},
			want: 124,
		},
		{
			name: "success-negative-number",
			args: args{
				a: "-22.8900",
				b: "8.90",
			},
			want: -13.989999771118164,
		},
		{
			name: "invalid-arg",
			args: args{
				a: "11",
				b: "ten",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sum(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				log.Println(got)
				t.Errorf("sum() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
