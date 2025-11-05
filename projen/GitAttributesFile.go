package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Assign attributes to file names in a git repository.
// See: https://git-scm.com/docs/gitattributes
//
// Experimental.
type GitAttributesFile interface {
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
	// The default end of line character for text files.
	// Experimental.
	EndOfLine() EndOfLine
	// Indicates if the file should be marked as executable.
	// Experimental.
	Executable() *bool
	// Experimental.
	SetExecutable(val *bool)
	// Whether the current gitattributes file has any LFS patterns.
	// Experimental.
	HasLfsPatterns() *bool
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
	// Maps a set of attributes to a set of files.
	// Experimental.
	AddAttributes(glob *string, attributes ...*string)
	// Add attributes necessary to mark these files as stored in LFS.
	// Experimental.
	AddLfsPattern(glob *string)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Removes attributes from a set of files.
	//
	// If no attributes are provided, the glob pattern will be removed completely.
	// Experimental.
	RemoveAttributes(glob *string, attributes ...*string)
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

// The jsii proxy struct for GitAttributesFile
type jsiiProxy_GitAttributesFile struct {
	jsiiProxy_FileBase
}

func (j *jsiiProxy_GitAttributesFile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) EndOfLine() EndOfLine {
	var returns EndOfLine
	_jsii_.Get(
		j,
		"endOfLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) HasLfsPatterns() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasLfsPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitAttributesFile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitAttributesFile(scope constructs.IConstruct, options *GitAttributesFileOptions) GitAttributesFile {
	_init_.Initialize()

	if err := validateNewGitAttributesFileParameters(scope, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitAttributesFile{}

	_jsii_.Create(
		"projen.GitAttributesFile",
		[]interface{}{scope, options},
		&j,
	)

	return &j
}

// Experimental.
func NewGitAttributesFile_Override(g GitAttributesFile, scope constructs.IConstruct, options *GitAttributesFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.GitAttributesFile",
		[]interface{}{scope, options},
		g,
	)
}

func (j *jsiiProxy_GitAttributesFile)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_GitAttributesFile)SetReadonly(val *bool) {
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
func GitAttributesFile_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGitAttributesFile_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.GitAttributesFile",
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
func GitAttributesFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGitAttributesFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.GitAttributesFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitAttributesFile) AddAttributes(glob *string, attributes ...*string) {
	if err := g.validateAddAttributesParameters(glob); err != nil {
		panic(err)
	}
	args := []interface{}{glob}
	for _, a := range attributes {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addAttributes",
		args,
	)
}

func (g *jsiiProxy_GitAttributesFile) AddLfsPattern(glob *string) {
	if err := g.validateAddLfsPatternParameters(glob); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addLfsPattern",
		[]interface{}{glob},
	)
}

func (g *jsiiProxy_GitAttributesFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitAttributesFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitAttributesFile) RemoveAttributes(glob *string, attributes ...*string) {
	if err := g.validateRemoveAttributesParameters(glob); err != nil {
		panic(err)
	}
	args := []interface{}{glob}
	for _, a := range attributes {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"removeAttributes",
		args,
	)
}

func (g *jsiiProxy_GitAttributesFile) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitAttributesFile) SynthesizeContent(resolver IResolver) *string {
	if err := g.validateSynthesizeContentParameters(resolver); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitAttributesFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

