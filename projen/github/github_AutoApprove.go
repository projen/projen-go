package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Auto approve pull requests that meet a criteria.
// Experimental.
type AutoApprove interface {
	projen.Component
	// Experimental.
	Label() *string
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

// The jsii proxy struct for AutoApprove
type jsiiProxy_AutoApprove struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_AutoApprove) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoApprove) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAutoApprove(github GitHub, options *AutoApproveOptions) AutoApprove {
	_init_.Initialize()

	j := jsiiProxy_AutoApprove{}

	_jsii_.Create(
		"projen.github.AutoApprove",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAutoApprove_Override(a AutoApprove, github GitHub, options *AutoApproveOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.AutoApprove",
		[]interface{}{github, options},
		a,
	)
}

func (a *jsiiProxy_AutoApprove) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoApprove) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoApprove) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

