package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
)

// Manages dependencies on the AWS CDK for Python projects.
// Experimental.
type AwsCdkDepsPy interface {
	AwsCdkDeps
	// The dependency requirement for the CDK CLI (e.g. '^2.3.4').
	//
	// Will return `^2` if the version was unspecified by the user.
	// Experimental.
	CdkCliVersion() *string
	// Whether CDK dependencies are added as normal dependencies (and peer dependencies).
	// Deprecated: Not used for CDK 2.x
	CdkDependenciesAsDeps() *bool
	// The major version of the AWS CDK (e.g. 1, 2, ...).
	// Experimental.
	CdkMajorVersion() *float64
	// The minimum version of the AWS CDK (e.g. `2.0.0`).
	// Experimental.
	CdkMinimumVersion() *string
	// The dependency requirement for AWS CDK (e.g. `^2.0.0`).
	// Experimental.
	CdkVersion() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Adds dependencies to AWS CDK modules.
	//
	// The type of dependency is determined by the `dependencyType` option.
	//
	// This method is not supported in CDK v2. Use `project.addPeerDeps()` or
	// `project.addDeps()` as appropriate.
	// Experimental.
	AddV1Dependencies(deps ...*string)
	// Adds AWS CDK modules as dev dependencies.
	//
	// This method is not supported in CDK v2. Use `project.addPeerDeps()` or
	// `project.addDeps()` as appropriate.
	// Experimental.
	AddV1DevDependencies(deps ...*string)
	// Return a configuration object with information about package naming in various languages.
	// Experimental.
	PackageNames() *AwsCdkPackageNames
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

// The jsii proxy struct for AwsCdkDepsPy
type jsiiProxy_AwsCdkDepsPy struct {
	jsiiProxy_AwsCdkDeps
}

func (j *jsiiProxy_AwsCdkDepsPy) CdkCliVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkCliVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsPy) CdkDependenciesAsDeps() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"cdkDependenciesAsDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsPy) CdkMajorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdkMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsPy) CdkMinimumVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkMinimumVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsPy) CdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsPy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsPy) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCdkDepsPy(project projen.Project, options *AwsCdkDepsOptions) AwsCdkDepsPy {
	_init_.Initialize()

	if err := validateNewAwsCdkDepsPyParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCdkDepsPy{}

	_jsii_.Create(
		"projen.awscdk.AwsCdkDepsPy",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCdkDepsPy_Override(a AwsCdkDepsPy, project projen.Project, options *AwsCdkDepsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AwsCdkDepsPy",
		[]interface{}{project, options},
		a,
	)
}

// Test whether the given construct is a component.
// Experimental.
func AwsCdkDepsPy_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCdkDepsPy_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.AwsCdkDepsPy",
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
func AwsCdkDepsPy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCdkDepsPy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.AwsCdkDepsPy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkDepsPy) AddV1Dependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addV1Dependencies",
		args,
	)
}

func (a *jsiiProxy_AwsCdkDepsPy) AddV1DevDependencies(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addV1DevDependencies",
		args,
	)
}

func (a *jsiiProxy_AwsCdkDepsPy) PackageNames() *AwsCdkPackageNames {
	var returns *AwsCdkPackageNames

	_jsii_.Invoke(
		a,
		"packageNames",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkDepsPy) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDepsPy) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDepsPy) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDepsPy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

