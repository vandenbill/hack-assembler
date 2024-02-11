package parser_test

import (
	"testing"

	"github.com/vandenbill/hack-assembler/parser"
)

func TestDeleteWhiteSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},                                 // Test empty string
		{"   ", ""},                              // Test string with only whitespaces
		{"Hello World", "HelloWorld"},            // Test string without comments or whitespaces
		{"// Comment line", ""},                  // Test string with comment line
		{"//Comment line without space", ""},     // Test string with comment line without space
		{"// ", ""},                              // Test string with only comment
		{"   Some text  // Comment", "Sometext"}, // Test string with text and comment
		{"D; JEQ // jump if the different between i and b is 0", "D;JEQ"}, // Test actual code
	}

	for _, test := range tests {
		result := parser.DeleteWhiteSpace(test.input)
		if result != test.expected {
			t.Errorf("input '%s', expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestToInstructions(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"@0", "0000000000000000",
		}, {
			"D=M", "1111110000010000",
		}, {
			"@16", "0000000000001000",
		}, {
			"M=D", "1110001100001000",
		},
	}

	for _, test := range tests {
		result := parser.ToInstructions([]string{test.input})

		if len(test.expected) != len(result[0]) {
			t.Errorf("input '%s', expected '%s', different length", test.input, test.expected)
		}

		for i := range test.input {
			if test.expected[i] != result[0][i] {
				t.Errorf("input '%s', expected '%s', got '%s'", test.input, test.expected, result)
			}
		}
	}
}
