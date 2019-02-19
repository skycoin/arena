package testdata

import "testing"

var flagtests = []struct {
	in  string
	out string
}{
	{"a", "aa"},
	{"b", "ba"},
	{"c", "ca"},
	{"d", "da"},
	{"e", "ea"},
	{"f", "fa"},
	{"g", "ga"},
	{"h", "ha"},
	{"e", "ba"},
	{"k", "ka"},
	{"l", "lt"},
	{"m", "ma"},
}

func TestFlagParser(t *testing.T) {
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			s := tt.in + "a"
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
