package uvconfig


// Experimental.
type PrereleaseMode string

const (
	// Disallow all pre-release versions.
	//
	// (disallow).
	// Experimental.
	PrereleaseMode_DISALLOW PrereleaseMode = "DISALLOW"
	// Allow all pre-release versions.
	//
	// (allow).
	// Experimental.
	PrereleaseMode_ALLOW PrereleaseMode = "ALLOW"
	// Allow pre-release versions if all versions of a package are pre-release.
	//
	// (if-necessary).
	// Experimental.
	PrereleaseMode_IF_HYPHEN_NECESSARY PrereleaseMode = "IF_HYPHEN_NECESSARY"
	// Allow pre-release versions for first-party packages with explicit pre-release markers in their version requirements.
	//
	// (explicit).
	// Experimental.
	PrereleaseMode_EXPLICIT PrereleaseMode = "EXPLICIT"
	// Allow pre-release versions if all versions of a package are pre-release, or if the package has an explicit pre-release marker in its version requirements.
	//
	// (if-necessary-or-explicit).
	// Experimental.
	PrereleaseMode_IF_HYPHEN_NECESSARY_HYPHEN_OR_HYPHEN_EXPLICIT PrereleaseMode = "IF_HYPHEN_NECESSARY_HYPHEN_OR_HYPHEN_EXPLICIT"
)

