// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Minimal Makefile.
// Experimental.
type Makefile interface {
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
	// List of rule definitions.
	// Experimental.
	Rules() *[]*Rule
	// Add a target to all.
	// Experimental.
	AddAll(target *string) Makefile
	// Add multiple targets to all.
	// Experimental.
	AddAlls(targets ...*string) Makefile
	// Add a rule to the Makefile.
	// Experimental.
	AddRule(rule *Rule) Makefile
	// Add multiple rules to the Makefile.
	// Experimental.
	AddRules(rules ...*Rule) Makefile
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
}

// The jsii proxy struct for Makefile
type jsiiProxy_Makefile struct {
	jsiiProxy_FileBase
}

func (j *jsiiProxy_Makefile) AbsolutePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absolutePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Changed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"changed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Executable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"executable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Makefile) Rules() *[]*Rule {
	var returns *[]*Rule
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}


// Experimental.
func NewMakefile(project Project, filePath *string, options *MakefileOptions) Makefile {
	_init_.Initialize()

	j := jsiiProxy_Makefile{}

	_jsii_.Create(
		"projen.Makefile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewMakefile_Override(m Makefile, project Project, filePath *string, options *MakefileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Makefile",
		[]interface{}{project, filePath, options},
		m,
	)
}

func (j *jsiiProxy_Makefile) SetExecutable(val *bool) {
	_jsii_.Set(
		j,
		"executable",
		val,
	)
}

func (j *jsiiProxy_Makefile) SetReadonly(val *bool) {
	_jsii_.Set(
		j,
		"readonly",
		val,
	)
}

func (m *jsiiProxy_Makefile) AddAll(target *string) Makefile {
	var returns Makefile

	_jsii_.Invoke(
		m,
		"addAll",
		[]interface{}{target},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Makefile) AddAlls(targets ...*string) Makefile {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	var returns Makefile

	_jsii_.Invoke(
		m,
		"addAlls",
		args,
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Makefile) AddRule(rule *Rule) Makefile {
	var returns Makefile

	_jsii_.Invoke(
		m,
		"addRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Makefile) AddRules(rules ...*Rule) Makefile {
	args := []interface{}{}
	for _, a := range rules {
		args = append(args, a)
	}

	var returns Makefile

	_jsii_.Invoke(
		m,
		"addRules",
		args,
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Makefile) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Makefile) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Makefile) Synthesize() {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Makefile) SynthesizeContent(resolver IResolver) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"synthesizeContent",
		[]interface{}{resolver},
		&returns,
	)

	return returns
}

