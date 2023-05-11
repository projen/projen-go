package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// The Gitpod component which emits .gitpod.yml.
// Experimental.
type Gitpod interface {
	Component
	IDevEnvironment
	// Direct access to the gitpod configuration (escape hatch).
	// Experimental.
	Config() interface{}
	// Experimental.
	Project() Project
	// Add a task with more granular options.
	//
	// By default, all tasks will be run in parallel. To run tasks in sequence,
	// create a new `Task` and set the other tasks as subtasks.
	// Experimental.
	AddCustomTask(options *GitpodTask)
	// Add a custom Docker image or Dockerfile for the container.
	// Experimental.
	AddDockerImage(image DevEnvironmentDockerImage)
	// Add ports that should be exposed (forwarded) from the container.
	// Experimental.
	AddPorts(ports ...*string)
	// Add a prebuilds configuration for the Gitpod App.
	// Experimental.
	AddPrebuilds(config *GitpodPrebuilds)
	// Add tasks to run when gitpod starts.
	//
	// By default, all tasks will be run in parallel. To run tasks in sequence,
	// create a new `Task` and specify the other tasks as subtasks.
	// Experimental.
	AddTasks(tasks ...Task)
	// Add a list of VSCode extensions that should be automatically installed in the container.
	//
	// These must be in the format defined in the Open VSX registry.
	//
	// Example:
	//   'scala-lang.scala@0.3.9:O5XmjwY5Gz+0oDZAmqneJw=='
	//
	// See: https://www.gitpod.io/docs/vscode-extensions/
	//
	// Experimental.
	AddVscodeExtensions(extensions ...*string)
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

// The jsii proxy struct for Gitpod
type jsiiProxy_Gitpod struct {
	jsiiProxy_Component
	jsiiProxy_IDevEnvironment
}

func (j *jsiiProxy_Gitpod) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Gitpod) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitpod(project Project, options *GitpodOptions) Gitpod {
	_init_.Initialize()

	if err := validateNewGitpodParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Gitpod{}

	_jsii_.Create(
		"projen.Gitpod",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewGitpod_Override(g Gitpod, project Project, options *GitpodOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Gitpod",
		[]interface{}{project, options},
		g,
	)
}

func (g *jsiiProxy_Gitpod) AddCustomTask(options *GitpodTask) {
	if err := g.validateAddCustomTaskParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addCustomTask",
		[]interface{}{options},
	)
}

func (g *jsiiProxy_Gitpod) AddDockerImage(image DevEnvironmentDockerImage) {
	if err := g.validateAddDockerImageParameters(image); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addDockerImage",
		[]interface{}{image},
	)
}

func (g *jsiiProxy_Gitpod) AddPorts(ports ...*string) {
	args := []interface{}{}
	for _, a := range ports {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addPorts",
		args,
	)
}

func (g *jsiiProxy_Gitpod) AddPrebuilds(config *GitpodPrebuilds) {
	if err := g.validateAddPrebuildsParameters(config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addPrebuilds",
		[]interface{}{config},
	)
}

func (g *jsiiProxy_Gitpod) AddTasks(tasks ...Task) {
	args := []interface{}{}
	for _, a := range tasks {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addTasks",
		args,
	)
}

func (g *jsiiProxy_Gitpod) AddVscodeExtensions(extensions ...*string) {
	args := []interface{}{}
	for _, a := range extensions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addVscodeExtensions",
		args,
	)
}

func (g *jsiiProxy_Gitpod) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Gitpod) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Gitpod) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

