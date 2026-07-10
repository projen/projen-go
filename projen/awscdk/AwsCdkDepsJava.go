package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
)

// Manages dependencies on the AWS CDK for Java projects.
// Experimental.
type AwsCdkDepsJava interface {
	AwsCdkDeps
	// The dependency requirement for the CDK CLI (e.g. '^2.3.4').
	//
	// Will return `^2` if the version was unspecified by the user.
	// Experimental.
	CdkCliVersion() *string
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
	// Return a configuration object with information about package naming in various languages.
	// Experimental.
	PackageNames() *AwsCdkPackageNames
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

// The jsii proxy struct for AwsCdkDepsJava
type jsiiProxy_AwsCdkDepsJava struct {
	jsiiProxy_AwsCdkDeps
}

func (j *jsiiProxy_AwsCdkDepsJava) CdkCliVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkCliVersion",
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

func (j *jsiiProxy_AwsCdkDepsJava) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

// Test whether the given construct is a component.
// Experimental.
func AwsCdkDepsJava_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCdkDepsJava_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.AwsCdkDepsJava",
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
func AwsCdkDepsJava_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCdkDepsJava_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.AwsCdkDepsJava",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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

func (a *jsiiProxy_AwsCdkDepsJava) PostProjectCreation(initProject *projen.InitProject) {
	if err := a.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"postProjectCreation",
		[]interface{}{initProject},
	)
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

func (a *jsiiProxy_AwsCdkDepsJava) ProjectCreation(initProject *projen.InitProject) {
	if err := a.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (a *jsiiProxy_AwsCdkDepsJava) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCdkDepsJava) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCdkDepsJava) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

