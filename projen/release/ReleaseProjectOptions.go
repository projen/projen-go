package release

import (
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/workflows"
)

// Project options for release.
// Experimental.
type ReleaseProjectOptions struct {
	// Version requirement of `publib` which is used to publish modules to npm.
	// Default: "latest".
	//
	// Experimental.
	JsiiReleaseVersion *string `field:"optional" json:"jsiiReleaseVersion" yaml:"jsiiReleaseVersion"`
	// Major version to release from the default branch.
	//
	// If this is specified, we bump the latest version of this major version line.
	// If not specified, we bump the global latest version.
	// Default: - Major version is not enforced.
	//
	// Experimental.
	MajorVersion *float64 `field:"optional" json:"majorVersion" yaml:"majorVersion"`
	// Minimal Major version to release.
	//
	// This can be useful to set to 1, as breaking changes before the 1.x major
	// release are not incrementing the major version number.
	//
	// Can not be set together with `majorVersion`.
	// Default: - No minimum version is being enforced.
	//
	// Experimental.
	MinMajorVersion *float64 `field:"optional" json:"minMajorVersion" yaml:"minMajorVersion"`
	// The npmDistTag to use when publishing from the default branch.
	//
	// To set the npm dist-tag for release branches, set the `npmDistTag` property
	// for each branch.
	// Default: "latest".
	//
	// Experimental.
	NpmDistTag *string `field:"optional" json:"npmDistTag" yaml:"npmDistTag"`
	// Steps to execute after build as part of the release workflow.
	// Default: [].
	//
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `field:"optional" json:"postBuildSteps" yaml:"postBuildSteps"`
	// Bump versions from the default branch as pre-releases (e.g. "beta", "alpha", "pre").
	// Default: - normal semantic versions.
	//
	// Experimental.
	Prerelease *string `field:"optional" json:"prerelease" yaml:"prerelease"`
	// Instead of actually publishing to package managers, just print the publishing command.
	// Default: false.
	//
	// Experimental.
	PublishDryRun *bool `field:"optional" json:"publishDryRun" yaml:"publishDryRun"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Default: false.
	//
	// Experimental.
	PublishTasks *bool `field:"optional" json:"publishTasks" yaml:"publishTasks"`
	// Find commits that should be considered releasable Used to decide if a release is required.
	// Default: ReleasableCommits.everyCommit()
	//
	// Experimental.
	ReleasableCommits projen.ReleasableCommits `field:"optional" json:"releasableCommits" yaml:"releasableCommits"`
	// Defines additional release branches.
	//
	// A workflow will be created for each
	// release branch which will publish releases from commits in this branch.
	// Each release branch _must_ be assigned a major version number which is used
	// to enforce that versions published from that branch always use that major
	// version. If multiple branches are used, the `majorVersion` field must also
	// be provided for the default branch.
	// Default: - no additional branches are used for release. you can use
	// `addBranch()` to add additional branches.
	//
	// Experimental.
	ReleaseBranches *map[string]*BranchOptions `field:"optional" json:"releaseBranches" yaml:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Default: true.
	//
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.continuous()` instead
	ReleaseEveryCommit *bool `field:"optional" json:"releaseEveryCommit" yaml:"releaseEveryCommit"`
	// Create a github issue on every failed publishing task.
	// Default: false.
	//
	// Experimental.
	ReleaseFailureIssue *bool `field:"optional" json:"releaseFailureIssue" yaml:"releaseFailureIssue"`
	// The label to apply to issues indicating publish failures.
	//
	// Only applies if `releaseFailureIssue` is true.
	// Default: "failed-release".
	//
	// Experimental.
	ReleaseFailureIssueLabel *string `field:"optional" json:"releaseFailureIssueLabel" yaml:"releaseFailureIssueLabel"`
	// CRON schedule to trigger new releases.
	// Default: - no scheduled releases.
	//
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.scheduled()` instead
	ReleaseSchedule *string `field:"optional" json:"releaseSchedule" yaml:"releaseSchedule"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Default: "v".
	//
	// Experimental.
	ReleaseTagPrefix *string `field:"optional" json:"releaseTagPrefix" yaml:"releaseTagPrefix"`
	// The release trigger to use.
	// Default: - Continuous releases (`ReleaseTrigger.continuous()`)
	//
	// Experimental.
	ReleaseTrigger ReleaseTrigger `field:"optional" json:"releaseTrigger" yaml:"releaseTrigger"`
	// The name of the default release workflow.
	// Default: "Release".
	//
	// Experimental.
	ReleaseWorkflowName *string `field:"optional" json:"releaseWorkflowName" yaml:"releaseWorkflowName"`
	// A set of workflow steps to execute in order to setup the workflow container.
	// Experimental.
	ReleaseWorkflowSetupSteps *[]*workflows.JobStep `field:"optional" json:"releaseWorkflowSetupSteps" yaml:"releaseWorkflowSetupSteps"`
	// Custom configuration used when creating changelog with standard-version package.
	//
	// Given values either append to default configuration or overwrite values in it.
	// Default: - standard configuration applicable for GitHub repositories.
	//
	// Experimental.
	VersionrcOptions *map[string]interface{} `field:"optional" json:"versionrcOptions" yaml:"versionrcOptions"`
	// Container image to use for GitHub workflows.
	// Default: - default image.
	//
	// Experimental.
	WorkflowContainerImage *string `field:"optional" json:"workflowContainerImage" yaml:"workflowContainerImage"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	WorkflowRunsOn *[]*string `field:"optional" json:"workflowRunsOn" yaml:"workflowRunsOn"`
	// Github Runner Group selection options.
	// Experimental.
	WorkflowRunsOnGroup *projen.GroupRunnerOptions `field:"optional" json:"workflowRunsOnGroup" yaml:"workflowRunsOnGroup"`
}

