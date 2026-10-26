package biomeconfig


// Experimental.
type ArrowParentheses string

const (
	// Always prints parentheses around arrow-function parameters.
	//
	// (always).
	// Experimental.
	ArrowParentheses_ALWAYS ArrowParentheses = "ALWAYS"
	// Omits parentheses when an arrow function has one parameter and the syntax permits it.
	//
	// (asNeeded).
	// Experimental.
	ArrowParentheses_AS_NEEDED ArrowParentheses = "AS_NEEDED"
)

