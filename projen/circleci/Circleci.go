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
	// Called once, right after `postSynthesize()`, only when the project is created for the first time.
	//
	// It does not run on later `projen` invocations. It only fires for `projen new` (or `Projects.createProject`).
	// It is also skipped when post-synthesis steps are disabled, e.g. `--no-post` or `PROJEN_DISABLE_POST`.
	// Use it for one-off setup that can be turned off by the user, like running a task to give the user immediate
	// feedback on their new project. Order across components is not guaranteed.
	// Experimental.
	PostProjectCreation(initProject *projen.InitProject)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Called once, right after `synthesize()`, only when the project is created for the first time.
	//
	// It does not run on later `projen` invocations. It only fires for `projen new` (or `Projects.createProject`).
	// Use it for deterministic, one-off file generation. Order across components is not guaranteed.
	// Experimental.
	ProjectCreation(initProject *projen.InitProject)
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
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

func (c *jsiiProxy_Circleci) PostProjectCreation(initProject *projen.InitProject) {
	if err := c.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"postProjectCreation",
		[]interface{}{initProject},
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

func (c *jsiiProxy_Circleci) ProjectCreation(initProject *projen.InitProject) {
	if err := c.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"projectCreation",
		[]interface{}{initProject},
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

func (c *jsiiProxy_Circleci) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

