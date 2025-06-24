package biomeconfig


// Controls whether void-elements should be self closed.
// Experimental.
type SelfCloseVoidElements string

const (
	// never.
	// Experimental.
	SelfCloseVoidElements_NEVER SelfCloseVoidElements = "NEVER"
	// always.
	// Experimental.
	SelfCloseVoidElements_ALWAYS SelfCloseVoidElements = "ALWAYS"
)

