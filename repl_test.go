package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
		errMsg   string
	}{
		{
			input:    "SIN MAyuSCUlas",
			expected: []string{"sin", "mayusculas"},
			errMsg:   "all letters must be lower case. got: '%v' expect: '%v'",
		},
		{
			input:    "holi careboli",
			expected: []string{"holi", "careboli"},
			errMsg:   "must slice words in the ' '. got: '%v' expect: '%v'",
		},
		{
			input:    " holi care     boli  ",
			expected: []string{"holi", "care", "boli"},
			errMsg:   "must remove trailling and leading spaces. got: '%v' expect: '%v'",
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Error("the length expected is different from the actual")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf(c.errMsg, word, expectedWord)
			}
		}
	}
}
