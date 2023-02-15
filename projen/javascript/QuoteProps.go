package javascript


// Experimental.
type QuoteProps string

const (
	// Only add quotes around object properties where required.
	// Experimental.
	QuoteProps_ASNEEDED QuoteProps = "ASNEEDED"
	// If at least one property in an object requires quotes, quote all properties.
	// Experimental.
	QuoteProps_CONSISTENT QuoteProps = "CONSISTENT"
	// Respect the input use of quotes in object properties.
	// Experimental.
	QuoteProps_PRESERVE QuoteProps = "PRESERVE"
)

