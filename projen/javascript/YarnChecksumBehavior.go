package javascript


// https://yarnpkg.com/configuration/yarnrc#checksumBehavior.
// Experimental.
type YarnChecksumBehavior string

const (
	// Experimental.
	YarnChecksumBehavior_THROW YarnChecksumBehavior = "THROW"
	// Experimental.
	YarnChecksumBehavior_UPDATE YarnChecksumBehavior = "UPDATE"
	// Experimental.
	YarnChecksumBehavior_RESET YarnChecksumBehavior = "RESET"
	// Experimental.
	YarnChecksumBehavior_IGNORE YarnChecksumBehavior = "IGNORE"
)

