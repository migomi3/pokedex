package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		Expected string
	}{
		{"test", "test"},
		{"Test", "test"},
		{"TEST", "test"},
		{" test", "test"},
		{" test ", "test"},
		{"Test ", "test"},
		{"	test	", "test"},
		{"	test check", "test"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Running Input: %s\n", tc.input), func(t *testing.T) {
			result := cleanInput(tc.input)
			if tc.Expected != result {
				t.Errorf("Expected output [%s] does not match Actual output [%s]", tc.Expected, result)
			}
		})

	}

}
