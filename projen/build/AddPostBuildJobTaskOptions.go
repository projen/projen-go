package build

import (
	"github.com/projen/projen-go/projen"
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
	// Github Runner Group selection options.
	// Experimental.
	RunsOnGroup *projen.GroupRunnerOptions `field:"optional" json:"runsOnGroup" yaml:"runsOnGroup"`
	// Tools that should be installed before the task is run.
	// Experimental.
	Tools *workflows.Tools `field:"optional" json:"tools" yaml:"tools"`
}

