package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Configure validations to run on GitHub pull requests.
//
// Only generates a file if at least one linter is configured.
// Experimental.
type PullRequestLint interface {
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

// The jsii proxy struct for PullRequestLint
type jsiiProxy_PullRequestLint struct {
	internal.Type__projenComponent
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

