package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "  hello world  ",
			output: []string{"hello", "world"},
		}, {
			input:  "  foo bar baz  ",
			output: []string{"foo", "bar", "baz"},
		}, {
			input:  "  a b c d e  ",
			output: []string{"a", "b", "c", "d", "e"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.output) {
			t.Errorf("length of output slice is not equal to length of expected slice")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.output[i]
			if word != expectedWord {
				t.Errorf("output %v does not match expected %v", word, expectedWord)
			}
		}
	}
}
