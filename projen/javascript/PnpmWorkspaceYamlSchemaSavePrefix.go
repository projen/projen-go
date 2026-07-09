package javascript


// Configure how versions of packages installed to a package.json file get prefixed.
// Experimental.
type PnpmWorkspaceYamlSchemaSavePrefix string

const (
	// ^.
	// Experimental.
	PnpmWorkspaceYamlSchemaSavePrefix_VALUE_CARAT PnpmWorkspaceYamlSchemaSavePrefix = "VALUE_CARAT"
	// ~.
	// Experimental.
	PnpmWorkspaceYamlSchemaSavePrefix_VALUE_TILDE PnpmWorkspaceYamlSchemaSavePrefix = "VALUE_TILDE"
)

