package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Installs the following npm scripts:.
//
// - `test` will run `jest --passWithNoTests`
// - `test:watch` will run `jest --watch`
// - `test:update` will run `jest -u`.
// Experimental.
type Jest interface {
	// Escape hatch.
	// Experimental.
	Config() interface{}
	// Experimental.
	JestVersion() *string
	// Experimental.
	AddIgnorePattern(pattern *string)
	// Experimental.
	AddReporter(reporter interface{})
	// Experimental.
	AddSnapshotResolver(file *string)
	// Adds a test match pattern.
	// Experimental.
	AddTestMatch(pattern *string)
	// Adds a watch ignore pattern.
	// Experimental.
	AddWatchIgnorePattern(pattern *string)
}

// The jsii proxy struct for Jest
type jsiiProxy_Jest struct {
	_ byte // padding
}

func (j *jsiiProxy_Jest) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Jest) JestVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jestVersion",
		&returns,
	)
	return returns
}


// Experimental.
func NewJest(project NodeProject, options *JestOptions) Jest {
	_init_.Initialize()

	j := jsiiProxy_Jest{}

	_jsii_.Create(
		"projen.javascript.Jest",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewJest_Override(j Jest, project NodeProject, options *JestOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Jest",
		[]interface{}{project, options},
		j,
	)
}

func (j *jsiiProxy_Jest) AddIgnorePattern(pattern *string) {
	_jsii_.InvokeVoid(
		j,
		"addIgnorePattern",
		[]interface{}{pattern},
	)
}

func (j *jsiiProxy_Jest) AddReporter(reporter interface{}) {
	_jsii_.InvokeVoid(
		j,
		"addReporter",
		[]interface{}{reporter},
	)
}

func (j *jsiiProxy_Jest) AddSnapshotResolver(file *string) {
	_jsii_.InvokeVoid(
		j,
		"addSnapshotResolver",
		[]interface{}{file},
	)
}

func (j *jsiiProxy_Jest) AddTestMatch(pattern *string) {
	_jsii_.InvokeVoid(
		j,
		"addTestMatch",
		[]interface{}{pattern},
	)
}

func (j *jsiiProxy_Jest) AddWatchIgnorePattern(pattern *string) {
	_jsii_.InvokeVoid(
		j,
		"addWatchIgnorePattern",
		[]interface{}{pattern},
	)
}

