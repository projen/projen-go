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
	// Prefer stable versions, falling back to pre-release versions when necessary.
	//
	// (if-necessary).
	// Experimental.
	PrereleaseMode_IF_HYPHEN_NECESSARY PrereleaseMode = "IF_HYPHEN_NECESSARY"
	// Prefer stable versions for first-party packages with explicit pre-release specifiers, falling back to pre-release versions when necessary.
	//
	// Disallow pre-release versions for all
	// other packages. (explicit)
	// Experimental.
	PrereleaseMode_EXPLICIT PrereleaseMode = "EXPLICIT"
	// Deprecated alias for `if-necessary`.
	//
	// (if-necessary-or-explicit).
	// Experimental.
	PrereleaseMode_IF_HYPHEN_NECESSARY_HYPHEN_OR_HYPHEN_EXPLICIT PrereleaseMode = "IF_HYPHEN_NECESSARY_HYPHEN_OR_HYPHEN_EXPLICIT"
)

