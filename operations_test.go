package arena

import (
	"testing"
)

func TestExecInt(t *testing.T) {
	type args []IntOp
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"base", args{Int(4), Int(6)}, 10, false},
		{"1 arg", args{Int(4)}, 0, true},
		{"wrong arg", args{StrInt("4"), StrInt("a")}, 0, true},
	}
	for _, tt := range tests {
		got, err := ExecInt(Add2(tt.args...))
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. ExecInt() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. ExecInt() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
