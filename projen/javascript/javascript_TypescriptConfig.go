package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
)

// Experimental.
type TypescriptConfig interface {
	// Experimental.
	CompilerOptions() *TypeScriptCompilerOptions
	// Experimental.
	Exclude() *[]*string
	// Experimental.
	File() projen.JsonFile
	// Experimental.
	FileName() *string
	// Experimental.
	Include() *[]*string
	// Experimental.
	AddExclude(pattern *string)
	// Experimental.
	AddInclude(pattern *string)
}

// The jsii proxy struct for TypescriptConfig
type jsiiProxy_TypescriptConfig struct {
	_ byte // padding
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


// Experimental.
func NewTypescriptConfig(project NodeProject, options *TypescriptConfigOptions) TypescriptConfig {
	_init_.Initialize()

	j := jsiiProxy_TypescriptConfig{}

	_jsii_.Create(
		"projen.javascript.TypescriptConfig",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTypescriptConfig_Override(t TypescriptConfig, project NodeProject, options *TypescriptConfigOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.TypescriptConfig",
		[]interface{}{project, options},
		t,
	)
}

func (t *jsiiProxy_TypescriptConfig) AddExclude(pattern *string) {
	_jsii_.InvokeVoid(
		t,
		"addExclude",
		[]interface{}{pattern},
	)
}

func (t *jsiiProxy_TypescriptConfig) AddInclude(pattern *string) {
	_jsii_.InvokeVoid(
		t,
		"addInclude",
		[]interface{}{pattern},
	)
}

