package javascript


// https://yarnpkg.com/configuration/yarnrc#pnpFallbackMode.
// Experimental.
type YarnPnpFallbackMode string

const (
	// Experimental.
	YarnPnpFallbackMode_NONE YarnPnpFallbackMode = "NONE"
	// Experimental.
	YarnPnpFallbackMode_DEPENDENCIES_ONLY YarnPnpFallbackMode = "DEPENDENCIES_ONLY"
	// Experimental.
	YarnPnpFallbackMode_ALL YarnPnpFallbackMode = "ALL"
)

