package javascript

import (
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
)

// Options for `UpgradeDependencies.workflowOptions`.
// Experimental.
type UpgradeDependenciesWorkflowOptions struct {
	// Assignees to add on the PR.
	// Default: - no assignees.
	//
	// Experimental.
	Assignees *[]*string `field:"optional" json:"assignees" yaml:"assignees"`
	// List of branches to create PR's for.
	// Default: - All release branches configured for the project.
	//
	// Experimental.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// Job container options.
	// Default: - defaults.
	//
	// Experimental.
	Container *workflows.ContainerOptions `field:"optional" json:"container" yaml:"container"`
	// The git identity to use for commits.
	// Default: "github-actions@github.com"
	//
	// Experimental.
	GitIdentity *github.GitIdentity `field:"optional" json:"gitIdentity" yaml:"gitIdentity"`
	// Labels to apply on the PR.
	// Default: - no labels.
	//
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Permissions granted to the upgrade job To limit job permissions for `contents`, the desired permissions have to be explicitly set, e.g.: `{ contents: JobPermission.NONE }`.
	// Default: `{ contents: JobPermission.READ }`
	//
	// Experimental.
	Permissions *workflows.JobPermissions `field:"optional" json:"permissions" yaml:"permissions"`
	// Choose a method for authenticating with GitHub for creating the PR.
	//
	// When using the default github token, PR's created by this workflow
	// will not trigger any subsequent workflows (i.e the build workflow), so
	// projen requires API access to be provided through e.g. a personal
	// access token or other method.
	// See: https://github.com/peter-evans/create-pull-request/issues/48
	//
	// Default: - personal access token named PROJEN_GITHUB_TOKEN.
	//
	// Experimental.
	ProjenCredentials github.GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Schedule to run on.
	// Default: UpgradeDependenciesSchedule.DAILY
	//
	// Experimental.
	Schedule UpgradeDependenciesSchedule `field:"optional" json:"schedule" yaml:"schedule"`
}

