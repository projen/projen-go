package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Experimental.
	AddIgnorePattern(pattern *string)
	// Experimental.
	AddReporter(reporter JestReporter)
	// Adds a a setup file to Jest's setupFiles configuration.
	// Experimental.
	AddSetupFile(file *string)
	// Adds a a setup file to Jest's setupFilesAfterEnv configuration.
	// Experimental.
	AddSetupFileAfterEnv(file *string)
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
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

func (j *jsiiProxy_Jest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

// Test whether the given construct is a component.
// Experimental.
func Jest_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJest_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Jest",
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
func Jest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Jest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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

func (j *jsiiProxy_Jest) AddReporter(reporter JestReporter) {
	if err := j.validateAddReporterParameters(reporter); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addReporter",
		[]interface{}{reporter},
	)
}

func (j *jsiiProxy_Jest) AddSetupFile(file *string) {
	if err := j.validateAddSetupFileParameters(file); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addSetupFile",
		[]interface{}{file},
	)
}

func (j *jsiiProxy_Jest) AddSetupFileAfterEnv(file *string) {
	if err := j.validateAddSetupFileAfterEnvParameters(file); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addSetupFileAfterEnv",
		[]interface{}{file},
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

func (j *jsiiProxy_Jest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

