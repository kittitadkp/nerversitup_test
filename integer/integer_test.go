package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOddOccurrence(t *testing.T) {
	type testCase struct {
		name     string
		input    []int
		expected int
	}

	testCases := []testCase{
		{
			name:     "should return -1",
			input:    []int{8, 8},
			expected: -1,
		},
		{
			name:     "should return 7",
			input:    []int{7},
			expected: 7,
		},
		{
			name:     "should return 0",
			input:    []int{0},
			expected: 0,
		},
		{
			name:     "should return 2",
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "should return 0",
			input:    []int{0, 1, 0, 1, 0},
			expected: 0,
		},
		{
			name:     "should return 4",
			input:    []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := FindOddOccurrence(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}
