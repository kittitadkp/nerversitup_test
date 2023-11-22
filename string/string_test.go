package permutation

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Should return empty string",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Should return ['apple', 'orange', 'banana', 'grape']",
			input:    []string{"apple", "orange", "banana", "apple", "grape", "banana"},
			expected: []string{"apple", "orange", "banana", "grape"},
		},
		{
			name:     "Should return ['a', 'b', 'c', 'd']",
			input:    []string{"a", "b", "a", "c", "b", "d", "a"},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "Should return ['red', 'blue', 'green']",
			input:    []string{"red", "blue", "green", "red", "blue", "green"},
			expected: []string{"red", "blue", "green"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeDuplicates(tc.input)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}

func TestGeneratePermutations(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected []string
	}

	testCases := []testCase{
		{
			name:  "Should return slice is nil",
			input: "",
		},
		{
			name:  "Should return slice is nil",
			input: " ",
		},
		{
			name:     "Should return ['a']",
			input:    "a",
			expected: []string{"a"},
		},
		{
			name:     "Should return ['ab', 'ba']",
			input:    "ab",
			expected: []string{"ab", "ba"},
		},
		{
			name:     "Should return ['abc', 'acb', 'bac', 'bca', 'cab', 'cba']",
			input:    "abc",
			expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name:     "Should return ['aabb', 'abab', 'abba', 'baab', 'baba', 'bbaa']",
			input:    "aabb",
			expected: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GeneratePermutations(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}

func TestCountSmileys(t *testing.T) {
	type testCase struct {
		name     string
		input    []string
		expected int
	}

	testCases := []testCase{
		{
			name:     "Should return 2",
			input:    []string{":)", ";(", ";}", ":-D"},
			expected: 2,
		},
		{
			name:     "Should return 3",
			input:    []string{";D", ":-(", ":-)", ";~)"},
			expected: 3,
		},
		{
			name:     "Should return 1",
			input:    []string{";]", ":[", ";*", ":$", ";-D"},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CountSmileys(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}
