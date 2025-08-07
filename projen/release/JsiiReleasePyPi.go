package release

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Deprecated: Use `PyPiPublishOptions` instead.
type JsiiReleasePyPi struct {
	// The GitHub Actions environment used for publishing.
	//
	// This can be used to add an explicit approval step to the release
	// or limit who can initiate a release through environment protection rules.
	//
	// Set this to overwrite a package level publishing environment just for this artifact.
	// Default: - no environment used, unless set at the package level.
	//
	// Deprecated: Use `PyPiPublishOptions` instead.
	GithubEnvironment *string `field:"optional" json:"githubEnvironment" yaml:"githubEnvironment"`
	// Steps to execute after executing the publishing command.
	//
	// These can be used
	// to add/update the release artifacts ot any other tasks needed.
	//
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPostPublishingSteps`.
	// Deprecated: Use `PyPiPublishOptions` instead.
	PostPublishSteps *[]*workflows.JobStep `field:"optional" json:"postPublishSteps" yaml:"postPublishSteps"`
	// Steps to execute before executing the publishing command. These can be used to prepare the artifact for publishing if needed.
	//
	// These steps are executed after `dist/` has been populated with the build
	// output.
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPrePublishingSteps`.
	// Deprecated: Use `PyPiPublishOptions` instead.
	PrePublishSteps *[]*workflows.JobStep `field:"optional" json:"prePublishSteps" yaml:"prePublishSteps"`
	// Additional tools to install in the publishing job.
	// Default: - no additional tools are installed.
	//
	// Deprecated: Use `PyPiPublishOptions` instead.
	PublishTools *workflows.Tools `field:"optional" json:"publishTools" yaml:"publishTools"`
	// Options for publishing to AWS CodeArtifact.
	// Default: - undefined.
	//
	// Deprecated: Use `PyPiPublishOptions` instead.
	CodeArtifactOptions *CodeArtifactOptions `field:"optional" json:"codeArtifactOptions" yaml:"codeArtifactOptions"`
	// The GitHub secret which contains PyPI password.
	// Default: "TWINE_PASSWORD".
	//
	// Deprecated: Use `PyPiPublishOptions` instead.
	TwinePasswordSecret *string `field:"optional" json:"twinePasswordSecret" yaml:"twinePasswordSecret"`
	// The registry url to use when releasing packages.
	// Default: - twine default.
	//
	// Deprecated: Use `PyPiPublishOptions` instead.
	TwineRegistryUrl *string `field:"optional" json:"twineRegistryUrl" yaml:"twineRegistryUrl"`
	// The GitHub secret which contains PyPI user name.
	// Default: "TWINE_USERNAME".
	//
	// Deprecated: Use `PyPiPublishOptions` instead.
	TwineUsernameSecret *string `field:"optional" json:"twineUsernameSecret" yaml:"twineUsernameSecret"`
}

