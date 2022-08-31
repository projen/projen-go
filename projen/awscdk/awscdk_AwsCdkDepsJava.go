package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
)

// Manages dependencies on the AWS CDK for Java projects.
// Experimental.
type AwsCdkDepsJava interface {
	AwsCdkDeps
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
}

// The jsii proxy struct for AwsCdkDepsJava
type jsiiProxy_AwsCdkDepsJava struct {
	jsiiProxy_AwsCdkDeps
}

func (j *jsiiProxy_AwsCdkDepsJava) CdkDependenciesAsDeps() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"cdkDependenciesAsDeps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJava) CdkMajorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdkMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJava) CdkMinimumVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkMinimumVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJava) CdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCdkDepsJava) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCdkDepsJava(project projen.Project, options *AwsCdkDepsOptions) AwsCdkDepsJava {
	_init_.Initialize()

	if err := validateNewAwsCdkDepsJavaParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCdkDepsJava{}

	_jsii_.Create(
		"projen.awscdk.AwsCdkDepsJava",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCdkDepsJava_Override(a AwsCdkDepsJava, project projen.Project, options *AwsCdkDepsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AwsCdkDepsJava",
		[]interface{}{project, options},
		a,
	)
}

func (a *jsiiProxy_AwsCdkDepsJava) AddV1Dependencies(deps ...*string) {
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

func (a *jsiiProxy_AwsCdkDepsJava) AddV1DevDependencies(deps ...*string) {
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

func (a *jsiiProxy_AwsCdkDepsJava) PackageNames() *AwsCdkPackageNames {
	var returns *AwsCdkPackageNames

	_jsii_.Invoke(
		a,
		"packageNames",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkDepsJava) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDepsJava) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDepsJava) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

