package build

import (
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
)

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
	// The task to execute in order to build the project.
	// Experimental.
	BuildTask projen.Task `field:"required" json:"buildTask" yaml:"buildTask"`
	// A name of a directory that includes build artifacts.
	// Default: "dist".
	//
	// Experimental.
	ArtifactsDirectory *string `field:"optional" json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// The container image to use for builds.
	// Default: - the default workflow container.
	//
	// Experimental.
	ContainerImage *string `field:"optional" json:"containerImage" yaml:"containerImage"`
	// Git identity to use for the workflow.
	// Default: - default identity.
	//
	// Experimental.
	GitIdentity *github.GitIdentity `field:"optional" json:"gitIdentity" yaml:"gitIdentity"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means that any files synthesized by projen or e.g. test snapshots will
	// always be up-to-date before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	//
	// This is enabled by default only if `githubTokenSecret` is set. Otherwise it
	// is disabled, which implies that file changes that happen during build will
	// not be pushed back to the branch.
	// Default: true.
	//
	// Experimental.
	MutableBuild *bool `field:"optional" json:"mutableBuild" yaml:"mutableBuild"`
	// Steps to execute after build.
	// Default: [].
	//
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `field:"optional" json:"postBuildSteps" yaml:"postBuildSteps"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Github Runner Group selection options.
	// Experimental.
	RunsOnGroup *projen.GroupRunnerOptions `field:"optional" json:"runsOnGroup" yaml:"runsOnGroup"`
}

