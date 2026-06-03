package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	tt := []struct {
		strs     []string
		expected string
	}{
		{
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
	}

	for _, tc := range tt {
		actual := longestCommonPrefix(tc.strs)
		assert.Equal(t, actual, tc.expected)
	}
}
