package javascript


// Experimental.
type ProseWrap string

const (
	// Wrap prose if it exceeds the print width.
	// Experimental.
	ProseWrap_ALWAYS ProseWrap = "ALWAYS"
	// Do not wrap prose.
	// Experimental.
	ProseWrap_NEVER ProseWrap = "NEVER"
	// Wrap prose as-is.
	// Experimental.
	ProseWrap_PRESERVE ProseWrap = "PRESERVE"
)

