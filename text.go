// Package consolestyle provides functions to style text printed to the console.
package consolestyle

const (
	boldCode           = "\x1b[1m"
	dimCode            = "\x1b[2m"
	italicCode         = "\x1b[3m"
	underlineCode      = "\x1b[4m"
	strikeCode         = "\x1b[9m"
	boldResetCode      = "\x1b[22m"
	italicResetCode    = "\x1b[23m"
	underlineResetCode = "\x1b[24m"
	strikeResetCode    = "\x1b[29m"
)

// Bold styles the text with bold font.
func Bold(s string) string {
	return boldCode + s + boldResetCode
}

// Dim styles the text with dim font.
func Dim(s string) string {
	return dimCode + s + boldResetCode
}

// Italic styles the text with italic font.
func Italic(s string) string {
	return italicCode + s + italicResetCode
}

// Underline underlines the text.
func Underline(s string) string {
	return underlineCode + s + underlineResetCode
}

// Strike strikes through the text.
func Strike(s string) string {
	return strikeCode + s + strikeResetCode
}
