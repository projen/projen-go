package javascript

import (
	"github.com/projen/projen-go/projen"
)

// Options for `UpgradeDependencies`.
// Experimental.
type UpgradeDependenciesOptions struct {
	// List of package names to exclude during the upgrade.
	// Default: - Nothing is excluded.
	//
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// List of package names to include during the upgrade.
	// Default: - Everything is included.
	//
	// Experimental.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
	// Title of the pull request to use (should be all lower-case).
	// Default: "upgrade dependencies".
	//
	// Experimental.
	PullRequestTitle *string `field:"optional" json:"pullRequestTitle" yaml:"pullRequestTitle"`
	// Check peer dependencies of installed packages and filter updates to compatible versions.
	//
	// By default, the upgrade workflow will adhere to version constraints from peer dependencies.
	// Sometimes this is not desirable and can be disabled.
	// See: https://github.com/raineorshine/npm-check-updates#peer
	//
	// Default: true.
	//
	// Experimental.
	SatisfyPeerDependencies *bool `field:"optional" json:"satisfyPeerDependencies" yaml:"satisfyPeerDependencies"`
	// The semantic commit type.
	// Default: 'chore'.
	//
	// Experimental.
	SemanticCommit *string `field:"optional" json:"semanticCommit" yaml:"semanticCommit"`
	// Add Signed-off-by line by the committer at the end of the commit log message.
	// Default: true.
	//
	// Experimental.
	Signoff *bool `field:"optional" json:"signoff" yaml:"signoff"`
	// Determines the target version to upgrade dependencies to.
	// See: https://github.com/raineorshine/npm-check-updates#target
	//
	// Default: "minor".
	//
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// The name of the task that will be created.
	//
	// This will also be the workflow name.
	// Default: "upgrade".
	//
	// Experimental.
	TaskName *string `field:"optional" json:"taskName" yaml:"taskName"`
	// Specify which dependency types the upgrade should operate on.
	// Default: - All dependency types.
	//
	// Experimental.
	Types *[]projen.DependencyType `field:"optional" json:"types" yaml:"types"`
	// Include a github workflow for creating PR's that upgrades the required dependencies, either by manual dispatch, or by a schedule.
	//
	// If this is `false`, only a local projen task is created, which can be executed manually to
	// upgrade the dependencies.
	// Default: - true for root projects, false for sub-projects.
	//
	// Experimental.
	Workflow *bool `field:"optional" json:"workflow" yaml:"workflow"`
	// Options for the github workflow.
	//
	// Only applies if `workflow` is true.
	// Default: - default options.
	//
	// Experimental.
	WorkflowOptions *UpgradeDependenciesWorkflowOptions `field:"optional" json:"workflowOptions" yaml:"workflowOptions"`
}

