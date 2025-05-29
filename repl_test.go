package main

import (
	"bytes"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "YelloW WorLd",
			expected: []string{"yellow", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch, expected: %v got: %v", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected: %q got: %q", expectedWord, word)
			}
		}
	}
}

func TestCommandHelp(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input: "help",
			expected: `
Welcome to the Pokedex!
Usage:

exit: Exit the Pokedex
help: Displays the help message
`,
		},
	}

	buffer := &bytes.Buffer{}

	for _, c := range cases {
		displayHelp(buffer)
		actual := buffer.String()

		if actual != c.expected {
			t.Errorf("Expected: %v Got: %v", c.expected, actual)
		}
	}
}
