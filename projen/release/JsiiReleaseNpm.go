package release

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Deprecated: Use `NpmPublishOptions` instead.
type JsiiReleaseNpm struct {
	// Steps to execute after executing the publishing command.
	//
	// These can be used
	// to add/update the release artifacts ot any other tasks needed.
	//
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPostPublishingSteps`.
	// Deprecated: Use `NpmPublishOptions` instead.
	PostPublishSteps *[]*workflows.JobStep `field:"optional" json:"postPublishSteps" yaml:"postPublishSteps"`
	// Steps to execute before executing the publishing command. These can be used to prepare the artifact for publishing if needed.
	//
	// These steps are executed after `dist/` has been populated with the build
	// output.
	//
	// Note that when using this in `publishToGitHubReleases` this will override steps added via `addGitHubPrePublishingSteps`.
	// Deprecated: Use `NpmPublishOptions` instead.
	PrePublishSteps *[]*workflows.JobStep `field:"optional" json:"prePublishSteps" yaml:"prePublishSteps"`
	// Additional tools to install in the publishing job.
	// Default: - no additional tools are installed.
	//
	// Deprecated: Use `NpmPublishOptions` instead.
	PublishTools *workflows.Tools `field:"optional" json:"publishTools" yaml:"publishTools"`
	// Options for publishing npm package to AWS CodeArtifact.
	// Default: - undefined.
	//
	// Deprecated: Use `NpmPublishOptions` instead.
	CodeArtifactOptions *CodeArtifactOptions `field:"optional" json:"codeArtifactOptions" yaml:"codeArtifactOptions"`
	// Tags can be used to provide an alias instead of version numbers.
	//
	// For example, a project might choose to have multiple streams of development
	// and use a different tag for each stream, e.g., stable, beta, dev, canary.
	//
	// By default, the `latest` tag is used by npm to identify the current version
	// of a package, and `npm install <pkg>` (without any `@<version>` or `@<tag>`
	// specifier) installs the latest tag. Typically, projects only use the
	// `latest` tag for stable release versions, and use other tags for unstable
	// versions such as prereleases.
	//
	// The `next` tag is used by some projects to identify the upcoming version.
	// Default: "latest".
	//
	// Deprecated: Use `npmDistTag` for each release branch instead.
	DistTag *string `field:"optional" json:"distTag" yaml:"distTag"`
	// Should provenance statements be generated when package is published.
	//
	// Note that this component is using `publib` to publish packages,
	// which is using npm internally and supports provenance statements independently of the package manager used.
	// See: https://docs.npmjs.com/generating-provenance-statements
	//
	// Default: - undefined.
	//
	// Deprecated: Use `NpmPublishOptions` instead.
	NpmProvenance *bool `field:"optional" json:"npmProvenance" yaml:"npmProvenance"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Default: - "NPM_TOKEN" or "GITHUB_TOKEN" if `registry` is set to `npm.pkg.github.com`.
	//
	// Deprecated: Use `NpmPublishOptions` instead.
	NpmTokenSecret *string `field:"optional" json:"npmTokenSecret" yaml:"npmTokenSecret"`
	// The domain name of the npm package registry.
	//
	// To publish to GitHub Packages, set this value to `"npm.pkg.github.com"`. In
	// this if `npmTokenSecret` is not specified, it will default to
	// `GITHUB_TOKEN` which means that you will be able to publish to the
	// repository's package store. In this case, make sure `repositoryUrl` is
	// correctly defined.
	//
	// Example:
	//   "npm.pkg.github.com"
	//
	// Default: "registry.npmjs.org"
	//
	// Deprecated: Use `NpmPublishOptions` instead.
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
}

