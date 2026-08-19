package javascript


// Controls how node_modules/.package-map.json is generated. standard - only declared dependencies are available through the package map. loose - also maps packages that are reachable through the installed node_modules layout, which can allow undeclared hoisted dependencies to resolve.
// Experimental.
type PnpmWorkspaceYamlSchemaNodePackageMapType string

const (
	// standard.
	// Experimental.
	PnpmWorkspaceYamlSchemaNodePackageMapType_STANDARD PnpmWorkspaceYamlSchemaNodePackageMapType = "STANDARD"
	// loose.
	// Experimental.
	PnpmWorkspaceYamlSchemaNodePackageMapType_LOOSE PnpmWorkspaceYamlSchemaNodePackageMapType = "LOOSE"
)

