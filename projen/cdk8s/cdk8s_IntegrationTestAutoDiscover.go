package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/cdk"
	"github.com/projen/projen-go/projen/cdk8s/internal"
)

// Discovers and creates integration tests from files in the test root.
// Experimental.
type IntegrationTestAutoDiscover interface {
	cdk.IntegrationTestAutoDiscoverBase
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

// The jsii proxy struct for IntegrationTestAutoDiscover
type jsiiProxy_IntegrationTestAutoDiscover struct {
	internal.Type__cdkIntegrationTestAutoDiscoverBase
}

func (j *jsiiProxy_IntegrationTestAutoDiscover) Entrypoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entrypoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestAutoDiscover) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegrationTestAutoDiscover(project projen.Project, options *IntegrationTestAutoDiscoverOptions) IntegrationTestAutoDiscover {
	_init_.Initialize()

	j := jsiiProxy_IntegrationTestAutoDiscover{}

	_jsii_.Create(
		"projen.cdk8s.IntegrationTestAutoDiscover",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegrationTestAutoDiscover_Override(i IntegrationTestAutoDiscover, project projen.Project, options *IntegrationTestAutoDiscoverOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdk8s.IntegrationTestAutoDiscover",
		[]interface{}{project, options},
		i,
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscover) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscover) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscover) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

