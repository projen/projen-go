package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Represents eslint configuration.
// Experimental.
type Eslint interface {
	projen.Component
	// Direct access to the eslint configuration (escape hatch).
	// Experimental.
	Config() interface{}
	// File patterns that should not be linted.
	// Experimental.
	IgnorePatterns() *[]*string
	// eslint overrides.
	// Experimental.
	Overrides() *[]*EslintOverride
	// Experimental.
	Project() projen.Project
	// eslint rules.
	// Experimental.
	Rules() *map[string]*[]interface{}
	// Adds an `extends` item to the eslint configuration.
	// Experimental.
	AddExtends(extendList ...*string)
	// Do not lint these files.
	// Experimental.
	AddIgnorePattern(pattern *string)
	// Add an eslint override.
	// Experimental.
	AddOverride(override *EslintOverride)
	// Adds an eslint plugin.
	// Experimental.
	AddPlugins(plugins ...*string)
	// Add an eslint rule.
	// Experimental.
	AddRules(rules *map[string]interface{})
	// Add a glob file pattern which allows importing dev dependencies.
	// Experimental.
	AllowDevDeps(pattern *string)
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

// The jsii proxy struct for Eslint
type jsiiProxy_Eslint struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Eslint) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eslint) IgnorePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignorePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eslint) Overrides() *[]*EslintOverride {
	var returns *[]*EslintOverride
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eslint) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eslint) Rules() *map[string]*[]interface{} {
	var returns *map[string]*[]interface{}
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}


// Experimental.
func NewEslint(project NodeProject, options *EslintOptions) Eslint {
	_init_.Initialize()

	j := jsiiProxy_Eslint{}

	_jsii_.Create(
		"projen.javascript.Eslint",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewEslint_Override(e Eslint, project NodeProject, options *EslintOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Eslint",
		[]interface{}{project, options},
		e,
	)
}

// Returns the singletone Eslint component of a project or undefined if there is none.
// Experimental.
func Eslint_Of(project projen.Project) Eslint {
	_init_.Initialize()

	var returns Eslint

	_jsii_.StaticInvoke(
		"projen.javascript.Eslint",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eslint) AddExtends(extendList ...*string) {
	args := []interface{}{}
	for _, a := range extendList {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addExtends",
		args,
	)
}

func (e *jsiiProxy_Eslint) AddIgnorePattern(pattern *string) {
	_jsii_.InvokeVoid(
		e,
		"addIgnorePattern",
		[]interface{}{pattern},
	)
}

func (e *jsiiProxy_Eslint) AddOverride(override *EslintOverride) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{override},
	)
}

func (e *jsiiProxy_Eslint) AddPlugins(plugins ...*string) {
	args := []interface{}{}
	for _, a := range plugins {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addPlugins",
		args,
	)
}

func (e *jsiiProxy_Eslint) AddRules(rules *map[string]interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addRules",
		[]interface{}{rules},
	)
}

func (e *jsiiProxy_Eslint) AllowDevDeps(pattern *string) {
	_jsii_.InvokeVoid(
		e,
		"allowDevDeps",
		[]interface{}{pattern},
	)
}

func (e *jsiiProxy_Eslint) PostSynthesize() {
	_jsii_.InvokeVoid(
		e,
		"postSynthesize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eslint) PreSynthesize() {
	_jsii_.InvokeVoid(
		e,
		"preSynthesize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eslint) Synthesize() {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		nil, // no parameters
	)
}

