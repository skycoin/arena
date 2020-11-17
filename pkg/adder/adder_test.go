package adder

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{name: "invalid number1", args: []string{"invalid", "123"}, want: "", wantErr: true},
		{name: "invalid number2", args: []string{"123", "invalid"}, want: "", wantErr: true},
		{name: "valid numbers", args: []string{"123", "123.15"}, want: "246.15", wantErr: false},
	}
	for _, tv := range tests {
		tt := tv
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args)
			if tt.wantErr {
				assert.Error(t, err, tt.name+" error returned")
				assert.True(t, strings.HasPrefix(err.Error(), tt.name))
			} else {
				assert.NotNil(t, got, tt.name+" adder initialized")
				if got != nil {
					assert.Equal(t, tt.want, got.Add().String(), tt.name+" adder return valid result")
				}
			}
		})
	}
}
