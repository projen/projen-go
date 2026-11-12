package github


// Dependency version bump levels eligible for Kodiak auto-merge.
// Experimental.
type KodiakDependencyVersion string

const (
	// Major semver updates.
	// Experimental.
	KodiakDependencyVersion_MAJOR KodiakDependencyVersion = "MAJOR"
	// Minor semver updates.
	// Experimental.
	KodiakDependencyVersion_MINOR KodiakDependencyVersion = "MINOR"
	// Patch semver updates.
	// Experimental.
	KodiakDependencyVersion_PATCH KodiakDependencyVersion = "PATCH"
)

