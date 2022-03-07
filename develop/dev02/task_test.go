package main

import (
	"testing"
)

type TestCase struct {
	arg      string
	expected string
}

func TestFind(t *testing.T) {
	cases := []TestCase{
		{
			arg:      "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			arg:      "abcd",
			expected: "abcd",
		},
		{
			arg:      "45",
			expected: "",
		},
		{
			arg:      "14",
			expected: "",
		},
		{
			arg:      "a1",
			expected: "a",
		},
		{
			arg:      "b2a",
			expected: "bba",
		},
	}

	for _, tc := range cases {
		got := unpackString(tc.arg)
		if got != tc.expected {
			t.Errorf("unpack(%s): Expected %s, got %s", tc.arg, tc.expected, got)
		}
	}
}
