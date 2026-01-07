package uvconfig


// Experimental.
type ForkStrategy string

const (
	// Optimize for selecting the fewest number of versions for each package.
	//
	// Older versions may
	// be preferred if they are compatible with a wider range of supported Python versions or
	// platforms. (fewest)
	// Experimental.
	ForkStrategy_FEWEST ForkStrategy = "FEWEST"
	// Optimize for selecting latest supported version of each package, for each supported Python version.
	//
	// (requires-python).
	// Experimental.
	ForkStrategy_REQUIRES_HYPHEN_PYTHON ForkStrategy = "REQUIRES_HYPHEN_PYTHON"
)

