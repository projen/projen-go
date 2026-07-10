package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Project-level logging utilities.
// Experimental.
type Logger interface {
	Component
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() Project
	// Log a message to stderr with DEBUG severity.
	// Experimental.
	Debug(text ...interface{})
	// Log a message to stderr with ERROR severity.
	// Experimental.
	Error(text ...interface{})
	// Log a message to stderr with INFO severity.
	// Experimental.
	Info(text ...interface{})
	// Log a message to stderr with a given logging level.
	//
	// The message will be
	// printed as long as `logger.level` is set to the message's severity or higher.
	// Experimental.
	Log(level LogLevel, text ...interface{})
	// Called once, right after `postSynthesize()`, only when the project is created for the first time.
	//
	// It does not run on later `projen` invocations. It only fires for `projen new` (or `Projects.createProject`).
	// It is also skipped when post-synthesis steps are disabled, e.g. `--no-post` or `PROJEN_DISABLE_POST`.
	// Use it for one-off setup that can be turned off by the user, like running a task to give the user immediate
	// feedback on their new project. Order across components is not guaranteed.
	// Experimental.
	PostProjectCreation(initProject *InitProject)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Called once, right after `synthesize()`, only when the project is created for the first time.
	//
	// It does not run on later `projen` invocations. It only fires for `projen new` (or `Projects.createProject`).
	// Use it for deterministic, one-off file generation. Order across components is not guaranteed.
	// Experimental.
	ProjectCreation(initProject *InitProject)
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Log a message to stderr with VERBOSE severity.
	// Experimental.
	Verbose(text ...interface{})
	// Log a message to stderr with WARN severity.
	// Experimental.
	Warn(text ...interface{})
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for Logger
type jsiiProxy_Logger struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Logger) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Logger) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewLogger(scope constructs.IConstruct, options *LoggerOptions) Logger {
	_init_.Initialize()

	if err := validateNewLoggerParameters(scope, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Logger{}

	_jsii_.Create(
		"projen.Logger",
		[]interface{}{scope, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLogger_Override(l Logger, scope constructs.IConstruct, options *LoggerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Logger",
		[]interface{}{scope, options},
		l,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Logger_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogger_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Logger",
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
func Logger_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogger_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Logger",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Logger) Debug(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"debug",
		args,
	)
}

func (l *jsiiProxy_Logger) Error(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"error",
		args,
	)
}

func (l *jsiiProxy_Logger) Info(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"info",
		args,
	)
}

func (l *jsiiProxy_Logger) Log(level LogLevel, text ...interface{}) {
	if err := l.validateLogParameters(level); err != nil {
		panic(err)
	}
	args := []interface{}{level}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"log",
		args,
	)
}

func (l *jsiiProxy_Logger) PostProjectCreation(initProject *InitProject) {
	if err := l.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"postProjectCreation",
		[]interface{}{initProject},
	)
}

func (l *jsiiProxy_Logger) PostSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"postSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Logger) PreSynthesize() {
	_jsii_.InvokeVoid(
		l,
		"preSynthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Logger) ProjectCreation(initProject *InitProject) {
	if err := l.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (l *jsiiProxy_Logger) Synthesize() {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Logger) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Logger) Verbose(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"verbose",
		args,
	)
}

func (l *jsiiProxy_Logger) Warn(text ...interface{}) {
	args := []interface{}{}
	for _, a := range text {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		l,
		"warn",
		args,
	)
}

func (l *jsiiProxy_Logger) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		l,
		"with",
		args,
		&returns,
	)

	return returns
}

