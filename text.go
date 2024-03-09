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

func Bold(s string) string {
	return boldCode + s + boldResetCode
}

func Dim(s string) string {
	return dimCode + s + boldResetCode
}

func Italic(s string) string {
	return italicCode + s + italicResetCode
}

func Underline(s string) string {
	return underlineCode + s + underlineResetCode
}

func Strike(s string) string {
	return strikeCode + s + strikeResetCode
}
