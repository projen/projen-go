package javascript


// Experimental.
type HTMLWhitespaceSensitivity string

const (
	// Respect the default value of CSS display property.
	// Experimental.
	HTMLWhitespaceSensitivity_CSS HTMLWhitespaceSensitivity = "CSS"
	// Whitespaces are considered insignificant.
	// Experimental.
	HTMLWhitespaceSensitivity_IGNORE HTMLWhitespaceSensitivity = "IGNORE"
	// Whitespaces are considered significant.
	// Experimental.
	HTMLWhitespaceSensitivity_STRICT HTMLWhitespaceSensitivity = "STRICT"
)

