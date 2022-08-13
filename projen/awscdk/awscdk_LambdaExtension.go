package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
)

// Create a Lambda Extension.
// Experimental.
type LambdaExtension interface {
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

// The jsii proxy struct for LambdaExtension
type jsiiProxy_LambdaExtension struct {
	internal.Type__projenComponent
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

