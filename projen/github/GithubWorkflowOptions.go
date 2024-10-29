package github


// Options for `GithubWorkflow`.
// Experimental.
type GithubWorkflowOptions struct {
	// Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time.
	//
	// Currently in beta.
	// See: https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#concurrency
	//
	// Default: - { group: ${{ github.workflow }}, cancelInProgress: false }
	//
	// Experimental.
	ConcurrencyOptions *ConcurrencyOptions `field:"optional" json:"concurrencyOptions" yaml:"concurrencyOptions"`
	// Set a custom file name for the workflow definition file. Must include either a .yml or .yaml file extension.
	//
	// Use this option to set a file name for the workflow file, that is different than the display name.
	//
	// Example:
	//   "my-workflow.yaml"
	//
	// Default: - a path-safe version of the workflow name plus the .yml file ending, e.g. build.yml
	//
	// Experimental.
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// Force the creation of the workflow even if `workflows` is disabled in `GitHub`.
	// Default: false.
	//
	// Experimental.
	Force *bool `field:"optional" json:"force" yaml:"force"`
	// Enable concurrency limitations.
	//
	// Use `concurrencyOptions` to configure specific non default values.
	// Default: false.
	//
	// Experimental.
	LimitConcurrency *bool `field:"optional" json:"limitConcurrency" yaml:"limitConcurrency"`
}

