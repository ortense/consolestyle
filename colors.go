// Package consolestyle provides functions to style text printed to the console.
package consolestyle

const (
	redCode          = "\x1b[31m"
	greenCode        = "\x1b[32m"
	yellowCode       = "\x1b[33m"
	blueCode         = "\x1b[34m"
	magentaCode      = "\x1b[35m"
	cyanCode         = "\x1b[36m"
	defaultColorCode = "\x1b[39m"
	bgRedCode        = "\x1b[41m"
	bgGreenCode      = "\x1b[42m"
	bgYellowCode     = "\x1b[43m"
	bgBlueCode       = "\x1b[44m"
	bgMagentaCode    = "\x1b[45m"
	bgCyanCode       = "\x1b[46m"
	bgDefaultColor   = "\x1b[49m"
	inverseCode      = "\x1b[7m"
	inverseResetCode = "\x1b[27m"
)

// Red styles the text with red color.
func Red(s string) string {
	return redCode + s + defaultColorCode
}

// Green styles the text with green color.
func Green(s string) string {
	return greenCode + s + defaultColorCode
}

// Yellow styles the text with yellow color.
func Yellow(s string) string {
	return yellowCode + s + defaultColorCode
}

// Blue styles the text with blue color.
func Blue(s string) string {
	return blueCode + s + defaultColorCode
}

// Magenta styles the text with magenta color.
func Magenta(s string) string {
	return magentaCode + s + defaultColorCode
}

// Cyan styles the text with cyan color.
func Cyan(s string) string {
	return cyanCode + s + defaultColorCode
}

// BgRed styles the text with red background.
func BgRed(s string) string {
	return bgRedCode + s + bgDefaultColor
}

// BgGreen styles the text with green background.
func BgGreen(s string) string {
	return bgGreenCode + s + bgDefaultColor
}

// BgYellow styles the text with yellow background.
func BgYellow(s string) string {
	return bgYellowCode + s + bgDefaultColor
}

// BgBlue styles the text with blue background.
func BgBlue(s string) string {
	return bgBlueCode + s + bgDefaultColor
}

// BgMagenta styles the text with magenta background.
func BgMagenta(s string) string {
	return bgMagentaCode + s + bgDefaultColor
}

// BgCyan styles the text with cyan background.
func BgCyan(s string) string {
	return bgCyanCode + s + bgDefaultColor
}

// Inverse inverts the text.
func Inverse(s string) string {
	return inverseCode + s + inverseResetCode
}
