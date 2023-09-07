package release

import (
	"github.com/projen/projen-go/projen"
)

// Options for `Publisher`.
// Experimental.
type PublisherOptions struct {
	// The name of the artifact to download (e.g. `dist`).
	//
	// The artifact is expected to include a subdirectory for each release target:
	// `go` (GitHub), `dotnet` (NuGet), `java` (Maven), `js` (npm), `python`
	// (PyPI).
	// See: https://github.com/aws/publib
	//
	// Experimental.
	ArtifactName *string `field:"required" json:"artifactName" yaml:"artifactName"`
	// The job ID that produces the build artifacts.
	//
	// All publish jobs will take a dependency on this job.
	// Experimental.
	BuildJobId *string `field:"required" json:"buildJobId" yaml:"buildJobId"`
	// A GitHub workflow expression used as a condition for publishers.
	// Default: - no condition.
	//
	// Experimental.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Do not actually publish, only print the commands that would be executed instead.
	//
	// Useful if you wish to block all publishing from a single option.
	// Experimental.
	DryRun *bool `field:"optional" json:"dryRun" yaml:"dryRun"`
	// Create an issue when a publish task fails.
	// Default: false.
	//
	// Experimental.
	FailureIssue *bool `field:"optional" json:"failureIssue" yaml:"failureIssue"`
	// The label to apply to the issue marking failed publish tasks.
	//
	// Only applies if `failureIssue` is true.
	// Default: "failed-release".
	//
	// Experimental.
	FailureIssueLabel *string `field:"optional" json:"failureIssueLabel" yaml:"failureIssueLabel"`
	// Deprecated: use `publibVersion` instead.
	JsiiReleaseVersion *string `field:"optional" json:"jsiiReleaseVersion" yaml:"jsiiReleaseVersion"`
	// Version requirement for `publib`.
	// Default: "latest".
	//
	// Experimental.
	PublibVersion *string `field:"optional" json:"publibVersion" yaml:"publibVersion"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Default: false.
	//
	// Experimental.
	PublishTasks *bool `field:"optional" json:"publishTasks" yaml:"publishTasks"`
	// Container image to use for GitHub workflows.
	// Default: - default image.
	//
	// Experimental.
	WorkflowContainerImage *string `field:"optional" json:"workflowContainerImage" yaml:"workflowContainerImage"`
	// Node version to setup in GitHub workflows if any node-based CLI utilities are needed.
	//
	// For example `publib`, the CLI projen uses to publish releases,
	// is an npm library.
	// Default: 16.x
	//
	// Experimental.
	WorkflowNodeVersion *string `field:"optional" json:"workflowNodeVersion" yaml:"workflowNodeVersion"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	WorkflowRunsOn *[]*string `field:"optional" json:"workflowRunsOn" yaml:"workflowRunsOn"`
	// Github Runner Group selection options.
	// Experimental.
	WorkflowRunsOnGroup *projen.GroupRunnerOptions `field:"optional" json:"workflowRunsOnGroup" yaml:"workflowRunsOnGroup"`
}

