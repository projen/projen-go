package biomeconfig


// Controls attribute placement in HTML-like languages.
// Experimental.
type AttributePosition string

const (
	// Places attributes automatically and wraps them when needed.
	//
	// (auto).
	// Experimental.
	AttributePosition_AUTO AttributePosition = "AUTO"
	// Places each attribute on its own line when an element has multiple attributes.
	//
	// (multiline).
	// Experimental.
	AttributePosition_MULTILINE AttributePosition = "MULTILINE"
)

