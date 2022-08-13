package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
)

// Discovers and creates integration tests and lambdas from code in the project's source and test trees.
// Experimental.
type AutoDiscover interface {
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

// The jsii proxy struct for AutoDiscover
type jsiiProxy_AutoDiscover struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_AutoDiscover) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAutoDiscover(project projen.Project, options *AutoDiscoverOptions) AutoDiscover {
	_init_.Initialize()

	j := jsiiProxy_AutoDiscover{}

	_jsii_.Create(
		"projen.awscdk.AutoDiscover",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAutoDiscover_Override(a AutoDiscover, project projen.Project, options *AutoDiscoverOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.AutoDiscover",
		[]interface{}{project, options},
		a,
	)
}

func (a *jsiiProxy_AutoDiscover) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoDiscover) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoDiscover) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

