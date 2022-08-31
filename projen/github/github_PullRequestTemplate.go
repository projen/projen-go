package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Template for GitHub pull requests.
// Experimental.
type PullRequestTemplate interface {
	projen.TextFile
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
	Project() projen.Project
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
	SynthesizeContent(_arg projen.IResolver) *string
}

// The jsii proxy struct for PullRequestTemplate
type jsiiProxy_PullRequestTemplate struct {
	internal.Type__projenTextFile
}

func (j *jsiiProxy_PullRequestTemplate) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PullRequestTemplate) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}


// Experimental.
func NewPullRequestTemplate(github GitHub, options *PullRequestTemplateOptions) PullRequestTemplate {
	_init_.Initialize()

	if err := validateNewPullRequestTemplateParameters(github, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PullRequestTemplate{}

	_jsii_.Create(
		"projen.github.PullRequestTemplate",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPullRequestTemplate_Override(p PullRequestTemplate, github GitHub, options *PullRequestTemplateOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.PullRequestTemplate",
		[]interface{}{github, options},
		p,
	)
}

func (j *jsiiProxy_PullRequestTemplate)SetExecutable(val *bool) {
	if err := j.validateSetExecutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_PullRequestTemplate)SetReadonly(val *bool) {
	if err := j.validateSetReadonlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func (p *jsiiProxy_PullRequestTemplate) AddLine(line *string) {
	if err := p.validateAddLineParameters(line); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addLine",
		[]interface{}{line},
	)
}

func (p *jsiiProxy_PullRequestTemplate) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestTemplate) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestTemplate) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PullRequestTemplate) SynthesizeContent(_arg projen.IResolver) *string {
	if err := p.validateSynthesizeContentParameters(_arg); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"synthesizeContent",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

