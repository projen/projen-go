// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Represents a project component.
// Experimental.
type Component interface {
	// Experimental.
	Project() Project
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

// The jsii proxy struct for Component
type jsiiProxy_Component struct {
	_ byte // padding
}

func (j *jsiiProxy_Component) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewComponent(project Project) Component {
	_init_.Initialize()

	if err := validateNewComponentParameters(project); err != nil {
		panic(err)
	}
	j := jsiiProxy_Component{}

	_jsii_.Create(
		"projen.Component",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewComponent_Override(c Component, project Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Component",
		[]interface{}{project},
		c,
	)
}

func (c *jsiiProxy_Component) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Component) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Component) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

