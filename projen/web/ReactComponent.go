package web

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/web/internal"
)

// Experimental.
type ReactComponent interface {
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

// The jsii proxy struct for ReactComponent
type jsiiProxy_ReactComponent struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_ReactComponent) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewReactComponent(project javascript.NodeProject, options *ReactComponentOptions) ReactComponent {
	_init_.Initialize()

	if err := validateNewReactComponentParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReactComponent{}

	_jsii_.Create(
		"projen.web.ReactComponent",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewReactComponent_Override(r ReactComponent, project javascript.NodeProject, options *ReactComponentOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.ReactComponent",
		[]interface{}{project, options},
		r,
	)
}

func (r *jsiiProxy_ReactComponent) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReactComponent) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReactComponent) Synthesize() {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		nil, // no parameters
	)
}

