package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	tt := []struct {
		roman    string
		expected int
	}{
		{
			roman:    "III",
			expected: 3,
		},
		{
			roman:    "LVIII",
			expected: 58,
		},
		{
			roman:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, tc := range tt {
		actual := romanToInt(tc.roman)
		assert.Equal(t, actual, tc.expected)
	}

}
