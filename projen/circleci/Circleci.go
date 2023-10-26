package circleci

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/circleci/internal"
)

// Circleci Class to manage `.circleci/config.yml`. Check projen's docs for more information.
// See: https://circleci.com/docs/2.0/configuration-reference/
//
// Experimental.
type Circleci interface {
	projen.Component
	// The yaml file for the Circleci pipeline.
	// Experimental.
	File() projen.YamlFile
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Add a Circleci Orb to pipeline.
	//
	// Will throw error if the orb already exists.
	// Experimental.
	AddOrb(name *string, orb *string)
	// add new workflow to existing pipeline.
	// Experimental.
	AddWorkflow(workflow *Workflow)
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

// The jsii proxy struct for Circleci
type jsiiProxy_Circleci struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Circleci) File() projen.YamlFile {
	var returns projen.YamlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Circleci) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Circleci) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewCircleci(project projen.Project, options *CircleCiProps) Circleci {
	_init_.Initialize()

	if err := validateNewCircleciParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Circleci{}

	_jsii_.Create(
		"projen.circleci.Circleci",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewCircleci_Override(c Circleci, project projen.Project, options *CircleCiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.circleci.Circleci",
		[]interface{}{project, options},
		c,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Circleci_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCircleci_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.circleci.Circleci",
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
func Circleci_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCircleci_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.circleci.Circleci",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Circleci) AddOrb(name *string, orb *string) {
	if err := c.validateAddOrbParameters(name, orb); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOrb",
		[]interface{}{name, orb},
	)
}

func (c *jsiiProxy_Circleci) AddWorkflow(workflow *Workflow) {
	if err := c.validateAddWorkflowParameters(workflow); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addWorkflow",
		[]interface{}{workflow},
	)
}

func (c *jsiiProxy_Circleci) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Circleci) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Circleci) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Circleci) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

