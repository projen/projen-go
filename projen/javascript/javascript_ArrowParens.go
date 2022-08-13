package javascript


// Experimental.
type ArrowParens string

const (
	// Always include parens.
	//
	// Example: `(x) => x`.
	// Experimental.
	ArrowParens_ALWAYS ArrowParens = "ALWAYS"
	// Omit parens when possible.
	//
	// Example: `x => x`.
	// Experimental.
	ArrowParens_AVOID ArrowParens = "AVOID"
)

