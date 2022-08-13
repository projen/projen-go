package cdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/cdk/internal"
)

// Base class for auto-discovering and creating project subcomponents.
// Experimental.
type AutoDiscoverBase interface {
	projen.Component
	// Auto-discovered entry points with paths relative to the project directory.
	// Experimental.
	Entrypoints() *[]*string
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

// The jsii proxy struct for AutoDiscoverBase
type jsiiProxy_AutoDiscoverBase struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_AutoDiscoverBase) Entrypoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entrypoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoDiscoverBase) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAutoDiscoverBase_Override(a AutoDiscoverBase, project projen.Project, options *AutoDiscoverBaseOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdk.AutoDiscoverBase",
		[]interface{}{project, options},
		a,
	)
}

func (a *jsiiProxy_AutoDiscoverBase) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoDiscoverBase) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoDiscoverBase) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

