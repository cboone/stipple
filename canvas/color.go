package canvas

// Color represents an ANSI foreground color.
type Color uint8

// Standard ANSI foreground colors (grouped, with ColorDefault at 0).
const (
	ColorDefault Color = iota
	ColorBlack
	ColorBlue
	ColorCyan
	ColorGreen
	ColorMagenta
	ColorRed
	ColorWhite
	ColorYellow
)

// ansiCodes maps Color values to ANSI escape sequences.
var ansiCodes = [...]string{
	ColorDefault: "",
	ColorBlack:   "\x1b[30m",
	ColorBlue:    "\x1b[34m",
	ColorCyan:    "\x1b[36m",
	ColorGreen:   "\x1b[32m",
	ColorMagenta: "\x1b[35m",
	ColorRed:     "\x1b[31m",
	ColorWhite:   "\x1b[37m",
	ColorYellow:  "\x1b[33m",
}

// ANSI returns the ANSI escape sequence for this color.
func (color Color) ANSI() string {
	if int(color) >= len(ansiCodes) {
		return ""
	}
	return ansiCodes[color]
}

// ANSIReset returns the ANSI reset escape sequence.
func ANSIReset() string {
	return "\x1b[0m"
}
