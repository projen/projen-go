package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
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
	// eslint task.
	// Experimental.
	EslintTask() projen.Task
	// File patterns that should not be linted.
	// Experimental.
	IgnorePatterns() *[]*string
	// Returns an immutable copy of the lintPatterns being used by this eslint configuration.
	// Experimental.
	LintPatterns() *[]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Add a file, glob pattern or directory with source files to lint (e.g. [ "src" ]).
	// Experimental.
	AddLintPattern(pattern *string)
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
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

func (j *jsiiProxy_Eslint) EslintTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"eslintTask",
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

func (j *jsiiProxy_Eslint) LintPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lintPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eslint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

	if err := validateNewEslintParameters(project, options); err != nil {
		panic(err)
	}
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

// Test whether the given construct is a component.
// Experimental.
func Eslint_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEslint_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Eslint",
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
func Eslint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEslint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Eslint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns the singleton Eslint component of a project or undefined if there is none.
// Experimental.
func Eslint_Of(project projen.Project) Eslint {
	_init_.Initialize()

	if err := validateEslint_OfParameters(project); err != nil {
		panic(err)
	}
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
	if err := e.validateAddIgnorePatternParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addIgnorePattern",
		[]interface{}{pattern},
	)
}

func (e *jsiiProxy_Eslint) AddLintPattern(pattern *string) {
	if err := e.validateAddLintPatternParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addLintPattern",
		[]interface{}{pattern},
	)
}

func (e *jsiiProxy_Eslint) AddOverride(override *EslintOverride) {
	if err := e.validateAddOverrideParameters(override); err != nil {
		panic(err)
	}
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
	if err := e.validateAddRulesParameters(rules); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addRules",
		[]interface{}{rules},
	)
}

func (e *jsiiProxy_Eslint) AllowDevDeps(pattern *string) {
	if err := e.validateAllowDevDepsParameters(pattern); err != nil {
		panic(err)
	}
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

func (e *jsiiProxy_Eslint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

