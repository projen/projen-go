package release

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Deprecated: Use `PyPiPublishOptions` instead.
type JsiiReleasePyPi struct {
	// Steps to execute before executing the publishing command. These can be used to prepare the artifact for publishing if neede.
	//
	// These steps are executed after `dist/` has been populated with the build
	// output.
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPrePublishingSteps`.
	// Deprecated: Use `PyPiPublishOptions` instead.
	PrePublishSteps *[]*workflows.JobStep `field:"optional" json:"prePublishSteps" yaml:"prePublishSteps"`
	// Additional tools to install in the publishing job.
	// Deprecated: Use `PyPiPublishOptions` instead.
	PublishTools *workflows.Tools `field:"optional" json:"publishTools" yaml:"publishTools"`
	// The GitHub secret which contains PyPI password.
	// Deprecated: Use `PyPiPublishOptions` instead.
	TwinePasswordSecret *string `field:"optional" json:"twinePasswordSecret" yaml:"twinePasswordSecret"`
	// The registry url to use when releasing packages.
	// Deprecated: Use `PyPiPublishOptions` instead.
	TwineRegistryUrl *string `field:"optional" json:"twineRegistryUrl" yaml:"twineRegistryUrl"`
	// The GitHub secret which contains PyPI user name.
	// Deprecated: Use `PyPiPublishOptions` instead.
	TwineUsernameSecret *string `field:"optional" json:"twineUsernameSecret" yaml:"twineUsernameSecret"`
}

