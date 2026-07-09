package javascript


// Defines what linker should be used for installing Node packages.
// Experimental.
type PnpmWorkspaceYamlSchemaNodeLinker string

const (
	// isolated.
	// Experimental.
	PnpmWorkspaceYamlSchemaNodeLinker_ISOLATED PnpmWorkspaceYamlSchemaNodeLinker = "ISOLATED"
	// hoisted.
	// Experimental.
	PnpmWorkspaceYamlSchemaNodeLinker_HOISTED PnpmWorkspaceYamlSchemaNodeLinker = "HOISTED"
	// pnp.
	// Experimental.
	PnpmWorkspaceYamlSchemaNodeLinker_PNP PnpmWorkspaceYamlSchemaNodeLinker = "PNP"
)

