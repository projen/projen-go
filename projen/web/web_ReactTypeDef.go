package web

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/web/internal"
)

// Deprecated: No longer used.
type ReactTypeDef interface {
	projen.FileBase
	// The absolute path of this file.
	// Deprecated: No longer used.
	AbsolutePath() *string
	// Indicates if the file has been changed during synthesis.
	//
	// This property is
	// only available in `postSynthesize()` hooks. If this is `undefined`, the
	// file has not been synthesized yet.
	// Deprecated: No longer used.
	Changed() *bool
	// Indicates if the file should be marked as executable.
	// Deprecated: No longer used.
	Executable() *bool
	// Deprecated: No longer used.
	SetExecutable(val *bool)
	// The projen marker, used to identify files as projen-generated.
	//
	// Value is undefined if the project is being ejected.
	// Deprecated: No longer used.
	Marker() *string
	// The file path, relative to the project root.
	// Deprecated: No longer used.
	Path() *string
	// Deprecated: No longer used.
	Project() projen.Project
	// Indicates if the file should be read-only or read-write.
	// Deprecated: No longer used.
	Readonly() *bool
	// Deprecated: No longer used.
	SetReadonly(val *bool)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Deprecated: No longer used.
	PostSynthesize()
	// Called before synthesis.
	// Deprecated: No longer used.
	PreSynthesize()
	// Writes the file to the project's output directory.
	// Deprecated: No longer used.
	Synthesize()
	// Implemented by derived classes and returns the contents of the file to emit.
	// Deprecated: No longer used.
	SynthesizeContent(_arg projen.IResolver) *string
}

// The jsii proxy struct for ReactTypeDef
type jsiiProxy_ReactTypeDef struct {
	internal.Type__projenFileBase
}

func (j *jsiiProxy_ReactTypeDef) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeDef) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeDef) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeDef) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeDef) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeDef) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReactTypeDef) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Deprecated: No longer used.
func NewReactTypeDef(project ReactTypeScriptProject, filePath *string, options *ReactTypeDefOptions) ReactTypeDef {
	_init_.Initialize()

	j := jsiiProxy_ReactTypeDef{}

	_jsii_.Create(
		"projen.web.ReactTypeDef",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Deprecated: No longer used.
func NewReactTypeDef_Override(r ReactTypeDef, project ReactTypeScriptProject, filePath *string, options *ReactTypeDefOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.ReactTypeDef",
		[]interface{}{project, filePath, options},
		r,
	)
}

func (j *jsiiProxy_ReactTypeDef) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_ReactTypeDef) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func (r *jsiiProxy_ReactTypeDef) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReactTypeDef) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReactTypeDef) Synthesize() {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReactTypeDef) SynthesizeContent(_arg projen.IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"synthesizeContent",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

