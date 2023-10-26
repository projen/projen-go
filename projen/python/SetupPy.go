package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Python packaging script where package metadata can be placed.
// Experimental.
type SetupPy interface {
	projen.FileBase
	// The absolute path of this file.
	// Experimental.
	AbsolutePath() *string
	// Indicates if the file has been changed during synthesis.
	//
	// This property is
	// only available in `postSynthesize()` hooks. If this is `undefined`, the
	// file has not been synthesized yet.
	// Experimental.
	Changed() *bool
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable() *bool
	// Experimental.
	SetExecutable(val *bool)
	// The projen marker, used to identify files as projen-generated.
	//
	// Value is undefined if the project is being ejected.
	// Experimental.
	Marker() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The file path, relative to the project's outdir.
	// Experimental.
	Path() *string
	// Experimental.
	Project() projen.Project
	// Indicates if the file should be read-only or read-write.
	// Experimental.
	Readonly() *bool
	// Experimental.
	SetReadonly(val *bool)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Writes the file to the project's output directory.
	// Experimental.
	Synthesize()
	// Implemented by derived classes and returns the contents of the file to emit.
	// Experimental.
	SynthesizeContent(resolver projen.IResolver) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SetupPy
type jsiiProxy_SetupPy struct {
	internal.Type__projenFileBase
}

func (j *jsiiProxy_SetupPy) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SetupPy) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewSetupPy(project projen.Project, options *SetupPyOptions) SetupPy {
	_init_.Initialize()

	if err := validateNewSetupPyParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SetupPy{}

	_jsii_.Create(
		"projen.python.SetupPy",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewSetupPy_Override(s SetupPy, project projen.Project, options *SetupPyOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.SetupPy",
		[]interface{}{project, options},
		s,
	)
}

func (j *jsiiProxy_SetupPy)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_SetupPy)SetReadonly(val *bool) {
	if err := j.validateSetReadonlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

// Test whether the given construct is a component.
// Experimental.
func SetupPy_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSetupPy_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.SetupPy",
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
func SetupPy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSetupPy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.SetupPy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SetupPy) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SetupPy) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SetupPy) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SetupPy) SynthesizeContent(resolver projen.IResolver) *string {
	if err := s.validateSynthesizeContentParameters(resolver); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SetupPy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

