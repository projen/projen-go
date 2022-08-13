package javascript


// Experimental.
type TrailingComma string

const (
	// Trailing commas wherever possible (including function arguments).
	// Experimental.
	TrailingComma_ALL TrailingComma = "ALL"
	// Trailing commas where valid in ES5 (objects, arrays, etc.).
	// Experimental.
	TrailingComma_ES5 TrailingComma = "ES5"
	// No trailing commas.
	// Experimental.
	TrailingComma_NONE TrailingComma = "NONE"
)

