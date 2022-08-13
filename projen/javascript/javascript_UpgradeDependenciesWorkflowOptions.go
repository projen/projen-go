package javascript

import (
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
)

// Options for `UpgradeDependencies.workflowOptions`.
// Experimental.
type UpgradeDependenciesWorkflowOptions struct {
	// Assignees to add on the PR.
	// Experimental.
	Assignees *[]*string `field:"optional" json:"assignees" yaml:"assignees"`
	// List of branches to create PR's for.
	// Experimental.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// Job container options.
	// Experimental.
	Container *workflows.ContainerOptions `field:"optional" json:"container" yaml:"container"`
	// The git identity to use for commits.
	// Experimental.
	GitIdentity *github.GitIdentity `field:"optional" json:"gitIdentity" yaml:"gitIdentity"`
	// Labels to apply on the PR.
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Choose a method for authenticating with GitHub for creating the PR.
	//
	// When using the default github token, PR's created by this workflow
	// will not trigger any subsequent workflows (i.e the build workflow), so
	// projen requires API access to be provided through e.g. a personal
	// access token or other method.
	// See: https://github.com/peter-evans/create-pull-request/issues/48
	//
	// Experimental.
	ProjenCredentials github.GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Schedule to run on.
	// Experimental.
	Schedule UpgradeDependenciesSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

