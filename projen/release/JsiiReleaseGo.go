package release

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Deprecated: Use `GoPublishOptions` instead.
type JsiiReleaseGo struct {
	// Steps to execute after executing the publishing command.
	//
	// These can be used
	// to add/update the release artifacts ot any other tasks needed.
	//
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPostPublishingSteps`.
	// Deprecated: Use `GoPublishOptions` instead.
	PostPublishSteps *[]*workflows.JobStep `field:"optional" json:"postPublishSteps" yaml:"postPublishSteps"`
	// Steps to execute before executing the publishing command. These can be used to prepare the artifact for publishing if neede.
	//
	// These steps are executed after `dist/` has been populated with the build
	// output.
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPrePublishingSteps`.
	// Deprecated: Use `GoPublishOptions` instead.
	PrePublishSteps *[]*workflows.JobStep `field:"optional" json:"prePublishSteps" yaml:"prePublishSteps"`
	// Additional tools to install in the publishing job.
	// Default: - no additional tools are installed.
	//
	// Deprecated: Use `GoPublishOptions` instead.
	PublishTools *workflows.Tools `field:"optional" json:"publishTools" yaml:"publishTools"`
	// Branch to push to.
	// Default: "main".
	//
	// Deprecated: Use `GoPublishOptions` instead.
	GitBranch *string `field:"optional" json:"gitBranch" yaml:"gitBranch"`
	// The commit message.
	// Default: "chore(release): $VERSION".
	//
	// Deprecated: Use `GoPublishOptions` instead.
	GitCommitMessage *string `field:"optional" json:"gitCommitMessage" yaml:"gitCommitMessage"`
	// The name of the secret that includes a GitHub deploy key used to push to the GitHub repository.
	//
	// Ignored if `githubUseSsh` is `false`.
	// Default: "GO_GITHUB_DEPLOY_KEY".
	//
	// Deprecated: Use `GoPublishOptions` instead.
	GithubDeployKeySecret *string `field:"optional" json:"githubDeployKeySecret" yaml:"githubDeployKeySecret"`
	// The name of the secret that includes a personal GitHub access token used to push to the GitHub repository.
	//
	// Ignored if `githubUseSsh` is `true`.
	// Default: "GO_GITHUB_TOKEN".
	//
	// Deprecated: Use `GoPublishOptions` instead.
	GithubTokenSecret *string `field:"optional" json:"githubTokenSecret" yaml:"githubTokenSecret"`
	// Use SSH to push to GitHub instead of a personal accses token.
	// Default: false.
	//
	// Deprecated: Use `GoPublishOptions` instead.
	GithubUseSsh *bool `field:"optional" json:"githubUseSsh" yaml:"githubUseSsh"`
	// The email to use in the release git commit.
	// Default: "github-actions@github.com"
	//
	// Deprecated: Use `GoPublishOptions` instead.
	GitUserEmail *string `field:"optional" json:"gitUserEmail" yaml:"gitUserEmail"`
	// The user name to use for the release git commit.
	// Default: "github-actions".
	//
	// Deprecated: Use `GoPublishOptions` instead.
	GitUserName *string `field:"optional" json:"gitUserName" yaml:"gitUserName"`
}

