package github


// Experimental.
type GitHubOptions struct {
	// Whether mergify should be enabled on this repository or not.
	// Experimental.
	Mergify *bool `field:"optional" json:"mergify" yaml:"mergify"`
	// Options for Mergify.
	// Experimental.
	MergifyOptions *MergifyOptions `field:"optional" json:"mergifyOptions" yaml:"mergifyOptions"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Experimental.
	ProjenCredentials GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Deprecated: - use `projenCredentials`.
	ProjenTokenSecret *string `field:"optional" json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// Add a workflow that performs basic checks for pull requests, like validating that PRs follow Conventional Commits.
	// Experimental.
	PullRequestLint *bool `field:"optional" json:"pullRequestLint" yaml:"pullRequestLint"`
	// Options for configuring a pull request linter.
	// Experimental.
	PullRequestLintOptions *PullRequestLintOptions `field:"optional" json:"pullRequestLintOptions" yaml:"pullRequestLintOptions"`
	// Enables GitHub workflows.
	//
	// If this is set to `false`, workflows will not be created.
	// Experimental.
	Workflows *bool `field:"optional" json:"workflows" yaml:"workflows"`
}

