package javascript


// Configure pnpm.
// Experimental.
type PnpmOptions struct {
	// The `pnpm-workspace.yaml` configuration.
	// Default: - a blank pnpm-workspace.yaml file
	//
	// Experimental.
	WorkspaceYamlOptions *PnpmWorkspaceYamlOptions `field:"optional" json:"workspaceYamlOptions" yaml:"workspaceYamlOptions"`
}

