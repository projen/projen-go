// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Defines renovatebot configuration for projen project.
//
// Ignores the versions controlled by Projen.
// Experimental.
type Renovatebot interface {
	Component
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

// The jsii proxy struct for Renovatebot
type jsiiProxy_Renovatebot struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Renovatebot) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewRenovatebot(project Project, options *RenovatebotOptions) Renovatebot {
	_init_.Initialize()

	j := jsiiProxy_Renovatebot{}

	_jsii_.Create(
		"projen.Renovatebot",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewRenovatebot_Override(r Renovatebot, project Project, options *RenovatebotOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Renovatebot",
		[]interface{}{project, options},
		r,
	)
}

func (r *jsiiProxy_Renovatebot) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Renovatebot) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Renovatebot) Synthesize() {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		nil, // no parameters
	)
}

