package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a source file.
// Experimental.
type SourceCode interface {
	Component
	// Experimental.
	FilePath() *string
	// Experimental.
	Marker() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() Project
	// Decreases the indentation level and closes a code block.
	// Experimental.
	Close(code *string)
	// Emit a line of code.
	// Experimental.
	Line(code *string)
	// Opens a code block and increases the indentation level.
	// Experimental.
	Open(code *string)
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

// The jsii proxy struct for SourceCode
type jsiiProxy_SourceCode struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_SourceCode) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceCode) Marker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceCode) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SourceCode) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewSourceCode(project Project, filePath *string, options *SourceCodeOptions) SourceCode {
	_init_.Initialize()

	if err := validateNewSourceCodeParameters(project, filePath, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SourceCode{}

	_jsii_.Create(
		"projen.SourceCode",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewSourceCode_Override(s SourceCode, project Project, filePath *string, options *SourceCodeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.SourceCode",
		[]interface{}{project, filePath, options},
		s,
	)
}

// Test whether the given construct is a component.
// Experimental.
func SourceCode_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSourceCode_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.SourceCode",
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
func SourceCode_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSourceCode_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.SourceCode",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SourceCode) Close(code *string) {
	_jsii_.InvokeVoid(
		s,
		"close",
		[]interface{}{code},
	)
}

func (s *jsiiProxy_SourceCode) Line(code *string) {
	_jsii_.InvokeVoid(
		s,
		"line",
		[]interface{}{code},
	)
}

func (s *jsiiProxy_SourceCode) Open(code *string) {
	_jsii_.InvokeVoid(
		s,
		"open",
		[]interface{}{code},
	)
}

func (s *jsiiProxy_SourceCode) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SourceCode) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SourceCode) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SourceCode) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

