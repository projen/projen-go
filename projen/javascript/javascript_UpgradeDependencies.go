package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Upgrade node project dependencies.
// Experimental.
type UpgradeDependencies interface {
	projen.Component
	// Container definitions for the upgrade workflow.
	// Experimental.
	ContainerOptions() *workflows.ContainerOptions
	// Experimental.
	SetContainerOptions(val *workflows.ContainerOptions)
	// A task run after the upgrade task.
	// Experimental.
	PostUpgradeTask() projen.Task
	// Experimental.
	Project() projen.Project
	// The upgrade task.
	// Experimental.
	UpgradeTask() projen.Task
	// The workflows that execute the upgrades.
	//
	// One workflow per branch.
	// Experimental.
	Workflows() *[]github.GithubWorkflow
	// Add steps to execute a successful build.
	// Experimental.
	AddPostBuildSteps(steps ...*workflows.JobStep)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for UpgradeDependencies
type jsiiProxy_UpgradeDependencies struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_UpgradeDependencies) ContainerOptions() *workflows.ContainerOptions {
	var returns *workflows.ContainerOptions
	_jsii_.Get(
		j,
		"containerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) PostUpgradeTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postUpgradeTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) UpgradeTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"upgradeTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) Workflows() *[]github.GithubWorkflow {
	var returns *[]github.GithubWorkflow
	_jsii_.Get(
		j,
		"workflows",
		&returns,
	)
	return returns
}


// Experimental.
func NewUpgradeDependencies(project NodeProject, options *UpgradeDependenciesOptions) UpgradeDependencies {
	_init_.Initialize()

	j := jsiiProxy_UpgradeDependencies{}

	_jsii_.Create(
		"projen.javascript.UpgradeDependencies",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewUpgradeDependencies_Override(u UpgradeDependencies, project NodeProject, options *UpgradeDependenciesOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.UpgradeDependencies",
		[]interface{}{project, options},
		u,
	)
}

func (j *jsiiProxy_UpgradeDependencies) SetContainerOptions(val *workflows.ContainerOptions) {
	_jsii_.Set(
		j,
		"containerOptions",
		val,
	)
}

func (u *jsiiProxy_UpgradeDependencies) AddPostBuildSteps(steps ...*workflows.JobStep) {
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		u,
		"addPostBuildSteps",
		args,
	)
}

func (u *jsiiProxy_UpgradeDependencies) PostSynthesize() {
	_jsii_.InvokeVoid(
		u,
		"postSynthesize",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpgradeDependencies) PreSynthesize() {
	_jsii_.InvokeVoid(
		u,
		"preSynthesize",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpgradeDependencies) Synthesize() {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		nil, // no parameters
	)
}

