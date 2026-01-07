package uvconfig


// Experimental.
type ResolutionMode string

const (
	// Resolve the highest compatible version of each package.
	//
	// (highest).
	// Experimental.
	ResolutionMode_HIGHEST ResolutionMode = "HIGHEST"
	// Resolve the lowest compatible version of each package.
	//
	// (lowest).
	// Experimental.
	ResolutionMode_LOWEST ResolutionMode = "LOWEST"
	// Resolve the lowest compatible version of any direct dependencies, and the highest compatible version of any transitive dependencies.
	//
	// (lowest-direct).
	// Experimental.
	ResolutionMode_LOWEST_HYPHEN_DIRECT ResolutionMode = "LOWEST_HYPHEN_DIRECT"
)

