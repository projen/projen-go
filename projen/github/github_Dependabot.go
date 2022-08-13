package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Defines dependabot configuration for node projects.
//
// Since module versions are managed in projen, the versioning strategy will be
// configured to "lockfile-only" which means that only updates that can be done
// on the lockfile itself will be proposed.
// Experimental.
type Dependabot interface {
	projen.Component
	// The raw dependabot configuration.
	// See: https://docs.github.com/en/github/administering-a-repository/configuration-options-for-dependency-updates
	//
	// Experimental.
	Config() interface{}
	// Whether or not projen is also upgraded in this config,.
	// Experimental.
	IgnoresProjen() *bool
	// Experimental.
	Project() projen.Project
	// Ignores a dependency from automatic updates.
	// Experimental.
	AddIgnore(dependencyName *string, versions ...*string)
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

// The jsii proxy struct for Dependabot
type jsiiProxy_Dependabot struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Dependabot) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dependabot) IgnoresProjen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ignoresProjen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dependabot) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewDependabot(github GitHub, options *DependabotOptions) Dependabot {
	_init_.Initialize()

	j := jsiiProxy_Dependabot{}

	_jsii_.Create(
		"projen.github.Dependabot",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewDependabot_Override(d Dependabot, github GitHub, options *DependabotOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.Dependabot",
		[]interface{}{github, options},
		d,
	)
}

func (d *jsiiProxy_Dependabot) AddIgnore(dependencyName *string, versions ...*string) {
	args := []interface{}{dependencyName}
	for _, a := range versions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addIgnore",
		args,
	)
}

func (d *jsiiProxy_Dependabot) PostSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"postSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dependabot) PreSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"preSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dependabot) Synthesize() {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		nil, // no parameters
	)
}

