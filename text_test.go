package consolestyle_test

import (
	"testing"

	"github.com/ortense/consolestyle"
)

func TestText(t *testing.T) {
	run(t, map[string]testcase{
		"consolestyle.Bold": {
			input:    "text with style",
			act:      consolestyle.Bold,
			expected: "\x1b[1mtext with style\x1b[22m",
		},
		"consolestyle.Dim": {
			input:    "text with style",
			act:      consolestyle.Dim,
			expected: "\x1b[2mtext with style\x1b[22m",
		},
		"consolestyle.Italic": {
			input:    "text with style",
			act:      consolestyle.Italic,
			expected: "\x1b[3mtext with style\x1b[23m",
		},
		"consolestyle.Underline": {
			input:    "text with style",
			act:      consolestyle.Underline,
			expected: "\x1b[4mtext with style\x1b[24m",
		},
		"consolestyle.Strike": {
			input:    "text with style",
			act:      consolestyle.Strike,
			expected: "\x1b[9mtext with style\x1b[29m",
		},
	})
}
