package biomeconfig


// Experimental.
type LineEnding string

const (
	// Line Feed only (\n), common on Linux and macOS as well as inside git repos (lf).
	// Experimental.
	LineEnding_LF LineEnding = "LF"
	// Carriage Return + Line Feed characters (\r\n), common on Windows (crlf).
	// Experimental.
	LineEnding_CRLF LineEnding = "CRLF"
	// Carriage Return character only (\r), used very rarely (cr).
	// Experimental.
	LineEnding_CR LineEnding = "CR"
	// Automatically use CRLF on Windows and LF on other platforms (auto).
	// Experimental.
	LineEnding_AUTO LineEnding = "AUTO"
)

