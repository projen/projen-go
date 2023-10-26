package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Configure validations to run on GitHub pull requests.
//
// Only generates a file if at least one linter is configured.
// Experimental.
type PullRequestLint interface {
	projen.Component
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PullRequestLint
type jsiiProxy_PullRequestLint struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_PullRequestLint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestLint) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPullRequestLint(github GitHub, options *PullRequestLintOptions) PullRequestLint {
	_init_.Initialize()

	if err := validateNewPullRequestLintParameters(github, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PullRequestLint{}

	_jsii_.Create(
		"projen.github.PullRequestLint",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPullRequestLint_Override(p PullRequestLint, github GitHub, options *PullRequestLintOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.PullRequestLint",
		[]interface{}{github, options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func PullRequestLint_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePullRequestLint_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.PullRequestLint",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func PullRequestLint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePullRequestLint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.PullRequestLint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PullRequestLint) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestLint) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestLint) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestLint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

