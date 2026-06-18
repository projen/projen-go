package web

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
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
	// Indicates if the file will be committed.
	// Deprecated: No longer used.
	Committed() *bool
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
	// The tree node.
	// Deprecated: No longer used.
	Node() constructs.Node
	// The file path, relative to the project's outdir.
	// Deprecated: No longer used.
	Path() *string
	// Deprecated: No longer used.
	Project() projen.Project
	// Indicates if the file should be read-only or read-write.
	// Deprecated: No longer used.
	Readonly() *bool
	// Deprecated: No longer used.
	SetReadonly(val *bool)
	// Returns a unified diff of the old and new file contents with context lines and hunk headers.
	//
	// Only available after synthesis.
	//
	// This is an expensive operation and should only be used on non time-critical
	// code paths, like debug output.
	//
	// Returns: the diff as an array of lines, or `undefined` if the file was
	// not changed or has not been synthesized yet.
	// Default: 3.
	//
	// Deprecated: No longer used.
	Diff(colorize *bool, contextLines *float64) *[]*string
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
	SynthesizeContent(resolver projen.IResolver) *string
	// Returns a string representation of this construct.
	// Deprecated: No longer used.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Deprecated: No longer used.
	With(mixins ...constructs.IMixin) constructs.IConstruct
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

func (j *jsiiProxy_ReactTypeDef) Committed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"committed",
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

func (j *jsiiProxy_ReactTypeDef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

	if err := validateNewReactTypeDefParameters(project, filePath, options); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_ReactTypeDef)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_ReactTypeDef)SetReadonly(val *bool) {
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
// Deprecated: No longer used.
func ReactTypeDef_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateReactTypeDef_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.web.ReactTypeDef",
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
// Deprecated: No longer used.
func ReactTypeDef_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateReactTypeDef_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.web.ReactTypeDef",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReactTypeDef) Diff(colorize *bool, contextLines *float64) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"diff",
		[]interface{}{colorize, contextLines},
		&returns,
	)

	return returns
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

func (r *jsiiProxy_ReactTypeDef) SynthesizeContent(resolver projen.IResolver) *string {
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

func (r *jsiiProxy_ReactTypeDef) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReactTypeDef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		r,
		"with",
		args,
		&returns,
	)

	return returns
}

