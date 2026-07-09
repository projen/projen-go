package javascript


// Controlling if and how dependencies are added to the default catalog.
// Experimental.
type PnpmWorkspaceYamlSchemaCatalogMode string

const (
	// strict.
	// Experimental.
	PnpmWorkspaceYamlSchemaCatalogMode_STRICT PnpmWorkspaceYamlSchemaCatalogMode = "STRICT"
	// prefer.
	// Experimental.
	PnpmWorkspaceYamlSchemaCatalogMode_PREFER PnpmWorkspaceYamlSchemaCatalogMode = "PREFER"
	// manual.
	// Experimental.
	PnpmWorkspaceYamlSchemaCatalogMode_MANUAL PnpmWorkspaceYamlSchemaCatalogMode = "MANUAL"
)

