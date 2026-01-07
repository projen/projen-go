package uvconfig


// The default version specifier when adding a dependency.
// Experimental.
type AddBoundsKind string

const (
	// Only a lower bound, e.g., `>=1.2.3`. (lower).
	// Experimental.
	AddBoundsKind_LOWER AddBoundsKind = "LOWER"
	// Allow the same major version, similar to the semver caret, e.g., `>=1.2.3, <2.0.0`.
	//
	// Leading zeroes are skipped, e.g. `>=0.1.2, <0.2.0`. (major)
	// Experimental.
	AddBoundsKind_MAJOR AddBoundsKind = "MAJOR"
	// Allow the same minor version, similar to the semver tilde, e.g., `>=1.2.3, <1.3.0`.
	//
	// Leading zeroes are skipped, e.g. `>=0.1.2, <0.1.3`. (minor)
	// Experimental.
	AddBoundsKind_MINOR AddBoundsKind = "MINOR"
	// Pin the exact version, e.g., `==1.2.3`.
	//
	// This option is not recommended, as versions are already pinned in the uv lockfile. (exact)
	// Experimental.
	AddBoundsKind_EXACT AddBoundsKind = "EXACT"
)

