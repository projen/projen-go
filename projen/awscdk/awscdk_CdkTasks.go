package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
)

// Adds standard AWS CDK tasks to your project.
// Experimental.
type CdkTasks interface {
	projen.Component
	// Deploys your app.
	// Experimental.
	Deploy() projen.Task
	// Destroys all the stacks.
	// Experimental.
	Destroy() projen.Task
	// Diff against production.
	// Experimental.
	Diff() projen.Task
	// Experimental.
	Project() projen.Project
	// Synthesizes your app.
	// Experimental.
	Synth() projen.Task
	// Synthesizes your app and suppresses stdout.
	// Experimental.
	SynthSilent() projen.Task
	// Watch task.
	// Experimental.
	Watch() projen.Task
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

// The jsii proxy struct for CdkTasks
type jsiiProxy_CdkTasks struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_CdkTasks) Deploy() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"deploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) Destroy() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"destroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) Diff() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"diff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) Synth() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"synth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) SynthSilent() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"synthSilent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkTasks) Watch() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"watch",
		&returns,
	)
	return returns
}


// Experimental.
func NewCdkTasks(project projen.Project) CdkTasks {
	_init_.Initialize()

	if err := validateNewCdkTasksParameters(project); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdkTasks{}

	_jsii_.Create(
		"projen.awscdk.CdkTasks",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewCdkTasks_Override(c CdkTasks, project projen.Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.CdkTasks",
		[]interface{}{project},
		c,
	)
}

func (c *jsiiProxy_CdkTasks) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkTasks) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkTasks) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

