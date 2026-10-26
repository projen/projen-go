package biomeconfig


// Experimental.
type QuoteProperties string

const (
	// Removes quotes when the property name is valid without them.
	//
	// (asNeeded).
	// Experimental.
	QuoteProperties_AS_NEEDED QuoteProperties = "AS_NEEDED"
	// Preserves quotes around property names that were quoted in the input.
	//
	// (preserve).
	// Experimental.
	QuoteProperties_PRESERVE QuoteProperties = "PRESERVE"
)

