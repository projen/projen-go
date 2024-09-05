package cdk

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Go target configuration.
// Experimental.
type JsiiGoTarget struct {
	// Steps to execute after executing the publishing command.
	//
	// These can be used
	// to add/update the release artifacts ot any other tasks needed.
	//
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPostPublishingSteps`.
	// Experimental.
	PostPublishSteps *[]*workflows.JobStep `field:"optional" json:"postPublishSteps" yaml:"postPublishSteps"`
	// Steps to execute before executing the publishing command. These can be used to prepare the artifact for publishing if needed.
	//
	// These steps are executed after `dist/` has been populated with the build
	// output.
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPrePublishingSteps`.
	// Experimental.
	PrePublishSteps *[]*workflows.JobStep `field:"optional" json:"prePublishSteps" yaml:"prePublishSteps"`
	// Additional tools to install in the publishing job.
	// Default: - no additional tools are installed.
	//
	// Experimental.
	PublishTools *workflows.Tools `field:"optional" json:"publishTools" yaml:"publishTools"`
	// Branch to push to.
	// Default: "main".
	//
	// Experimental.
	GitBranch *string `field:"optional" json:"gitBranch" yaml:"gitBranch"`
	// The commit message.
	// Default: "chore(release): $VERSION".
	//
	// Experimental.
	GitCommitMessage *string `field:"optional" json:"gitCommitMessage" yaml:"gitCommitMessage"`
	// The name of the secret that includes a GitHub deploy key used to push to the GitHub repository.
	//
	// Ignored if `githubUseSsh` is `false`.
	// Default: "GO_GITHUB_DEPLOY_KEY".
	//
	// Experimental.
	GithubDeployKeySecret *string `field:"optional" json:"githubDeployKeySecret" yaml:"githubDeployKeySecret"`
	// The name of the secret that includes a personal GitHub access token used to push to the GitHub repository.
	//
	// Ignored if `githubUseSsh` is `true`.
	// Default: "GO_GITHUB_TOKEN".
	//
	// Experimental.
	GithubTokenSecret *string `field:"optional" json:"githubTokenSecret" yaml:"githubTokenSecret"`
	// Use SSH to push to GitHub instead of a personal accses token.
	// Default: false.
	//
	// Experimental.
	GithubUseSsh *bool `field:"optional" json:"githubUseSsh" yaml:"githubUseSsh"`
	// The email to use in the release git commit.
	// Default: "github-actions@github.com"
	//
	// Experimental.
	GitUserEmail *string `field:"optional" json:"gitUserEmail" yaml:"gitUserEmail"`
	// The user name to use for the release git commit.
	// Default: "github-actions".
	//
	// Experimental.
	GitUserName *string `field:"optional" json:"gitUserName" yaml:"gitUserName"`
	// The name of the target repository in which this module will be published (e.g. github.com/owner/repo).
	//
	// The module itself will always be published under a subdirectory named according
	// to the `packageName` of the module (e.g. github.com/foo/bar/pkg).
	//
	// Example:
	//   github.com/owner/repo
	//
	// Experimental.
	ModuleName *string `field:"required" json:"moduleName" yaml:"moduleName"`
	// The name of the Go package name.
	//
	// If not specified, package name will be derived from the JavaScript module name
	// by removing non-alphanumeric characters (e.g.
	// Default: - derived from the JavaScript module name.
	//
	// Experimental.
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// A suffix appended at the end of the module version (e.g `"-devprefix"`).
	// Default: - none.
	//
	// Experimental.
	VersionSuffix *string `field:"optional" json:"versionSuffix" yaml:"versionSuffix"`
}

