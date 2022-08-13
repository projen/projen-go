package release


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
	// Experimental.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Do not actually publish, only print the commands that would be executed instead.
	//
	// Useful if you wish to block all publishing from a single option.
	// Experimental.
	DryRun *bool `field:"optional" json:"dryRun" yaml:"dryRun"`
	// Create an issue when a publish task fails.
	// Experimental.
	FailureIssue *bool `field:"optional" json:"failureIssue" yaml:"failureIssue"`
	// The label to apply to the issue marking failed publish tasks.
	//
	// Only applies if `failureIssue` is true.
	// Experimental.
	FailureIssueLabel *string `field:"optional" json:"failureIssueLabel" yaml:"failureIssueLabel"`
	// Deprecated: use `publibVersion` instead.
	JsiiReleaseVersion *string `field:"optional" json:"jsiiReleaseVersion" yaml:"jsiiReleaseVersion"`
	// Version requirement for `publib`.
	// Experimental.
	PublibVersion *string `field:"optional" json:"publibVersion" yaml:"publibVersion"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Experimental.
	PublishTasks *bool `field:"optional" json:"publishTasks" yaml:"publishTasks"`
	// Github Runner selection labels.
	// Experimental.
	WorkflowRunsOn *[]*string `field:"optional" json:"workflowRunsOn" yaml:"workflowRunsOn"`
}

