package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"test", []string{"test"}},
		{"Test", []string{"test"}},
		{"TEST", []string{"test"}},
		{" test", []string{"test"}},
		{" test ", []string{"test"}},
		{"Test ", []string{"test"}},
		{"	test	", []string{"test"}},
		{"	test check", []string{"test", "check"}},
		{"	test   check", []string{"test", "check"}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running Input: %s\n", tc.input), func(t *testing.T) {
			result := cleanInput(tc.input)

			if len(tc.expected) != len(result) {
				t.Errorf("lengths don't match: '%v' vs '%v'", result, tc.expected)
			}

			for i := range tc.expected {
				if tc.expected[i] != result[i] {
					t.Errorf("Expected output [%s] does not match Actual output [%s]", tc.expected, result)
				}
			}

		})

	}

}
