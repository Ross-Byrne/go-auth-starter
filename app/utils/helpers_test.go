package utils

import (
	"fmt"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input  string
		expect bool
	}{
		{"example@example.com", true},
		{"test@test.com", true},
		{"e@e", true},
		{"exam ple@example.com", false},
		{"@example.com", false},
		{"e@e.", false},
		{"jksdhfgklshdfgsdfg", false},
		{"@", false},
		{"test.com", false},
		{"@test.com", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Testing validity of %s", test.input), func(t *testing.T) {
			t.Parallel()

			actual := IsValidEmail(test.input)
			if actual != test.expect {
				t.Errorf("Incorrect Result, got: %v, expected: %v.", actual, test.expect)
			}
		})
	}
}
