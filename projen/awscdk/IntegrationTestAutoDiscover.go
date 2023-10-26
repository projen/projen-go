package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
	"github.com/projen/projen-go/projen/cdk"
)

// Creates integration tests from entry points discovered in the test tree.
// Experimental.
type IntegrationTestAutoDiscover interface {
	cdk.IntegrationTestAutoDiscoverBase
	// Auto-discovered entry points with paths relative to the project directory.
	// Experimental.
	Entrypoints() *[]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
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

// The jsii proxy struct for IntegrationTestAutoDiscover
type jsiiProxy_IntegrationTestAutoDiscover struct {
	internal.Type__cdkIntegrationTestAutoDiscoverBase
}

func (j *jsiiProxy_IntegrationTestAutoDiscover) Entrypoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entrypoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestAutoDiscover) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestAutoDiscover) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegrationTestAutoDiscover(project projen.Project, options *IntegrationTestAutoDiscoverOptions) IntegrationTestAutoDiscover {
	_init_.Initialize()

	if err := validateNewIntegrationTestAutoDiscoverParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationTestAutoDiscover{}

	_jsii_.Create(
		"projen.awscdk.IntegrationTestAutoDiscover",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegrationTestAutoDiscover_Override(i IntegrationTestAutoDiscover, project projen.Project, options *IntegrationTestAutoDiscoverOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.IntegrationTestAutoDiscover",
		[]interface{}{project, options},
		i,
	)
}

// Test whether the given construct is a component.
// Experimental.
func IntegrationTestAutoDiscover_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationTestAutoDiscover_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.IntegrationTestAutoDiscover",
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
func IntegrationTestAutoDiscover_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationTestAutoDiscover_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.IntegrationTestAutoDiscover",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationTestAutoDiscover) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscover) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscover) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscover) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

