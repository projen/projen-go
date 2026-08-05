package javascript


// Controls where release changelog content is stored.
// Experimental.
type PnpmWorkspaceYamlSchemaVersioningChangelog struct {
	// Changelog storage mode.
	//
	// `registry` composes changelog entries at publish time, while `repository` commits CHANGELOG.md files in packages.
	// Experimental.
	Storage PnpmWorkspaceYamlSchemaVersioningChangelogStorage `field:"optional" json:"storage" yaml:"storage"`
}

