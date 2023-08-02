package build

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Options for `BuildWorkflow.addPostBuildJobCommands`.
// Experimental.
type AddPostBuildJobCommandsOptions struct {
	// Check out the repository at the pull request branch before commands are run.
	// Default: false.
	//
	// Experimental.
	CheckoutRepo *bool `field:"optional" json:"checkoutRepo" yaml:"checkoutRepo"`
	// Install project dependencies before running commands. `checkoutRepo` must also be set to true.
	//
	// Currently only supported for `NodeProject`.
	// Default: false.
	//
	// Experimental.
	InstallDeps *bool `field:"optional" json:"installDeps" yaml:"installDeps"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Tools that should be installed before the commands are run.
	// Experimental.
	Tools *workflows.Tools `field:"optional" json:"tools" yaml:"tools"`
}

