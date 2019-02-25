package services

import "testing"

func Test_serviceInput_SumsUp(t *testing.T) {
	type fields struct {
		input []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{input: []int{1, 2, 3}},
			want:    6,
			wantErr: false,
		},
		{
			name:    "fail-emptyInput",
			fields:  fields{input: []int{}},
			want:    6,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := serviceInput{
				input: tt.fields.input,
			}
			got, err := s.SumsUp()
			if (err != nil) != tt.wantErr {
				t.Errorf("serviceInput.SumsUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("serviceInput.SumsUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
