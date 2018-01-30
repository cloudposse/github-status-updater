package main

import (
	"testing"
)

var validStateTCs = []struct {
	state    string
	expected bool
}{
	{"error", true},
	{"failure", true},
	{"pending", true},
	{"success", true},
	{"foobar", false},
}

func TestValidState(t *testing.T) {
	for _, tc := range validStateTCs {
		result := validState(tc.state)
		if result != tc.expected {
			t.Errorf("Error: expected '%t' to be\n'%t'", tc.expected, result)
		}
	}
}
