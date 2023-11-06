package javascript


// https://yarnpkg.com/configuration/yarnrc#defaultSemverRangePrefix.
// Experimental.
type YarnDefaultSemverRangePrefix string

const (
	// Experimental.
	YarnDefaultSemverRangePrefix_CARET YarnDefaultSemverRangePrefix = "CARET"
	// Experimental.
	YarnDefaultSemverRangePrefix_TILDE YarnDefaultSemverRangePrefix = "TILDE"
	// Experimental.
	YarnDefaultSemverRangePrefix_EMPTY_STRING YarnDefaultSemverRangePrefix = "EMPTY_STRING"
)

