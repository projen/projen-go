package build

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Options for `BuildWorkflow.addPostBuildJobTask`.
// Experimental.
type AddPostBuildJobTaskOptions struct {
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Tools that should be installed before the task is run.
	// Experimental.
	Tools *workflows.Tools `field:"optional" json:"tools" yaml:"tools"`
}

