package release

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Publishing options for GitHub releases.
// Experimental.
type GitHubReleasesPublishOptions struct {
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
	// The location of an .md file (relative to `dist/`) that includes the changelog for the release.
	//
	// Example:
	//   changelog.md
	//
	// Experimental.
	ChangelogFile *string `field:"required" json:"changelogFile" yaml:"changelogFile"`
	// The location of a text file (relative to `dist/`) that contains the release tag.
	//
	// Example:
	//   releasetag.txt
	//
	// Experimental.
	ReleaseTagFile *string `field:"required" json:"releaseTagFile" yaml:"releaseTagFile"`
	// The location of a text file (relative to `dist/`) that contains the version number.
	//
	// Example:
	//   version.txt
	//
	// Experimental.
	VersionFile *string `field:"required" json:"versionFile" yaml:"versionFile"`
}

