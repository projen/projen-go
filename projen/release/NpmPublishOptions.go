package release

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Options for npm release.
// Experimental.
type NpmPublishOptions struct {
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
	// Options for publishing npm package to AWS CodeArtifact.
	// Default: - package is not published to.
	//
	// Experimental.
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
	//
	// Only works in supported CI/CD environments.
	// See: https://docs.npmjs.com/generating-provenance-statements
	//
	// Default: - enabled for for public packages using trusted publishing, disabled otherwise.
	//
	// Experimental.
	NpmProvenance *bool `field:"optional" json:"npmProvenance" yaml:"npmProvenance"`
	// GitHub secret which contains the NPM token to use for publishing packages.
	// Default: - "NPM_TOKEN" or "GITHUB_TOKEN" if `registry` is set to `npm.pkg.github.com`.
	//
	// Experimental.
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
	// Experimental.
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
	// Use trusted publishing for publishing to npmjs.com Needs to be pre-configured on npm.js to work.
	//
	// Requires npm CLI version 11.5.1 or later, this is NOT ensured automatically.
	// When used, `npmTokenSecret` will be ignored.
	// See: https://docs.npmjs.com/trusted-publishers
	//
	// Default: - false.
	//
	// Experimental.
	TrustedPublishing *bool `field:"optional" json:"trustedPublishing" yaml:"trustedPublishing"`
}

