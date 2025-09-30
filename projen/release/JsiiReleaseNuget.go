package release

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Deprecated: Use `NugetPublishOptions` instead.
type JsiiReleaseNuget struct {
	// The GitHub Actions environment used for publishing.
	//
	// This can be used to add an explicit approval step to the release
	// or limit who can initiate a release through environment protection rules.
	//
	// Set this to overwrite a package level publishing environment just for this artifact.
	// Default: - no environment used, unless set at the package level.
	//
	// Deprecated: Use `NugetPublishOptions` instead.
	GithubEnvironment *string `field:"optional" json:"githubEnvironment" yaml:"githubEnvironment"`
	// Steps to execute after executing the publishing command.
	//
	// These can be used
	// to add/update the release artifacts ot any other tasks needed.
	//
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPostPublishingSteps`.
	// Deprecated: Use `NugetPublishOptions` instead.
	PostPublishSteps *[]*workflows.JobStep `field:"optional" json:"postPublishSteps" yaml:"postPublishSteps"`
	// Steps to execute before executing the publishing command. These can be used to prepare the artifact for publishing if needed.
	//
	// These steps are executed after `dist/` has been populated with the build
	// output.
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPrePublishingSteps`.
	// Deprecated: Use `NugetPublishOptions` instead.
	PrePublishSteps *[]*workflows.JobStep `field:"optional" json:"prePublishSteps" yaml:"prePublishSteps"`
	// Additional tools to install in the publishing job.
	// Default: - no additional tools are installed.
	//
	// Deprecated: Use `NugetPublishOptions` instead.
	PublishTools *workflows.Tools `field:"optional" json:"publishTools" yaml:"publishTools"`
	// GitHub secret which contains the API key for NuGet.
	// Default: "NUGET_API_KEY".
	//
	// Deprecated: Use `NugetPublishOptions` instead.
	NugetApiKeySecret *string `field:"optional" json:"nugetApiKeySecret" yaml:"nugetApiKeySecret"`
	// NuGet Server URL (defaults to nuget.org).
	// Deprecated: Use `NugetPublishOptions` instead.
	NugetServer *string `field:"optional" json:"nugetServer" yaml:"nugetServer"`
	// The NuGet.org username (profile name, not email address) for trusted publisher authentication.
	//
	// Required when using trusted publishing.
	// Default: "NUGET_USERNAME".
	//
	// Deprecated: Use `NugetPublishOptions` instead.
	NugetUsernameSecret *string `field:"optional" json:"nugetUsernameSecret" yaml:"nugetUsernameSecret"`
	// Use NuGet trusted publishing instead of API keys.
	//
	// Needs to be setup in NuGet.org.
	// See: https://learn.microsoft.com/en-us/nuget/nuget-org/trusted-publishing
	//
	// Deprecated: Use `NugetPublishOptions` instead.
	TrustedPublishing *bool `field:"optional" json:"trustedPublishing" yaml:"trustedPublishing"`
}

