package javascript


// Experimental.
type PnpmWorkspaceYamlSchemaUpdateConfig struct {
	// A list of packages that should be ignored when running "pnpm outdated" or "pnpm update --latest".
	// Experimental.
	IgnoreDependencies *[]*string `field:"optional" json:"ignoreDependencies" yaml:"ignoreDependencies"`
}

