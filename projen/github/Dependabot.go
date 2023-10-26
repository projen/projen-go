package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Defines dependabot configuration for node projects.
//
// Since module versions are managed in projen, the versioning strategy will be
// configured to "lockfile-only" which means that only updates that can be done
// on the lockfile itself will be proposed.
// Experimental.
type Dependabot interface {
	projen.Component
	// The raw dependabot configuration.
	// See: https://docs.github.com/en/github/administering-a-repository/configuration-options-for-dependency-updates
	//
	// Experimental.
	Config() interface{}
	// Whether or not projen is also upgraded in this config,.
	// Experimental.
	IgnoresProjen() *bool
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Allows a dependency from automatic updates.
	// Experimental.
	AddAllow(dependencyName *string)
	// Ignores a dependency from automatic updates.
	// Experimental.
	AddIgnore(dependencyName *string, versions ...*string)
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

// The jsii proxy struct for Dependabot
type jsiiProxy_Dependabot struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Dependabot) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dependabot) IgnoresProjen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ignoresProjen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dependabot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dependabot) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewDependabot(github GitHub, options *DependabotOptions) Dependabot {
	_init_.Initialize()

	if err := validateNewDependabotParameters(github, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Dependabot{}

	_jsii_.Create(
		"projen.github.Dependabot",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewDependabot_Override(d Dependabot, github GitHub, options *DependabotOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.Dependabot",
		[]interface{}{github, options},
		d,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Dependabot_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDependabot_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.Dependabot",
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
func Dependabot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDependabot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.Dependabot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dependabot) AddAllow(dependencyName *string) {
	if err := d.validateAddAllowParameters(dependencyName); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addAllow",
		[]interface{}{dependencyName},
	)
}

func (d *jsiiProxy_Dependabot) AddIgnore(dependencyName *string, versions ...*string) {
	if err := d.validateAddIgnoreParameters(dependencyName); err != nil {
		panic(err)
	}
	args := []interface{}{dependencyName}
	for _, a := range versions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addIgnore",
		args,
	)
}

func (d *jsiiProxy_Dependabot) PostSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"postSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dependabot) PreSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"preSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dependabot) Synthesize() {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dependabot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

