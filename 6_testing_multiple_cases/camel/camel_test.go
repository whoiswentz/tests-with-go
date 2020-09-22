package camel

import (
	"testing"
)

func TestCamel_simple(t *testing.T) {
	testCases := []struct {
		arg  string
		want string
	}{
		{arg: "thisIsACamelCaseString", want: "this_is_a_camel_case_string"},
	}

	for _, tt := range testCases {
		got := Camel(tt.arg)
		if got != tt.want {
			t.Errorf("Camel(%q) = %q; want = %q", tt.arg, got, tt.want)
		}
	}
}

func TestCamel(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "should pass",
			args: "thisIsACamelCaseString",
			want: "this_is_a_camel_case_string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel(tt.args); got != tt.want {
				t.Errorf("Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}
