// Package consolestyle provides functions to style text printed to the console.
package consolestyle

import "strings"

// styled represents a styled string with a fluent api.
type styled struct {
	lines []string
	value string
}

// Style creates a new styled fluent interface.
func Style(value string) *styled {
	var lines []string
	return &styled{lines, value}
}

// Red styles the string with red color.
func (s *styled) Red() *styled {
	s.value = Red(s.value)
	return s
}

// Green styles the string with green color.
func (s *styled) Green() *styled {
	s.value = Green(s.value)
	return s
}

// Yellow styles the string with yellow color.
func (s *styled) Yellow() *styled {
	s.value = Yellow(s.value)
	return s
}

// Blue styles the string with blue color.
func (s *styled) Blue() *styled {
	s.value = Blue(s.value)
	return s
}

// Magenta styles the string with magenta color.
func (s *styled) Magenta() *styled {
	s.value = Magenta(s.value)
	return s
}

// Cyan styles the string with cyan color.
func (s *styled) Cyan() *styled {
	s.value = Cyan(s.value)
	return s
}

// BgRed styles the string with red background.
func (s *styled) BgRed() *styled {
	s.value = BgRed(s.value)
	return s
}

// BgGreen styles the string with green background.
func (s *styled) BgGreen() *styled {
	s.value = BgGreen(s.value)
	return s
}

// BgYellow styles the string with yellow background.
func (s *styled) BgYellow() *styled {
	s.value = BgYellow(s.value)
	return s
}

// BgBlue styles the string with blue background.
func (s *styled) BgBlue() *styled {
	s.value = BgBlue(s.value)
	return s
}

// BgMagenta styles the string with magenta background.
func (s *styled) BgMagenta() *styled {
	s.value = BgMagenta(s.value)
	return s
}

// BgCyan styles the string with cyan background.
func (s *styled) BgCyan() *styled {
	s.value = BgCyan(s.value)
	return s
}

// Inverse inverts the string.
func (s *styled) Inverse() *styled {
	s.value = Inverse(s.value)
	return s
}

// Bold makes the string bold.
func (s *styled) Bold() *styled {
	s.value = Bold(s.value)
	return s
}

// Dim makes the string dim.
func (s *styled) Dim() *styled {
	s.value = Dim(s.value)
	return s
}

// Italic makes the string italic.
func (s *styled) Italic() *styled {
	s.value = Italic(s.value)
	return s
}

// Underline underlines the string.
func (s *styled) Underline() *styled {
	s.value = Underline(s.value)
	return s
}

// Strike strikes through the string.
func (s *styled) Strike() *styled {
	s.value = Strike(s.value)
	return s
}

// NewLine adds a new line with the provided value to the styled string.
func (s *styled) NewLine(value string) *styled {
	s.lines = append(s.lines, s.value)
	s.value = value
	return s
}

// EmptyLine adds an empty line to the styled string.
func (s *styled) EmptyLine() *styled {
	s.lines = append(s.lines, s.value)
	s.value = ""
	return s
}

// String returns the styled string with formatting applied.
func (s *styled) String() string {
	return strings.Join(append(s.lines, s.value), "\n")
}
