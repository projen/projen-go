package web

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/web/internal"
)

// Experimental.
type NextComponent interface {
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

// The jsii proxy struct for NextComponent
type jsiiProxy_NextComponent struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_NextComponent) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewNextComponent(project javascript.NodeProject, options *NextComponentOptions) NextComponent {
	_init_.Initialize()

	if err := validateNewNextComponentParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_NextComponent{}

	_jsii_.Create(
		"projen.web.NextComponent",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNextComponent_Override(n NextComponent, project javascript.NodeProject, options *NextComponentOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.NextComponent",
		[]interface{}{project, options},
		n,
	)
}

func (n *jsiiProxy_NextComponent) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NextComponent) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NextComponent) Synthesize() {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		nil, // no parameters
	)
}

