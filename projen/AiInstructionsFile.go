package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type AiInstructionsFile interface {
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
	// Adds instructions to the instruction file.
	// Experimental.
	AddInstructions(instructions ...*string)
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
	SynthesizeContent(resolver IResolver) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiInstructionsFile
type jsiiProxy_AiInstructionsFile struct {
	jsiiProxy_FileBase
}

func (j *jsiiProxy_AiInstructionsFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiInstructionsFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiInstructionsFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiInstructionsFile) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiInstructionsFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiInstructionsFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiInstructionsFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiInstructionsFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewAiInstructionsFile(scope constructs.IConstruct, filePath *string, options *FileBaseOptions) AiInstructionsFile {
	_init_.Initialize()

	if err := validateNewAiInstructionsFileParameters(scope, filePath, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiInstructionsFile{}

	_jsii_.Create(
		"projen.AiInstructionsFile",
		[]interface{}{scope, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAiInstructionsFile_Override(a AiInstructionsFile, scope constructs.IConstruct, filePath *string, options *FileBaseOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.AiInstructionsFile",
		[]interface{}{scope, filePath, options},
		a,
	)
}

func (j *jsiiProxy_AiInstructionsFile)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_AiInstructionsFile)SetReadonly(val *bool) {
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
func AiInstructionsFile_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiInstructionsFile_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.AiInstructionsFile",
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
func AiInstructionsFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiInstructionsFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.AiInstructionsFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiInstructionsFile) AddInstructions(instructions ...*string) {
	args := []interface{}{}
	for _, a := range instructions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addInstructions",
		args,
	)
}

func (a *jsiiProxy_AiInstructionsFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiInstructionsFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiInstructionsFile) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiInstructionsFile) SynthesizeContent(resolver IResolver) *string {
	if err := a.validateSynthesizeContentParameters(resolver); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiInstructionsFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

