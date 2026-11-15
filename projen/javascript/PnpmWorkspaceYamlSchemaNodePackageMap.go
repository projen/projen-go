package javascript


// Groups the package-map settings, which control node_modules/.package-map.json during isolated and hoisted installs. Replaces nodeExperimentalPackageMap and nodePackageMapType, which still work; nodePackageMap takes precedence when both are set in the same place. Added in pnpm v12.13.0.
// Experimental.
type PnpmWorkspaceYamlSchemaNodePackageMap struct {
	// off - pnpm writes no package map, and an install removes a map that an earlier install left.
	//
	// write-only - pnpm writes the package map for other tools to read; Node.js does not load it. load - pnpm also adds Node's --experimental-package-map option to NODE_OPTIONS in pnpm-managed Node.js script environments.
	// Experimental.
	Mode PnpmWorkspaceYamlSchemaNodePackageMapMode `field:"optional" json:"mode" yaml:"mode"`
	// Controls how node_modules/.package-map.json is generated. standard - only declared dependencies are available through the package map. loose - also maps packages that are reachable through the installed node_modules layout, which can allow undeclared hoisted dependencies to resolve.
	// Experimental.
	Type PnpmWorkspaceYamlSchemaNodePackageMapType `field:"optional" json:"type" yaml:"type"`
}

