package biomeconfig


// Experimental.
type Expand string

const (
	// Objects are expanded when the first property has a leading newline.
	//
	// Arrays remain on one
	// line if they fit. (auto)
	// Experimental.
	Expand_AUTO Expand = "AUTO"
	// Objects and arrays are always expanded.
	//
	// (always).
	// Experimental.
	Expand_ALWAYS Expand = "ALWAYS"
	// Objects and arrays remain on one line if they fit.
	//
	// (never).
	// Experimental.
	Expand_NEVER Expand = "NEVER"
)

