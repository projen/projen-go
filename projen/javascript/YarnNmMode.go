package javascript


// https://yarnpkg.com/configuration/yarnrc#nmMode.
// Experimental.
type YarnNmMode string

const (
	// Experimental.
	YarnNmMode_CLASSIC YarnNmMode = "CLASSIC"
	// Experimental.
	YarnNmMode_HARDLINKS_LOCAL YarnNmMode = "HARDLINKS_LOCAL"
	// Experimental.
	YarnNmMode_HARDLINKS_GLOBAL YarnNmMode = "HARDLINKS_GLOBAL"
)

