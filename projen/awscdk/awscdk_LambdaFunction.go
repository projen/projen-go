package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

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
}

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__projenComponent
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

