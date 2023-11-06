package javascript


// https://yarnpkg.com/configuration/yarnrc#cacheMigrationMode.
// Experimental.
type YarnCacheMigrationMode string

const (
	// Experimental.
	YarnCacheMigrationMode_REQUIRED_ONLY YarnCacheMigrationMode = "REQUIRED_ONLY"
	// Experimental.
	YarnCacheMigrationMode_MATCH_SPEC YarnCacheMigrationMode = "MATCH_SPEC"
	// Experimental.
	YarnCacheMigrationMode_ALWAYS YarnCacheMigrationMode = "ALWAYS"
)

