package cdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
)

// Base class for locating integration tests in the project's test tree.
// Experimental.
type IntegrationTestAutoDiscoverBase interface {
	AutoDiscoverBase
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

// The jsii proxy struct for IntegrationTestAutoDiscoverBase
type jsiiProxy_IntegrationTestAutoDiscoverBase struct {
	jsiiProxy_AutoDiscoverBase
}

func (j *jsiiProxy_IntegrationTestAutoDiscoverBase) Entrypoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entrypoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestAutoDiscoverBase) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegrationTestAutoDiscoverBase(project projen.Project, options *IntegrationTestAutoDiscoverBaseOptions) IntegrationTestAutoDiscoverBase {
	_init_.Initialize()

	if err := validateNewIntegrationTestAutoDiscoverBaseParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationTestAutoDiscoverBase{}

	_jsii_.Create(
		"projen.cdk.IntegrationTestAutoDiscoverBase",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegrationTestAutoDiscoverBase_Override(i IntegrationTestAutoDiscoverBase, project projen.Project, options *IntegrationTestAutoDiscoverBaseOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdk.IntegrationTestAutoDiscoverBase",
		[]interface{}{project, options},
		i,
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscoverBase) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscoverBase) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscoverBase) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

