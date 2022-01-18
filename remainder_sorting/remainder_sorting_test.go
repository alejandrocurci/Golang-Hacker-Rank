package remainder_sorting

import (
	"fmt"
	"testing"
)

type test struct {
	input    []string
	expected []string
}

func TestRemainderSorting(t *testing.T) {
	cases := []test{
		{
			input:    []string{"Colorado", "Utah", "Wisconsin", "Oregon"},
			expected: []string{"Oregon", "Wisconsin", "Utah", "Colorado"},
		},
		{
			input:    []string{"ace", "ron", "fit", "zip", "leg"},
			expected: []string{"ace", "fit", "leg", "ron", "zip"},
		},
		{
			input:    []string{"jhon", "bob", "lionel", "smith"},
			expected: []string{"bob", "lionel", "jhon", "smith"},
		},
		{
			input:    []string{"alex", "cat", "diego"},
			expected: []string{"cat", "alex", "diego"},
		},
		{
			input:    []string{"a", "ab", "bc", "abc"},
			expected: []string{"abc", "a", "ab", "bc"},
		},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprintf("Test: %v", i+1), func(t *testing.T) {
			actual := RemainderSorting(tt.input)
			if len(actual) != len(tt.expected) {
				t.Errorf("expected %v elements, returned %v", len(tt.expected), len(actual))
			}
			for j, s := range actual {
				if s != tt.expected[j] {
					t.Errorf("expected %s, got %s", tt.expected[j], s)
				}
			}
		})
	}
}
