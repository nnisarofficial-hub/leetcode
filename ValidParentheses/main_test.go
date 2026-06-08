package main

import "testing"

func TestValidParenthesis(t *testing.T) {
	tt := []struct {
		s        string
		expected bool
	}{
		{
			s: "()",
			expected: true,
		},
		{
			s: "()[]{}",
			expected: true
		},
		{
			s: "(]",
			expected: false ,
		},
		{
			s: "([])",
			expected: true,
		},

		{
			s: "([)]",
			expected: false,
		},	
	}
	for _, tc := range tt {
		actual := validParenthesis(tc.s)
		assert.Equal(t, actual, tc.expected)
	}

}
