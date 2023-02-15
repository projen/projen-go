package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Warns and then closes issues and PRs that have had no activity for a specified amount of time.
//
// The default configuration will:
//
//   * Add a "Stale" label to pull requests after 14 days and closed after 2 days
//   * Add a "Stale" label to issues after 60 days and closed after 7 days
//   * If a comment is added, the label will be removed and timer is restarted.
// See: https://github.com/actions/stale
//
// Experimental.
type Stale interface {
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

// The jsii proxy struct for Stale
type jsiiProxy_Stale struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Stale) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewStale(github GitHub, options *StaleOptions) Stale {
	_init_.Initialize()

	if err := validateNewStaleParameters(github, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Stale{}

	_jsii_.Create(
		"projen.github.Stale",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewStale_Override(s Stale, github GitHub, options *StaleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.Stale",
		[]interface{}{github, options},
		s,
	)
}

func (s *jsiiProxy_Stale) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stale) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stale) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

