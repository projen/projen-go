package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
	"github.com/projen/projen-go/projen/cdk"
)

// Creates Lambda Extensions from entrypoints discovered in the project's source tree.
// Experimental.
type LambdaExtensionAutoDiscover interface {
	cdk.AutoDiscoverBase
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

// The jsii proxy struct for LambdaExtensionAutoDiscover
type jsiiProxy_LambdaExtensionAutoDiscover struct {
	internal.Type__cdkAutoDiscoverBase
}

func (j *jsiiProxy_LambdaExtensionAutoDiscover) Entrypoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entrypoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaExtensionAutoDiscover) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaExtensionAutoDiscover) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaExtensionAutoDiscover(project projen.Project, options *LambdaExtensionAutoDiscoverOptions) LambdaExtensionAutoDiscover {
	_init_.Initialize()

	if err := validateNewLambdaExtensionAutoDiscoverParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaExtensionAutoDiscover{}

	_jsii_.Create(
		"projen.awscdk.LambdaExtensionAutoDiscover",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaExtensionAutoDiscover_Override(l LambdaExtensionAutoDiscover, project projen.Project, options *LambdaExtensionAutoDiscoverOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.LambdaExtensionAutoDiscover",
		[]interface{}{project, options},
		l,
	)
}

// Test whether the given construct is a component.
// Experimental.
func LambdaExtensionAutoDiscover_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaExtensionAutoDiscover_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.LambdaExtensionAutoDiscover",
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
func LambdaExtensionAutoDiscover_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaExtensionAutoDiscover_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.LambdaExtensionAutoDiscover",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaExtensionAutoDiscover) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaExtensionAutoDiscover) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaExtensionAutoDiscover) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaExtensionAutoDiscover) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

