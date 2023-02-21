// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
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
	// Writes the file to the project's output directory.
	// Experimental.
	Synthesize()
	// Implemented by derived classes and returns the contents of the file to emit.
	// Experimental.
	SynthesizeContent(_arg IResolver) *string
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
func NewGitAttributesFile(project Project) GitAttributesFile {
	_init_.Initialize()

	if err := validateNewGitAttributesFileParameters(project); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitAttributesFile{}

	_jsii_.Create(
		"projen.GitAttributesFile",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewGitAttributesFile_Override(g GitAttributesFile, project Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.GitAttributesFile",
		[]interface{}{project},
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

func (g *jsiiProxy_GitAttributesFile) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitAttributesFile) SynthesizeContent(_arg IResolver) *string {
	if err := g.validateSynthesizeContentParameters(_arg); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"synthesizeContent",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

