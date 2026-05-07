package javascript

import (
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/workflows"
)

// Build workflow options for NodeProject.
// Experimental.
type BuildWorkflowOptions struct {
	// Build environment variables.
	// Default: {}.
	//
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Name of the buildfile (e.g. "build" becomes "build.yml").
	// Default: "build".
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Permissions granted to the build job To limit job permissions for `contents`, the desired permissions have to be explicitly set, e.g.: `{ contents: JobPermission.NONE }`.
	// Default: `{ contents: JobPermission.WRITE }`
	//
	// Experimental.
	Permissions *workflows.JobPermissions `field:"optional" json:"permissions" yaml:"permissions"`
	// Steps to execute before the build.
	// Default: [].
	//
	// Experimental.
	PreBuildSteps *[]*workflows.JobStep `field:"optional" json:"preBuildSteps" yaml:"preBuildSteps"`
	// Build workflow triggers.
	// Default: "{ pullRequest: {}, workflowDispatch: {} }".
	//
	// Experimental.
	WorkflowTriggers *workflows.Triggers `field:"optional" json:"workflowTriggers" yaml:"workflowTriggers"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means that any files synthesized by projen or e.g. test snapshots will
	// always be up-to-date before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	// Default: true.
	//
	// Experimental.
	MutableBuild *bool `field:"optional" json:"mutableBuild" yaml:"mutableBuild"`
	// Perform a mutable (non-frozen) install during builds.
	//
	// This will update the
	// package lockfile during installs, which is useful when build steps modify
	// dependencies. Set to `false` to use frozen lockfile installs even when
	// `mutableBuild` is enabled.
	// Default: - value of `mutableBuild`.
	//
	// Experimental.
	MutableInstall *bool `field:"optional" json:"mutableInstall" yaml:"mutableInstall"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Github Runner Group selection options.
	// Experimental.
	RunsOnGroup *projen.GroupRunnerOptions `field:"optional" json:"runsOnGroup" yaml:"runsOnGroup"`
}

