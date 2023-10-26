package typescript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/typescript/internal"
)

// A projenrc file written in TypeScript.
//
// This component can be instantiated in any type of project
// and has no expectations around the project's main language.
//
// Requires that `npx` is available.
// Experimental.
type ProjenrcTs interface {
	projen.ProjenrcFile
	// The path of the projenrc file.
	// Experimental.
	FilePath() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// TypeScript configuration file used to compile projen source files.
	// Experimental.
	Tsconfig() javascript.TypescriptConfig
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

// The jsii proxy struct for ProjenrcTs
type jsiiProxy_ProjenrcTs struct {
	internal.Type__projenProjenrcFile
}

func (j *jsiiProxy_ProjenrcTs) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjenrcTs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjenrcTs) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjenrcTs) Tsconfig() javascript.TypescriptConfig {
	var returns javascript.TypescriptConfig
	_jsii_.Get(
		j,
		"tsconfig",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjenrcTs(project projen.Project, options *ProjenrcTsOptions) ProjenrcTs {
	_init_.Initialize()

	if err := validateNewProjenrcTsParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjenrcTs{}

	_jsii_.Create(
		"projen.typescript.ProjenrcTs",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewProjenrcTs_Override(p ProjenrcTs, project projen.Project, options *ProjenrcTsOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.typescript.ProjenrcTs",
		[]interface{}{project, options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func ProjenrcTs_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjenrcTs_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.typescript.ProjenrcTs",
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
func ProjenrcTs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjenrcTs_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.typescript.ProjenrcTs",
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
func ProjenrcTs_Of(project projen.Project) projen.ProjenrcFile {
	_init_.Initialize()

	if err := validateProjenrcTs_OfParameters(project); err != nil {
		panic(err)
	}
	var returns projen.ProjenrcFile

	_jsii_.StaticInvoke(
		"projen.typescript.ProjenrcTs",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjenrcTs) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcTs) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcTs) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcTs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

