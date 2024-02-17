package javascript


// Charset for esbuild's output.
// Experimental.
type Charset string

const (
	// ASCII.
	//
	// Any non-ASCII characters are escaped using backslash escape sequences.
	// Experimental.
	Charset_ASCII Charset = "ASCII"
	// UTF-8.
	//
	// Keep original characters without using escape sequences.
	// Experimental.
	Charset_UTF8 Charset = "UTF8"
)

