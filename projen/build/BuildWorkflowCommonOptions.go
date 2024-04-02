package build

import (
	"github.com/projen/projen-go/projen/github/workflows"
)

// Experimental.
type BuildWorkflowCommonOptions struct {
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
}

