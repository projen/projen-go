package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Sets up a javascript project to use TypeScript for projenrc.
// Experimental.
type Projenrc interface {
	projen.ProjenrcFile
	// The path of the projenrc file.
	// Experimental.
	FilePath() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
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

// The jsii proxy struct for Projenrc
type jsiiProxy_Projenrc struct {
	internal.Type__projenProjenrcFile
}

func (j *jsiiProxy_Projenrc) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Projenrc) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Projenrc) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjenrc(project projen.Project, options *ProjenrcOptions) Projenrc {
	_init_.Initialize()

	if err := validateNewProjenrcParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Projenrc{}

	_jsii_.Create(
		"projen.javascript.Projenrc",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewProjenrc_Override(p Projenrc, project projen.Project, options *ProjenrcOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Projenrc",
		[]interface{}{project, options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Projenrc_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjenrc_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Projenrc",
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
func Projenrc_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjenrc_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Projenrc",
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
func Projenrc_Of(project projen.Project) projen.ProjenrcFile {
	_init_.Initialize()

	if err := validateProjenrc_OfParameters(project); err != nil {
		panic(err)
	}
	var returns projen.ProjenrcFile

	_jsii_.StaticInvoke(
		"projen.javascript.Projenrc",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Projenrc) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Projenrc) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Projenrc) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Projenrc) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

