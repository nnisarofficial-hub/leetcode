package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		x        int
		expected bool
	}{
		{
			x:        121,
			expected: true,
		},
		{
			x:        -121,
			expected: false,
		},
		{
			x:        10,
			expected: false,
		},
		{
			x:        123454321,
			expected: true,
		},
	}

	for _, tc := range tt {
		actual := isPalindrome(tc.x)
		assert.Equal(t, actual, tc.expected)
	}
}
