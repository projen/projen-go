package cdk

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Experimental.
type JsiiPythonTarget struct {
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
	// The GitHub secret which contains PyPI password.
	// Experimental.
	TwinePasswordSecret *string `field:"optional" json:"twinePasswordSecret" yaml:"twinePasswordSecret"`
	// The registry url to use when releasing packages.
	// Experimental.
	TwineRegistryUrl *string `field:"optional" json:"twineRegistryUrl" yaml:"twineRegistryUrl"`
	// The GitHub secret which contains PyPI user name.
	// Experimental.
	TwineUsernameSecret *string `field:"optional" json:"twineUsernameSecret" yaml:"twineUsernameSecret"`
	// Experimental.
	DistName *string `field:"required" json:"distName" yaml:"distName"`
	// Experimental.
	Module *string `field:"required" json:"module" yaml:"module"`
}

