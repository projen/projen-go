package javascript


// Determines how pnpm resolves dependencies, See https://pnpm.io/settings#resolutionmode.
// Experimental.
type PnpmWorkspaceYamlSchemaResolutionMode string

const (
	// highest.
	// Experimental.
	PnpmWorkspaceYamlSchemaResolutionMode_HIGHEST PnpmWorkspaceYamlSchemaResolutionMode = "HIGHEST"
	// time-based.
	// Experimental.
	PnpmWorkspaceYamlSchemaResolutionMode_TIME_HYPHEN_BASED PnpmWorkspaceYamlSchemaResolutionMode = "TIME_HYPHEN_BASED"
	// lowest-direct.
	// Experimental.
	PnpmWorkspaceYamlSchemaResolutionMode_LOWEST_HYPHEN_DIRECT PnpmWorkspaceYamlSchemaResolutionMode = "LOWEST_HYPHEN_DIRECT"
)

