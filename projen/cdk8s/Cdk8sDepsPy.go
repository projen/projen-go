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

