package cdk

import (
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/release"
)

// Experimental.
type JsiiPythonTarget struct {
	// The GitHub Actions environment used for publishing.
	//
	// This can be used to add an explicit approval step to the release
	// or limit who can initiate a release through environment protection rules.
	//
	// Set this to overwrite a package level publishing environment just for this artifact.
	// Default: - no environment used, unless set at the package level.
	//
	// Experimental.
	GithubEnvironment *string `field:"optional" json:"githubEnvironment" yaml:"githubEnvironment"`
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
	// Generate and publish cryptographic attestations for files uploaded to PyPI.
	//
	// Attestations provide package provenance and integrity an can be viewed on PyPI.
	// They are only available when using a Trusted Publisher for publishing.
	// See: https://docs.pypi.org/attestations/producing-attestations/
	//
	// Default: - enabled when using trusted publishing, otherwise not applicable.
	//
	// Experimental.
	Attestations *bool `field:"optional" json:"attestations" yaml:"attestations"`
	// Options for publishing to AWS CodeArtifact.
	// Default: - undefined.
	//
	// Experimental.
	CodeArtifactOptions *release.CodeArtifactOptions `field:"optional" json:"codeArtifactOptions" yaml:"codeArtifactOptions"`
	// Use PyPI trusted publishing instead of tokens or username & password.
	//
	// Needs to be setup in PyPI.
	// See: https://docs.pypi.org/trusted-publishers/adding-a-publisher/
	//
	// Experimental.
	TrustedPublishing *bool `field:"optional" json:"trustedPublishing" yaml:"trustedPublishing"`
	// The GitHub secret which contains PyPI password.
	// Default: "TWINE_PASSWORD".
	//
	// Experimental.
	TwinePasswordSecret *string `field:"optional" json:"twinePasswordSecret" yaml:"twinePasswordSecret"`
	// The registry url to use when releasing packages.
	// Default: - twine default.
	//
	// Experimental.
	TwineRegistryUrl *string `field:"optional" json:"twineRegistryUrl" yaml:"twineRegistryUrl"`
	// The GitHub secret which contains PyPI user name.
	// Default: "TWINE_USERNAME".
	//
	// Experimental.
	TwineUsernameSecret *string `field:"optional" json:"twineUsernameSecret" yaml:"twineUsernameSecret"`
	// Experimental.
	DistName *string `field:"required" json:"distName" yaml:"distName"`
	// Experimental.
	Module *string `field:"required" json:"module" yaml:"module"`
}

