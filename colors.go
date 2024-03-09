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

func Red(s string) string {
	return redCode + s + defaultColorCode
}

func Green(s string) string {
	return greenCode + s + defaultColorCode
}

func Yellow(s string) string {
	return yellowCode + s + defaultColorCode
}

func Blue(s string) string {
	return blueCode + s + defaultColorCode
}

func Magenta(s string) string {
	return magentaCode + s + defaultColorCode
}

func Cyan(s string) string {
	return cyanCode + s + defaultColorCode
}

func BgRed(s string) string {
	return bgRedCode + s + bgDefaultColor
}

func BgGreen(s string) string {
	return bgGreenCode + s + bgDefaultColor
}

func BgYellow(s string) string {
	return bgYellowCode + s + bgDefaultColor
}

func BgBlue(s string) string {
	return bgBlueCode + s + bgDefaultColor
}

func BgMagenta(s string) string {
	return bgMagentaCode + s + bgDefaultColor
}

func BgCyan(s string) string {
	return bgCyanCode + s + bgDefaultColor
}

func Inverse(s string) string {
	return inverseCode + s + inverseResetCode
}
