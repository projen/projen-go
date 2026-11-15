package javascript


// Cargo (Rust) dependencies support: when enabled, `pnpm install` resolves and installs the workspace's Cargo dependencies next to the npm packages.
//
// Multi-ecosystem support is experimental. Added in: v12.4.0 (pnpm v12 only).
// Experimental.
type PnpmWorkspaceYamlSchemaCargo struct {
	// Whether pnpm install resolves and installs the workspace's Cargo dependencies.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

