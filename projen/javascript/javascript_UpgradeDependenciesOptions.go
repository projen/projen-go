package javascript


// Options for `UpgradeDependencies`.
// Experimental.
type UpgradeDependenciesOptions struct {
	// List of package names to exclude during the upgrade.
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// List of package names to include during the upgrade.
	// Experimental.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
	// Title of the pull request to use (should be all lower-case).
	// Experimental.
	PullRequestTitle *string `field:"optional" json:"pullRequestTitle" yaml:"pullRequestTitle"`
	// Add Signed-off-by line by the committer at the end of the commit log message.
	// Experimental.
	Signoff *bool `field:"optional" json:"signoff" yaml:"signoff"`
	// The name of the task that will be created.
	//
	// This will also be the workflow name.
	// Experimental.
	TaskName *string `field:"optional" json:"taskName" yaml:"taskName"`
	// Include a github workflow for creating PR's that upgrades the required dependencies, either by manual dispatch, or by a schedule.
	//
	// If this is `false`, only a local projen task is created, which can be executed manually to
	// upgrade the dependencies.
	// Experimental.
	Workflow *bool `field:"optional" json:"workflow" yaml:"workflow"`
	// Options for the github workflow.
	//
	// Only applies if `workflow` is true.
	// Experimental.
	WorkflowOptions *UpgradeDependenciesWorkflowOptions `field:"optional" json:"workflowOptions" yaml:"workflowOptions"`
}

