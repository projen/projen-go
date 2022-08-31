// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
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
	// The file path, relative to the project root.
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
func NewTextFile(project Project, filePath *string, options *TextFileOptions) TextFile {
	_init_.Initialize()

	if err := validateNewTextFileParameters(project, filePath, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TextFile{}

	_jsii_.Create(
		"projen.TextFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Defines a text file.
// Experimental.
func NewTextFile_Override(t TextFile, project Project, filePath *string, options *TextFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.TextFile",
		[]interface{}{project, filePath, options},
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

