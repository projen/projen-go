package javascript


// Caps the bump that a release from the current checkout may apply, after dependent propagation and fixed-group resolution.
// Experimental.
type PnpmWorkspaceYamlSchemaVersioningMaxBump string

const (
	// patch.
	// Experimental.
	PnpmWorkspaceYamlSchemaVersioningMaxBump_PATCH PnpmWorkspaceYamlSchemaVersioningMaxBump = "PATCH"
	// minor.
	// Experimental.
	PnpmWorkspaceYamlSchemaVersioningMaxBump_MINOR PnpmWorkspaceYamlSchemaVersioningMaxBump = "MINOR"
	// major.
	// Experimental.
	PnpmWorkspaceYamlSchemaVersioningMaxBump_MAJOR PnpmWorkspaceYamlSchemaVersioningMaxBump = "MAJOR"
)

