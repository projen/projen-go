package release

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Options for Go releases.
// Experimental.
type GoPublishOptions struct {
	// Steps to execute before executing the publishing command. These can be used to prepare the artifact for publishing if neede.
	//
	// These steps are executed after `dist/` has been populated with the build
	// output.
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPrePublishingSteps`.
	// Experimental.
	PrePublishSteps *[]*workflows.JobStep `field:"optional" json:"prePublishSteps" yaml:"prePublishSteps"`
	// Additional tools to install in the publishing job.
	// Experimental.
	PublishTools *workflows.Tools `field:"optional" json:"publishTools" yaml:"publishTools"`
	// Branch to push to.
	// Experimental.
	GitBranch *string `field:"optional" json:"gitBranch" yaml:"gitBranch"`
	// The commit message.
	// Experimental.
	GitCommitMessage *string `field:"optional" json:"gitCommitMessage" yaml:"gitCommitMessage"`
	// The name of the secret that includes a GitHub deploy key used to push to the GitHub repository.
	//
	// Ignored if `githubUseSsh` is `false`.
	// Experimental.
	GithubDeployKeySecret *string `field:"optional" json:"githubDeployKeySecret" yaml:"githubDeployKeySecret"`
	// GitHub repository to push to.
	// Experimental.
	GithubRepo *string `field:"optional" json:"githubRepo" yaml:"githubRepo"`
	// The name of the secret that includes a personal GitHub access token used to push to the GitHub repository.
	//
	// Ignored if `githubUseSsh` is `true`.
	// Experimental.
	GithubTokenSecret *string `field:"optional" json:"githubTokenSecret" yaml:"githubTokenSecret"`
	// Use SSH to push to GitHub instead of a personal accses token.
	// Experimental.
	GithubUseSsh *bool `field:"optional" json:"githubUseSsh" yaml:"githubUseSsh"`
	// The email to use in the release git commit.
	// Experimental.
	GitUserEmail *string `field:"optional" json:"gitUserEmail" yaml:"gitUserEmail"`
	// The user name to use for the release git commit.
	// Experimental.
	GitUserName *string `field:"optional" json:"gitUserName" yaml:"gitUserName"`
}

