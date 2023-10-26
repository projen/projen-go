package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// A component representing the projen runtime configuration.
// Experimental.
type ProjenrcFile interface {
	Component
	// The path of the projenrc file.
	// Experimental.
	FilePath() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() Project
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ProjenrcFile
type jsiiProxy_ProjenrcFile struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_ProjenrcFile) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjenrcFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjenrcFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjenrcFile_Override(p ProjenrcFile, scope constructs.IConstruct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.ProjenrcFile",
		[]interface{}{scope, id},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func ProjenrcFile_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjenrcFile_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.ProjenrcFile",
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
func ProjenrcFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjenrcFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.ProjenrcFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns the `Projenrc` instance associated with a project or `undefined` if there is no Projenrc.
//
// Returns: A Projenrc.
// Experimental.
func ProjenrcFile_Of(project Project) ProjenrcFile {
	_init_.Initialize()

	if err := validateProjenrcFile_OfParameters(project); err != nil {
		panic(err)
	}
	var returns ProjenrcFile

	_jsii_.StaticInvoke(
		"projen.ProjenrcFile",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjenrcFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcFile) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

