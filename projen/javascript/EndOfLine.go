package javascript


// Experimental.
type EndOfLine string

const (
	// Maintain existing (mixed values within one file are normalised by looking at what's used after the first line).
	// Experimental.
	EndOfLine_AUTO EndOfLine = "AUTO"
	// Carriage Return character only (\r), used very rarely.
	// Experimental.
	EndOfLine_CR EndOfLine = "CR"
	// Carriage Return + Line Feed characters (\r\n), common on Windows.
	// Experimental.
	EndOfLine_CRLF EndOfLine = "CRLF"
	// Line Feed only (\n), common on Linux and macOS as well as inside git repos.
	// Experimental.
	EndOfLine_LF EndOfLine = "LF"
)

