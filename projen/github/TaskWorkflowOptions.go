package github

import (
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/workflows"
)

// Experimental.
type TaskWorkflowOptions struct {
	// The workflow name.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Permissions for the build job.
	// Experimental.
	Permissions *workflows.JobPermissions `field:"required" json:"permissions" yaml:"permissions"`
	// The main task to be executed.
	// Experimental.
	Task projen.Task `field:"required" json:"task" yaml:"task"`
	// A directory name which contains artifacts to be uploaded (e.g. `dist`). If this is set, the contents of this directory will be uploaded as an artifact at the end of the workflow run, even if other steps fail.
	// Default: - not set.
	//
	// Experimental.
	ArtifactsDirectory *string `field:"optional" json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// Override for the `with` property of the source code checkout step.
	// Default: - not set.
	//
	// Experimental.
	CheckoutWith *map[string]interface{} `field:"optional" json:"checkoutWith" yaml:"checkoutWith"`
	// Adds an 'if' condition to the workflow.
	// Experimental.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Default: - default image.
	//
	// Experimental.
	Container *workflows.ContainerOptions `field:"optional" json:"container" yaml:"container"`
	// Whether to download files from Git LFS for this workflow.
	// Default: - Use the setting on the corresponding GitHub project.
	//
	// Experimental.
	DownloadLfs *bool `field:"optional" json:"downloadLfs" yaml:"downloadLfs"`
	// Workflow environment variables.
	// Default: {}.
	//
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// The git identity to use in this workflow.
	// Experimental.
	GitIdentity *GitIdentity `field:"optional" json:"gitIdentity" yaml:"gitIdentity"`
	// The primary job id.
	// Default: "build".
	//
	// Experimental.
	JobId *string `field:"optional" json:"jobId" yaml:"jobId"`
	// Mapping of job output names to values/expressions.
	// Default: {}.
	//
	// Experimental.
	Outputs *map[string]*workflows.JobStepOutput `field:"optional" json:"outputs" yaml:"outputs"`
	// Actions to run after the main build step.
	// Default: - not set.
	//
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `field:"optional" json:"postBuildSteps" yaml:"postBuildSteps"`
	// Steps to run before the main build step.
	// Default: - not set.
	//
	// Experimental.
	PreBuildSteps *[]*workflows.JobStep `field:"optional" json:"preBuildSteps" yaml:"preBuildSteps"`
	// Initial steps to run before the source code checkout.
	// Default: - not set.
	//
	// Experimental.
	PreCheckoutSteps *[]*workflows.JobStep `field:"optional" json:"preCheckoutSteps" yaml:"preCheckoutSteps"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// The triggers for the workflow.
	// Default: - by default workflows can only be triggered by manually.
	//
	// Experimental.
	Triggers *workflows.Triggers `field:"optional" json:"triggers" yaml:"triggers"`
}

