package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Installs the following npm scripts:.
//
// - `test`, intended for testing locally and in CI. Will update snapshots unless `updateSnapshot: UpdateSnapshot: NEVER` is set.
// - `test:watch`, intended for automatically rerunning tests when files change.
// - `test:update`, intended for testing locally and updating snapshots to match the latest unit under test. Only available when `updateSnapshot: UpdateSnapshot: NEVER`.
// Experimental.
type Jest interface {
	projen.Component
	// Escape hatch.
	// Experimental.
	Config() interface{}
	// Jest config file.
	//
	// `undefined` if settings are written to `package.json`
	// Experimental.
	File() projen.JsonFile
	// Jest version, including `@` symbol, like `@^29`.
	// Experimental.
	JestVersion() *string
	// Experimental.
	Project() projen.Project
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

// The jsii proxy struct for Jest
type jsiiProxy_Jest struct {
	internal.Type__projenComponent
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

func (j *jsiiProxy_Jest) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
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

func (j *jsiiProxy_Jest) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewJest(project NodeProject, options *JestOptions) Jest {
	_init_.Initialize()

	if err := validateNewJestParameters(project, options); err != nil {
		panic(err)
	}
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

// Returns the singletone Jest component of a project or undefined if there is none.
// Experimental.
func Jest_Of(project projen.Project) Jest {
	_init_.Initialize()

	if err := validateJest_OfParameters(project); err != nil {
		panic(err)
	}
	var returns Jest

	_jsii_.StaticInvoke(
		"projen.javascript.Jest",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Jest) AddIgnorePattern(pattern *string) {
	if err := j.validateAddIgnorePatternParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addIgnorePattern",
		[]interface{}{pattern},
	)
}

func (j *jsiiProxy_Jest) AddReporter(reporter interface{}) {
	if err := j.validateAddReporterParameters(reporter); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addReporter",
		[]interface{}{reporter},
	)
}

func (j *jsiiProxy_Jest) AddSnapshotResolver(file *string) {
	if err := j.validateAddSnapshotResolverParameters(file); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addSnapshotResolver",
		[]interface{}{file},
	)
}

func (j *jsiiProxy_Jest) AddTestMatch(pattern *string) {
	if err := j.validateAddTestMatchParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addTestMatch",
		[]interface{}{pattern},
	)
}

func (j *jsiiProxy_Jest) AddWatchIgnorePattern(pattern *string) {
	if err := j.validateAddWatchIgnorePatternParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addWatchIgnorePattern",
		[]interface{}{pattern},
	)
}

func (j *jsiiProxy_Jest) PostSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"postSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Jest) PreSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"preSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Jest) Synthesize() {
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		nil, // no parameters
	)
}

