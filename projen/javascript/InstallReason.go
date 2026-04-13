package javascript


// Why a dependency install was triggered during synthesis.
// Experimental.
type InstallReason string

const (
	// The node_modules directory does not exist.
	// Experimental.
	InstallReason_NO_NODE_MODULES InstallReason = "NO_NODE_MODULES"
	// The package.json file was modified during synthesis.
	// Experimental.
	InstallReason_PACKAGE_JSON_CHANGED InstallReason = "PACKAGE_JSON_CHANGED"
	// Wildcard dependency versions were resolved to concrete ranges.
	// Experimental.
	InstallReason_DEPS_RESOLVED InstallReason = "DEPS_RESOLVED"
)

