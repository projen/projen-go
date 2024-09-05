package cdk

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Experimental.
type JsiiDotNetTarget struct {
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
	// GitHub secret which contains the API key for NuGet.
	// Default: "NUGET_API_KEY".
	//
	// Experimental.
	NugetApiKeySecret *string `field:"optional" json:"nugetApiKeySecret" yaml:"nugetApiKeySecret"`
	// NuGet Server URL (defaults to nuget.org).
	// Experimental.
	NugetServer *string `field:"optional" json:"nugetServer" yaml:"nugetServer"`
	// Experimental.
	DotNetNamespace *string `field:"required" json:"dotNetNamespace" yaml:"dotNetNamespace"`
	// Experimental.
	PackageId *string `field:"required" json:"packageId" yaml:"packageId"`
	// Experimental.
	IconUrl *string `field:"optional" json:"iconUrl" yaml:"iconUrl"`
}

