package github


// Options for `GithubWorkflow`.
// Experimental.
type GithubWorkflowOptions struct {
	// Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time.
	//
	// Currently in beta.
	// See: https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#concurrency
	//
	// Experimental.
	Concurrency *string `field:"optional" json:"concurrency" yaml:"concurrency"`
	// Force the creation of the workflow even if `workflows` is disabled in `GitHub`.
	// Experimental.
	Force *bool `field:"optional" json:"force" yaml:"force"`
}

