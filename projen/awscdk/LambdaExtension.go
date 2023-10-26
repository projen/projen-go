package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
)

// Create a Lambda Extension.
// Experimental.
type LambdaExtension interface {
	projen.Component
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

// The jsii proxy struct for LambdaExtension
type jsiiProxy_LambdaExtension struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_LambdaExtension) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaExtension) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaExtension(project projen.Project, options *LambdaExtensionOptions) LambdaExtension {
	_init_.Initialize()

	if err := validateNewLambdaExtensionParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaExtension{}

	_jsii_.Create(
		"projen.awscdk.LambdaExtension",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaExtension_Override(l LambdaExtension, project projen.Project, options *LambdaExtensionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.LambdaExtension",
		[]interface{}{project, options},
		l,
	)
}

// Test whether the given construct is a component.
// Experimental.
func LambdaExtension_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaExtension_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.LambdaExtension",
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
func LambdaExtension_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaExtension_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.LambdaExtension",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaExtension) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaExtension) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaExtension) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaExtension) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

