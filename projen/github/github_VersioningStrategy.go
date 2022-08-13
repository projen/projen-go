package github


// The strategy to use when edits manifest and lock files.
// Experimental.
type VersioningStrategy string

const (
	// Only create pull requests to update lockfiles updates.
	//
	// Ignore any new
	// versions that would require package manifest changes.
	// Experimental.
	VersioningStrategy_LOCKFILE_ONLY VersioningStrategy = "LOCKFILE_ONLY"
	// - For apps, the version requirements are increased.
	//
	// - For libraries, the range of versions is widened.
	// Experimental.
	VersioningStrategy_AUTO VersioningStrategy = "AUTO"
	// Relax the version requirement to include both the new and old version, when possible.
	// Experimental.
	VersioningStrategy_WIDEN VersioningStrategy = "WIDEN"
	// Always increase the version requirement to match the new version.
	// Experimental.
	VersioningStrategy_INCREASE VersioningStrategy = "INCREASE"
	// Increase the version requirement only when required by the new version.
	// Experimental.
	VersioningStrategy_INCREASE_IF_NECESSARY VersioningStrategy = "INCREASE_IF_NECESSARY"
)

