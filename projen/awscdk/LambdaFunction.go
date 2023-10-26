package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
)

// Generates a pre-bundled AWS Lambda function construct from handler code.
//
// To use this, create an AWS Lambda handler file under your source tree with
// the `.lambda.ts` extension and add a `LambdaFunction` component to your
// typescript project pointing to this entrypoint.
//
// This will add a task to your "compile" step which will use `esbuild` to
// bundle the handler code into the build directory. It will also generate a
// file `src/foo-function.ts` with a custom AWS construct called `FooFunction`
// which extends `@aws-cdk/aws-lambda.Function` which is bound to the bundled
// handle through an asset.
//
// Example:
//   new LambdaFunction(myProject, {
//     srcdir: myProject.srcdir,
//     entrypoint: 'src/foo.lambda.ts',
//   });
//
// Experimental.
type LambdaFunction interface {
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

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_LambdaFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Defines a pre-bundled AWS Lambda function construct from handler code.
// Experimental.
func NewLambdaFunction(project projen.Project, options *LambdaFunctionOptions) LambdaFunction {
	_init_.Initialize()

	if err := validateNewLambdaFunctionParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunction{}

	_jsii_.Create(
		"projen.awscdk.LambdaFunction",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Defines a pre-bundled AWS Lambda function construct from handler code.
// Experimental.
func NewLambdaFunction_Override(l LambdaFunction, project projen.Project, options *LambdaFunctionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.LambdaFunction",
		[]interface{}{project, options},
		l,
	)
}

// Test whether the given construct is a component.
// Experimental.
func LambdaFunction_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaFunction_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.LambdaFunction",
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
func LambdaFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.LambdaFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

