package javascript


// Experimental.
type PnpmWorkspaceYamlSchemaVersioningEpics struct {
	// Lead workspace project name or a `./`-prefixed workspace-relative directory.
	// Experimental.
	Lead *string `field:"required" json:"lead" yaml:"lead"`
	// Project selectors matched in order, supporting name globs, `./`-prefixed directory globs, and `!`-prefixed negations.
	// Experimental.
	Packages *[]*string `field:"required" json:"packages" yaml:"packages"`
}

