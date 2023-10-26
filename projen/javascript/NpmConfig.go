package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// File representing the local NPM config in .npmrc.
// Experimental.
type NpmConfig interface {
	projen.Component
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// configure a generic property.
	// Experimental.
	AddConfig(name *string, value *string)
	// configure a scoped registry.
	// Experimental.
	AddRegistry(url *string, scope *string)
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

// The jsii proxy struct for NpmConfig
type jsiiProxy_NpmConfig struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_NpmConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NpmConfig) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewNpmConfig(project NodeProject, options *NpmConfigOptions) NpmConfig {
	_init_.Initialize()

	if err := validateNewNpmConfigParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_NpmConfig{}

	_jsii_.Create(
		"projen.javascript.NpmConfig",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNpmConfig_Override(n NpmConfig, project NodeProject, options *NpmConfigOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.NpmConfig",
		[]interface{}{project, options},
		n,
	)
}

// Test whether the given construct is a component.
// Experimental.
func NpmConfig_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNpmConfig_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.NpmConfig",
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
func NpmConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNpmConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.NpmConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NpmConfig) AddConfig(name *string, value *string) {
	if err := n.validateAddConfigParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addConfig",
		[]interface{}{name, value},
	)
}

func (n *jsiiProxy_NpmConfig) AddRegistry(url *string, scope *string) {
	if err := n.validateAddRegistryParameters(url); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addRegistry",
		[]interface{}{url, scope},
	)
}

func (n *jsiiProxy_NpmConfig) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NpmConfig) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NpmConfig) Synthesize() {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NpmConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

