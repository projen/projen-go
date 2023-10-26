package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Experimental.
type TypescriptConfig interface {
	projen.Component
	// Experimental.
	CompilerOptions() *TypeScriptCompilerOptions
	// Experimental.
	Exclude() *[]*string
	// Array of base `tsconfig.json` paths. Any absolute paths are resolved relative to this instance, while any relative paths are used as is.
	// Experimental.
	Extends() *[]*string
	// Experimental.
	File() projen.JsonFile
	// Experimental.
	FileName() *string
	// Experimental.
	Include() *[]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Experimental.
	AddExclude(pattern *string)
	// Extend from base `TypescriptConfig` instance.
	// Experimental.
	AddExtends(value TypescriptConfig)
	// Experimental.
	AddInclude(pattern *string)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Resolve valid TypeScript extends paths relative to this config.
	// Experimental.
	ResolveExtendsPath(configPath *string) *string
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TypescriptConfig
type jsiiProxy_TypescriptConfig struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_TypescriptConfig) CompilerOptions() *TypeScriptCompilerOptions {
	var returns *TypeScriptCompilerOptions
	_jsii_.Get(
		j,
		"compilerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypescriptConfig) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypescriptConfig) Extends() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extends",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypescriptConfig) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypescriptConfig) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypescriptConfig) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypescriptConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypescriptConfig) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewTypescriptConfig(project projen.Project, options *TypescriptConfigOptions) TypescriptConfig {
	_init_.Initialize()

	if err := validateNewTypescriptConfigParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TypescriptConfig{}

	_jsii_.Create(
		"projen.javascript.TypescriptConfig",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTypescriptConfig_Override(t TypescriptConfig, project projen.Project, options *TypescriptConfigOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.TypescriptConfig",
		[]interface{}{project, options},
		t,
	)
}

// Test whether the given construct is a component.
// Experimental.
func TypescriptConfig_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTypescriptConfig_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.TypescriptConfig",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func TypescriptConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTypescriptConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.TypescriptConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypescriptConfig) AddExclude(pattern *string) {
	if err := t.validateAddExcludeParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addExclude",
		[]interface{}{pattern},
	)
}

func (t *jsiiProxy_TypescriptConfig) AddExtends(value TypescriptConfig) {
	if err := t.validateAddExtendsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addExtends",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TypescriptConfig) AddInclude(pattern *string) {
	if err := t.validateAddIncludeParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addInclude",
		[]interface{}{pattern},
	)
}

func (t *jsiiProxy_TypescriptConfig) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TypescriptConfig) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TypescriptConfig) ResolveExtendsPath(configPath *string) *string {
	if err := t.validateResolveExtendsPathParameters(configPath); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"resolveExtendsPath",
		[]interface{}{configPath},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypescriptConfig) Synthesize() {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TypescriptConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

