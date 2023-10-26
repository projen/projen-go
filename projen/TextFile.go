package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A text file.
// Experimental.
type TextFile interface {
	FileBase
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
	// Adds a line to the text file.
	// Experimental.
	AddLine(line *string)
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
	SynthesizeContent(_arg IResolver) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TextFile
type jsiiProxy_TextFile struct {
	jsiiProxy_FileBase
}

func (j *jsiiProxy_TextFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Defines a text file.
// Experimental.
func NewTextFile(scope constructs.IConstruct, filePath *string, options *TextFileOptions) TextFile {
	_init_.Initialize()

	if err := validateNewTextFileParameters(scope, filePath, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TextFile{}

	_jsii_.Create(
		"projen.TextFile",
		[]interface{}{scope, filePath, options},
		&j,
	)

	return &j
}

// Defines a text file.
// Experimental.
func NewTextFile_Override(t TextFile, scope constructs.IConstruct, filePath *string, options *TextFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.TextFile",
		[]interface{}{scope, filePath, options},
		t,
	)
}

func (j *jsiiProxy_TextFile)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_TextFile)SetReadonly(val *bool) {
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
func TextFile_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTextFile_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.TextFile",
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
func TextFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTextFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.TextFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TextFile) AddLine(line *string) {
	if err := t.validateAddLineParameters(line); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addLine",
		[]interface{}{line},
	)
}

func (t *jsiiProxy_TextFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TextFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TextFile) Synthesize() {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TextFile) SynthesizeContent(_arg IResolver) *string {
	if err := t.validateSynthesizeContentParameters(_arg); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"synthesizeContent",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TextFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

