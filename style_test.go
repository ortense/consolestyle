package consolestyle_test

import (
	"testing"

	"github.com/ortense/consolestyle"
)

func TestStyle(t *testing.T) {
	run(t, map[string]testcase{
		"styled.Red": {
			input:    "text with style",
			expected: "\x1b[31mtext with style\x1b[39m",
			act: func(input string) string {
				return consolestyle.Style(input).Red().String()
			},
		},
		"styled.Green": {
			input:    "text with style",
			expected: "\x1b[32mtext with style\x1b[39m",
			act: func(input string) string {
				return consolestyle.Style(input).Green().String()
			},
		},
		"styled.Yellow": {
			input:    "text with style",
			expected: "\x1b[33mtext with style\x1b[39m",
			act: func(input string) string {
				return consolestyle.Style(input).Yellow().String()
			},
		},
		"styled.Blue": {
			input:    "text with style",
			expected: "\x1b[34mtext with style\x1b[39m",
			act: func(input string) string {
				return consolestyle.Style(input).Blue().String()
			},
		},
		"styled.Magenta": {
			input:    "text with style",
			expected: "\x1b[35mtext with style\x1b[39m",
			act: func(input string) string {
				return consolestyle.Style(input).Magenta().String()
			},
		},
		"styled.Cyan": {
			input:    "text with style",
			expected: "\x1b[36mtext with style\x1b[39m",
			act: func(input string) string {
				return consolestyle.Style(input).Cyan().String()
			},
		},
		"styled.BgRed": {
			input:    "text with style",
			expected: "\x1b[41mtext with style\x1b[49m",
			act: func(input string) string {
				return consolestyle.Style(input).BgRed().String()
			},
		},
		"styled.BgGreen": {
			input:    "text with style",
			expected: "\x1b[42mtext with style\x1b[49m",
			act: func(input string) string {
				return consolestyle.Style(input).BgGreen().String()
			},
		},
		"styled.BgYellow": {
			input:    "text with style",
			expected: "\x1b[43mtext with style\x1b[49m",
			act: func(input string) string {
				return consolestyle.Style(input).BgYellow().String()
			},
		},
		"styled.BgBlue": {
			input:    "text with style",
			expected: "\x1b[44mtext with style\x1b[49m",
			act: func(input string) string {
				return consolestyle.Style(input).BgBlue().String()
			},
		},
		"styled.BgMagenta": {
			input:    "text with style",
			expected: "\x1b[45mtext with style\x1b[49m",
			act: func(input string) string {
				return consolestyle.Style(input).BgMagenta().String()
			},
		},
		"styled.BgCyan": {
			input:    "text with style",
			expected: "\x1b[46mtext with style\x1b[49m",
			act: func(input string) string {
				return consolestyle.Style(input).BgCyan().String()
			},
		},
		"styled.Inverse": {
			input:    "text with style",
			expected: "\x1b[7mtext with style\x1b[27m",
			act: func(input string) string {
				return consolestyle.Style(input).Inverse().String()
			},
		},
		"styled.Bold": {
			input:    "text with style",
			expected: "\x1b[1mtext with style\x1b[22m",
			act: func(input string) string {
				return consolestyle.Style(input).Bold().String()
			},
		},
		"styled.Dim": {
			input:    "text with style",
			expected: "\x1b[2mtext with style\x1b[22m",
			act: func(input string) string {
				return consolestyle.Style(input).Dim().String()
			},
		},
		"styled.Italic": {
			input:    "text with style",
			expected: "\x1b[3mtext with style\x1b[23m",
			act: func(input string) string {
				return consolestyle.Style(input).Italic().String()
			},
		},
		"styled.Underline": {
			input:    "text with style",
			expected: "\x1b[4mtext with style\x1b[24m",
			act: func(input string) string {
				return consolestyle.Style(input).Underline().String()
			},
		},
		"styled.Strike": {
			input:    "text with style",
			expected: "\x1b[9mtext with style\x1b[29m",
			act: func(input string) string {
				return consolestyle.Style(input).Strike().String()
			},
		},
		"styled.NewLine": {
			input:    "text with style",
			expected: "text with style\nnew line",
			act: func(input string) string {
				return consolestyle.Style(input).NewLine("new line").String()
			},
		},
		"styled.EmptyLine": {
			input:    "text with style",
			expected: "text with style\n",
			act: func(input string) string {
				return consolestyle.Style(input).EmptyLine().String()
			},
		},
	})
}
