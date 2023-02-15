// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Experimental.
type IgnoreFile interface {
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
	// Add ignore patterns.
	//
	// Files that match this pattern will be ignored. If the
	// pattern starts with a negation mark `!`, files that match will _not_ be
	// ignored.
	//
	// Comment lines (start with `#`) are ignored.
	// Experimental.
	AddPatterns(patterns ...*string)
	// Ignore the files that match these patterns.
	// Experimental.
	Exclude(patterns ...*string)
	// Always include the specified file patterns.
	// Experimental.
	Include(patterns ...*string)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Removes patterns previously added from the ignore file.
	//
	// If `addPattern()` is called after this, the pattern will be added again.
	// Experimental.
	RemovePatterns(patterns ...*string)
	// Writes the file to the project's output directory.
	// Experimental.
	Synthesize()
	// Implemented by derived classes and returns the contents of the file to emit.
	// Experimental.
	SynthesizeContent(resolver IResolver) *string
}

// The jsii proxy struct for IgnoreFile
type jsiiProxy_IgnoreFile struct {
	jsiiProxy_FileBase
}

func (j *jsiiProxy_IgnoreFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IgnoreFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IgnoreFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IgnoreFile) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IgnoreFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IgnoreFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IgnoreFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewIgnoreFile(project Project, filePath *string) IgnoreFile {
	_init_.Initialize()

	if err := validateNewIgnoreFileParameters(project, filePath); err != nil {
		panic(err)
	}
	j := jsiiProxy_IgnoreFile{}

	_jsii_.Create(
		"projen.IgnoreFile",
		[]interface{}{project, filePath},
		&j,
	)

	return &j
}

// Experimental.
func NewIgnoreFile_Override(i IgnoreFile, project Project, filePath *string) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.IgnoreFile",
		[]interface{}{project, filePath},
		i,
	)
}

func (j *jsiiProxy_IgnoreFile)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_IgnoreFile)SetReadonly(val *bool) {
	if err := j.validateSetReadonlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func (i *jsiiProxy_IgnoreFile) AddPatterns(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addPatterns",
		args,
	)
}

func (i *jsiiProxy_IgnoreFile) Exclude(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"exclude",
		args,
	)
}

func (i *jsiiProxy_IgnoreFile) Include(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"include",
		args,
	)
}

func (i *jsiiProxy_IgnoreFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IgnoreFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IgnoreFile) RemovePatterns(patterns ...*string) {
	args := []interface{}{}
	for _, a := range patterns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"removePatterns",
		args,
	)
}

func (i *jsiiProxy_IgnoreFile) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IgnoreFile) SynthesizeContent(resolver IResolver) *string {
	if err := i.validateSynthesizeContentParameters(resolver); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

