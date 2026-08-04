package javascript


// Experimental.
type PnpmWorkspaceYamlSchemaUpdate struct {
	// When true, pnpm update writes a change intent after updating workspace manifests.
	// Experimental.
	Changeset *bool `field:"optional" json:"changeset" yaml:"changeset"`
	// When true, pnpm update and pnpm outdated also check the GitHub Actions referenced by the repository's workflow files.
	// Experimental.
	GithubActions *bool `field:"optional" json:"githubActions" yaml:"githubActions"`
	// The base URL of the GitHub server that hosts the repositories of the GitHub Actions referenced by the workflow files.
	// Experimental.
	GithubActionsServer *string `field:"optional" json:"githubActionsServer" yaml:"githubActionsServer"`
	// A list of dependency name patterns that pnpm update and pnpm outdated should skip.
	// Experimental.
	IgnoreDeps *[]*string `field:"optional" json:"ignoreDeps" yaml:"ignoreDeps"`
}

