package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
)

// Manages dependencies on the AWS CDK.
// Experimental.
type AwsCdkDeps interface {
	projen.Component
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

// The jsii proxy struct for AwsCdkDeps
type jsiiProxy_AwsCdkDeps struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_AwsCdkDeps) CdkDependenciesAsDeps() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"cdkDependenciesAsDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDeps) CdkMajorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdkMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDeps) CdkMinimumVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkMinimumVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDeps) CdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDeps) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDeps) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCdkDeps_Override(a AwsCdkDeps, project projen.Project, options *AwsCdkDepsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AwsCdkDeps",
		[]interface{}{project, options},
		a,
	)
}

// Test whether the given construct is a component.
// Experimental.
func AwsCdkDeps_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCdkDeps_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.AwsCdkDeps",
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
func AwsCdkDeps_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCdkDeps_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.AwsCdkDeps",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkDeps) AddV1Dependencies(deps ...*string) {
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

func (a *jsiiProxy_AwsCdkDeps) AddV1DevDependencies(deps ...*string) {
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

func (a *jsiiProxy_AwsCdkDeps) PackageNames() *AwsCdkPackageNames {
	var returns *AwsCdkPackageNames

	_jsii_.Invoke(
		a,
		"packageNames",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkDeps) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDeps) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDeps) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDeps) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

