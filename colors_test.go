package consolestyle_test

import (
	"testing"

	"github.com/ortense/consolestyle"
)

func TestColors(t *testing.T) {
	run(t, map[string]testcase{
		"consolestyle.Red": {
			input:    "text with color",
			act:      consolestyle.Red,
			expected: "\x1b[31mtext with color\x1b[39m",
		},
		"consolestyle.Green": {
			input:    "text with color",
			act:      consolestyle.Green,
			expected: "\x1b[32mtext with color\x1b[39m",
		},
		"consolestyle.Yellow": {
			input:    "text with color",
			act:      consolestyle.Yellow,
			expected: "\x1b[33mtext with color\x1b[39m",
		},
		"consolestyle.Blue": {
			input:    "text with color",
			act:      consolestyle.Blue,
			expected: "\x1b[34mtext with color\x1b[39m",
		},
		"consolestyle.Magenta": {
			input:    "text with color",
			act:      consolestyle.Magenta,
			expected: "\x1b[35mtext with color\x1b[39m",
		},
		"consolestyle.Cyan": {
			input:    "text with color",
			act:      consolestyle.Cyan,
			expected: "\x1b[36mtext with color\x1b[39m",
		},
		"consolestyle.BgRed": {
			input:    "text with color",
			act:      consolestyle.BgRed,
			expected: "\x1b[41mtext with color\x1b[49m",
		},
		"consolestyle.BgGree": {
			input:    "text with color",
			act:      consolestyle.BgGreen,
			expected: "\x1b[42mtext with color\x1b[49m",
		},
		"colors.bgYellow": {
			input:    "text with color",
			act:      consolestyle.BgYellow,
			expected: "\x1b[43mtext with color\x1b[49m",
		},
		"consolestyle.BgBlue,": {
			input:    "text with color",
			act:      consolestyle.BgBlue,
			expected: "\x1b[44mtext with color\x1b[49m",
		},
		"consolestyle.BgMagenta": {
			input:    "text with color",
			act:      consolestyle.BgMagenta,
			expected: "\x1b[45mtext with color\x1b[49m",
		},
		"consolestyle.BgCyan": {
			input:    "text with color",
			act:      consolestyle.BgCyan,
			expected: "\x1b[46mtext with color\x1b[49m",
		},
		"consolestyle.Inverse": {
			input:    "text with color",
			act:      consolestyle.Inverse,
			expected: "\x1b[7mtext with color\x1b[27m",
		},
	})
}
