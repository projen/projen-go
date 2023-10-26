package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type FileBase interface {
	Component
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
	Project() Project
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
	//
	// Returns: the content to synthesize or undefined to skip the file.
	// Experimental.
	SynthesizeContent(resolver IResolver) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FileBase
type jsiiProxy_FileBase struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_FileBase) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileBase) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewFileBase_Override(f FileBase, scope constructs.IConstruct, filePath *string, options *FileBaseOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.FileBase",
		[]interface{}{scope, filePath, options},
		f,
	)
}

func (j *jsiiProxy_FileBase)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_FileBase)SetReadonly(val *bool) {
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
func FileBase_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFileBase_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.FileBase",
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
func FileBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFileBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.FileBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileBase) PostSynthesize() {
	_jsii_.InvokeVoid(
		f,
		"postSynthesize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileBase) PreSynthesize() {
	_jsii_.InvokeVoid(
		f,
		"preSynthesize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileBase) Synthesize() {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileBase) SynthesizeContent(resolver IResolver) *string {
	if err := f.validateSynthesizeContentParameters(resolver); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

