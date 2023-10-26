package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
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

func (j *jsiiProxy_CdkTasks) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

// Test whether the given construct is a component.
// Experimental.
func CdkTasks_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdkTasks_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.CdkTasks",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CdkTasks_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdkTasks_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.CdkTasks",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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

func (c *jsiiProxy_CdkTasks) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

