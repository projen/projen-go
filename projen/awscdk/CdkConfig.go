package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
)

// Represents cdk.json file.
// Experimental.
type CdkConfig interface {
	projen.Component
	// Name of the cdk.out directory.
	// Experimental.
	Cdkout() *string
	// List of glob patterns to be excluded by CDK.
	// Experimental.
	Exclude() *[]*string
	// List of glob patterns to be included by CDK.
	// Experimental.
	Include() *[]*string
	// Represents the JSON file.
	// Experimental.
	Json() projen.JsonFile
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Add excludes to `cdk.json`.
	// Experimental.
	AddExcludes(patterns ...*string)
	// Add includes to `cdk.json`.
	// Experimental.
	AddIncludes(patterns ...*string)
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

// The jsii proxy struct for CdkConfig
type jsiiProxy_CdkConfig struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_CdkConfig) Cdkout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdkout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Json() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkConfig) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewCdkConfig(project projen.Project, options *CdkConfigOptions) CdkConfig {
	_init_.Initialize()

	if err := validateNewCdkConfigParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdkConfig{}

	_jsii_.Create(
		"projen.awscdk.CdkConfig",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewCdkConfig_Override(c CdkConfig, project projen.Project, options *CdkConfigOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.CdkConfig",
		[]interface{}{project, options},
		c,
	)
}

// Test whether the given construct is a component.
// Experimental.
func CdkConfig_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdkConfig_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.CdkConfig",
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
func CdkConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdkConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.CdkConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkConfig) AddExcludes(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addExcludes",
		args,
	)
}

func (c *jsiiProxy_CdkConfig) AddIncludes(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addIncludes",
		args,
	)
}

func (c *jsiiProxy_CdkConfig) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkConfig) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkConfig) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

