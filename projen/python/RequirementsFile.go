package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Specifies a list of packages to be installed using pip.
// See: https://pip.pypa.io/en/stable/reference/pip_install/#requirements-file-format
//
// Experimental.
type RequirementsFile interface {
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
	// Adds the specified packages provided in semver format.
	//
	// Comment lines (start with `#`) are ignored.
	// Experimental.
	AddPackages(packages ...*string)
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

// The jsii proxy struct for RequirementsFile
type jsiiProxy_RequirementsFile struct {
	internal.Type__projenFileBase
}

func (j *jsiiProxy_RequirementsFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequirementsFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewRequirementsFile(project projen.Project, filePath *string, options *RequirementsFileOptions) RequirementsFile {
	_init_.Initialize()

	if err := validateNewRequirementsFileParameters(project, filePath, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_RequirementsFile{}

	_jsii_.Create(
		"projen.python.RequirementsFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewRequirementsFile_Override(r RequirementsFile, project projen.Project, filePath *string, options *RequirementsFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.RequirementsFile",
		[]interface{}{project, filePath, options},
		r,
	)
}

func (j *jsiiProxy_RequirementsFile)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_RequirementsFile)SetReadonly(val *bool) {
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
func RequirementsFile_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRequirementsFile_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.RequirementsFile",
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
func RequirementsFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRequirementsFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.RequirementsFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequirementsFile) AddPackages(packages ...*string) {
	args := []interface{}{}
	for _, a := range packages {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addPackages",
		args,
	)
}

func (r *jsiiProxy_RequirementsFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RequirementsFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RequirementsFile) Synthesize() {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RequirementsFile) SynthesizeContent(resolver projen.IResolver) *string {
	if err := r.validateSynthesizeContentParameters(resolver); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequirementsFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

