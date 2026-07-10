package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
)

// Experimental.
type Cdk8sDepsPy interface {
	Cdk8sDeps
	// The major version of the CDK8s (e.g. 1, 2, ...).
	// Experimental.
	Cdk8sMajorVersion() *float64
	// The minimum version of the CDK8s (e.g. `2.0.0`).
	// Experimental.
	Cdk8sMinimumVersion() *string
	// The dependency requirement for CDK8s.
	// Experimental.
	Cdk8sVersion() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Return a configuration object with information about package naming in various languages.
	// Experimental.
	PackageNames() *Cdk8sPackageNames
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

// The jsii proxy struct for Cdk8sDepsPy
type jsiiProxy_Cdk8sDepsPy struct {
	jsiiProxy_Cdk8sDeps
}

func (j *jsiiProxy_Cdk8sDepsPy) Cdk8sMajorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdk8sMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sDepsPy) Cdk8sMinimumVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdk8sMinimumVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sDepsPy) Cdk8sVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdk8sVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sDepsPy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cdk8sDepsPy) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewCdk8sDepsPy(project projen.Project, options *Cdk8sDepsOptions) Cdk8sDepsPy {
	_init_.Initialize()

	if err := validateNewCdk8sDepsPyParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cdk8sDepsPy{}

	_jsii_.Create(
		"projen.cdk8s.Cdk8sDepsPy",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewCdk8sDepsPy_Override(c Cdk8sDepsPy, project projen.Project, options *Cdk8sDepsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdk8s.Cdk8sDepsPy",
		[]interface{}{project, options},
		c,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Cdk8sDepsPy_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdk8sDepsPy_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.cdk8s.Cdk8sDepsPy",
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
func Cdk8sDepsPy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdk8sDepsPy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.cdk8s.Cdk8sDepsPy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cdk8sDepsPy) PackageNames() *Cdk8sPackageNames {
	var returns *Cdk8sPackageNames

	_jsii_.Invoke(
		c,
		"packageNames",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cdk8sDepsPy) PostProjectCreation(initProject *projen.InitProject) {
	if err := c.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"postProjectCreation",
		[]interface{}{initProject},
	)
}

func (c *jsiiProxy_Cdk8sDepsPy) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cdk8sDepsPy) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cdk8sDepsPy) ProjectCreation(initProject *projen.InitProject) {
	if err := c.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (c *jsiiProxy_Cdk8sDepsPy) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cdk8sDepsPy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cdk8sDepsPy) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

