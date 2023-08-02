package release

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Deprecated: Use `NugetPublishOptions` instead.
type JsiiReleaseNuget struct {
	// Steps to execute before executing the publishing command. These can be used to prepare the artifact for publishing if neede.
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
}

