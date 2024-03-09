package consolestyle

import "strings"

type styled struct {
	lines []string
	value string
}

func Style(value string) *styled {
	var lines []string
	return &styled{lines, value}
}

func (s *styled) Red() *styled {
	s.value = Red(s.value)
	return s
}

func (s *styled) Green() *styled {
	s.value = Green(s.value)
	return s
}

func (s *styled) Yellow() *styled {
	s.value = Yellow(s.value)
	return s
}

func (s *styled) Blue() *styled {
	s.value = Blue(s.value)
	return s
}

func (s *styled) Magenta() *styled {
	s.value = Magenta(s.value)
	return s
}

func (s *styled) Cyan() *styled {
	s.value = Cyan(s.value)
	return s
}

func (s *styled) BgRed() *styled {
	s.value = BgRed(s.value)
	return s
}

func (s *styled) BgGreen() *styled {
	s.value = BgGreen(s.value)
	return s
}

func (s *styled) BgYellow() *styled {
	s.value = BgYellow(s.value)
	return s
}

func (s *styled) BgBlue() *styled {
	s.value = BgBlue(s.value)
	return s
}

func (s *styled) BgMagenta() *styled {
	s.value = BgMagenta(s.value)
	return s
}

func (s *styled) BgCyan() *styled {
	s.value = BgCyan(s.value)
	return s
}

func (s *styled) Inverse() *styled {
	s.value = Inverse(s.value)
	return s
}

func (s *styled) Bold() *styled {
	s.value = Bold(s.value)
	return s
}

func (s *styled) Dim() *styled {
	s.value = Dim(s.value)
	return s
}

func (s *styled) Italic() *styled {
	s.value = Italic(s.value)
	return s
}

func (s *styled) Underline() *styled {
	s.value = Underline(s.value)
	return s
}

func (s *styled) Strike() *styled {
	s.value = Strike(s.value)
	return s
}

func (s *styled) NewLine(value string) *styled {
	s.lines = append(s.lines, s.value)
	s.value = value
	return s
}

func (s *styled) EmptyLine() *styled {
	s.lines = append(s.lines, s.value)
	s.value = ""
	return s
}

func (s *styled) String() string {
	return strings.Join(append(s.lines, s.value), "\n")
}
