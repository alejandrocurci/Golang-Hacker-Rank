package string_sorting

import (
	"fmt"
	"testing"
)

type test struct {
	input    []string
	expected []string
}

func TestCustomSorting(t *testing.T) {
	cases := []test{
		{
			input:    []string{"abc", "ab", "abcde", "a", "abcd"},
			expected: []string{"a", "abc", "abcde", "abcd", "ab"},
		},
		{
			input:    []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"},
			expected: []string{"Earth", "Venus", "Jupiter", "Mercury", "Neptune", "Saturn", "Uranus", "Mars"},
		},
	}
	for i, tt := range cases {
		t.Run(fmt.Sprintf("Test: %v", i+1), func(t *testing.T) {
			actual := CustomSorting(tt.input)
			if len(actual) != len(tt.expected) {
				t.Errorf("expected %v elements, returned %v", len(tt.expected), len(actual))
			}
			for j := range actual {
				if actual[j] != tt.expected[j] {
					t.Errorf("expected %s, got %s", tt.expected[j], actual[j])
				}
			}
		})
	}
}
