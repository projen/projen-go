package javascript


// off - pnpm writes no package map, and an install removes a map that an earlier install left.
//
// write-only - pnpm writes the package map for other tools to read; Node.js does not load it. load - pnpm also adds Node's --experimental-package-map option to NODE_OPTIONS in pnpm-managed Node.js script environments.
// Experimental.
type PnpmWorkspaceYamlSchemaNodePackageMapMode string

const (
	// off.
	// Experimental.
	PnpmWorkspaceYamlSchemaNodePackageMapMode_OFF PnpmWorkspaceYamlSchemaNodePackageMapMode = "OFF"
	// write-only.
	// Experimental.
	PnpmWorkspaceYamlSchemaNodePackageMapMode_WRITE_HYPHEN_ONLY PnpmWorkspaceYamlSchemaNodePackageMapMode = "WRITE_HYPHEN_ONLY"
	// load.
	// Experimental.
	PnpmWorkspaceYamlSchemaNodePackageMapMode_LOAD PnpmWorkspaceYamlSchemaNodePackageMapMode = "LOAD"
)

