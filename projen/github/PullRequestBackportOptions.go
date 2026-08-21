package github

import (
	"github.com/projen/projen-go/projen"
)

// Experimental.
type PullRequestBackportOptions struct {
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Github Runner Group selection options.
	// Experimental.
	RunsOnGroup *projen.GroupRunnerOptions `field:"optional" json:"runsOnGroup" yaml:"runsOnGroup"`
	// Automatically approve backport PRs if the 'auto approve' workflow is available.
	// Default: true.
	//
	// Experimental.
	AutoApproveBackport *bool `field:"optional" json:"autoApproveBackport" yaml:"autoApproveBackport"`
	// The prefix used to name backport branches.
	//
	// Make sure to include a separator at the end like `/` or `_`.
	// Default: "backport/".
	//
	// Experimental.
	BackportBranchNamePrefix *string `field:"optional" json:"backportBranchNamePrefix" yaml:"backportBranchNamePrefix"`
	// The labels added to the created backport PR.
	// Default: ["backport"].
	//
	// Experimental.
	BackportPRLabels *[]*string `field:"optional" json:"backportPRLabels" yaml:"backportPRLabels"`
	// List of branches that can be a target for backports.
	// Default: - allow backports to all release branches.
	//
	// Experimental.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// Should this created Backport PRs with conflicts.
	//
	// Conflicts will have to be resolved manually, but a PR is always created.
	// Set to `false` to prevent the backport PR from being created if there are conflicts.
	// Default: true.
	//
	// Experimental.
	CreateWithConflicts *bool `field:"optional" json:"createWithConflicts" yaml:"createWithConflicts"`
	// The prefix used to detect PRs that should be backported.
	// Default: "backport-to-".
	//
	// Experimental.
	LabelPrefix *string `field:"optional" json:"labelPrefix" yaml:"labelPrefix"`
	// The name of the workflow.
	// Default: "backport".
	//
	// Experimental.
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

