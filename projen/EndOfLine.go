package projen


// The end of line characters supported by git.
// Experimental.
type EndOfLine string

const (
	// Maintain existing (mixed values within one file are normalised by looking at what's used after the first line).
	// Experimental.
	EndOfLine_AUTO EndOfLine = "AUTO"
	// Carriage Return + Line Feed characters (\r\n), common on Windows.
	// Experimental.
	EndOfLine_CRLF EndOfLine = "CRLF"
	// Line Feed only (\n), common on Linux and macOS as well as inside git repos.
	// Experimental.
	EndOfLine_LF EndOfLine = "LF"
	// Disable and do not configure the end of line character.
	// Experimental.
	EndOfLine_NONE EndOfLine = "NONE"
)

