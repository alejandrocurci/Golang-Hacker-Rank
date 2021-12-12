package string_operations

import (
	"fmt"
	"testing"
)

type test struct {
	input    string
	expected string
}

func TestModifyString(t *testing.T) {
	cases := []test{
		{"de75s1rev2er", "reversed"},
		{"oll123eH56", "Hello"},
		{"!y7aw5-91o320n", "no-way!"},
	}

	for _, tt := range cases {
		name := fmt.Sprintf("Case %s", tt.expected)
		t.Run(name, func(t *testing.T) {
			actual := ModifyString(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %s, actual %s", tt.expected, actual)
			}
		})
	}
}
