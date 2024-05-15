package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Experimental.
type PullRequestBackport interface {
	projen.Component
	// Experimental.
	File() projen.JsonFile
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Experimental.
	Workflow() GithubWorkflow
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

// The jsii proxy struct for PullRequestBackport
type jsiiProxy_PullRequestBackport struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_PullRequestBackport) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestBackport) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestBackport) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestBackport) Workflow() GithubWorkflow {
	var returns GithubWorkflow
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}


// Experimental.
func NewPullRequestBackport(scope constructs.IConstruct, options *PullRequestBackportOptions) PullRequestBackport {
	_init_.Initialize()

	if err := validateNewPullRequestBackportParameters(scope, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PullRequestBackport{}

	_jsii_.Create(
		"projen.github.PullRequestBackport",
		[]interface{}{scope, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPullRequestBackport_Override(p PullRequestBackport, scope constructs.IConstruct, options *PullRequestBackportOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.PullRequestBackport",
		[]interface{}{scope, options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func PullRequestBackport_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePullRequestBackport_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.PullRequestBackport",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func PullRequestBackport_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePullRequestBackport_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.PullRequestBackport",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PullRequestBackport) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestBackport) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestBackport) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestBackport) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

