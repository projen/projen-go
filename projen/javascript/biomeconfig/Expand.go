package biomeconfig


// Experimental.
type Expand string

const (
	// Objects are expanded when the first property has a leading newline.
	//
	// Arrays are always
	// expanded if they are shorter than the line width. (auto)
	// Experimental.
	Expand_AUTO Expand = "AUTO"
	// Objects and arrays are always expanded.
	//
	// (always).
	// Experimental.
	Expand_ALWAYS Expand = "ALWAYS"
	// Objects and arrays are never expanded, if they are shorter than the line width.
	//
	// (never).
	// Experimental.
	Expand_NEVER Expand = "NEVER"
)

