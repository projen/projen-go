package javascript


// Added a new hoistingLimits setting for `nodeLinker: hoisted` installs, mirroring yarn's `nmHoistingLimits`.
//
// It accepts `none` (the default — hoist as far as possible), workspaces (hoist only as far as each workspace package), or dependencies (hoist only up to each workspace package's direct dependencies).
// Experimental.
type PnpmWorkspaceYamlSchemaHoistingLimits string

const (
	// node.
	// Experimental.
	PnpmWorkspaceYamlSchemaHoistingLimits_NODE PnpmWorkspaceYamlSchemaHoistingLimits = "NODE"
	// workspaces.
	// Experimental.
	PnpmWorkspaceYamlSchemaHoistingLimits_WORKSPACES PnpmWorkspaceYamlSchemaHoistingLimits = "WORKSPACES"
	// dependencies.
	// Experimental.
	PnpmWorkspaceYamlSchemaHoistingLimits_DEPENDENCIES PnpmWorkspaceYamlSchemaHoistingLimits = "DEPENDENCIES"
)

