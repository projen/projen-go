package javascript


// Sets the default for pnpm publish --access.
//
// When set, it takes precedence over publishConfig.access. Configure it in pnpm-workspace.yaml or the global config.yaml, or with PNPM_CONFIG_ACCESS. Read from configuration since pnpm v12.23.0.
// Experimental.
type PnpmWorkspaceYamlSchemaAccess string

const (
	// public.
	// Experimental.
	PnpmWorkspaceYamlSchemaAccess_PUBLIC PnpmWorkspaceYamlSchemaAccess = "PUBLIC"
	// restricted.
	// Experimental.
	PnpmWorkspaceYamlSchemaAccess_RESTRICTED PnpmWorkspaceYamlSchemaAccess = "RESTRICTED"
)

