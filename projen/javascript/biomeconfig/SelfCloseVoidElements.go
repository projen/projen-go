package biomeconfig


// Controls whether void-elements should be self closed.
// Experimental.
type SelfCloseVoidElements string

const (
	// The `/` inside void elements is removed by the formatter (never).
	// Experimental.
	SelfCloseVoidElements_NEVER SelfCloseVoidElements = "NEVER"
	// The `/` inside void elements is always added (always).
	// Experimental.
	SelfCloseVoidElements_ALWAYS SelfCloseVoidElements = "ALWAYS"
)

