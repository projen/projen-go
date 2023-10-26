package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Sets up a project to use JSON for projenrc.
// Experimental.
type ProjenrcJson interface {
	ProjenrcFile
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

// The jsii proxy struct for ProjenrcJson
type jsiiProxy_ProjenrcJson struct {
	jsiiProxy_ProjenrcFile
}

func (j *jsiiProxy_ProjenrcJson) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjenrcJson) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjenrcJson) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjenrcJson(project Project, options *ProjenrcJsonOptions) ProjenrcJson {
	_init_.Initialize()

	if err := validateNewProjenrcJsonParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjenrcJson{}

	_jsii_.Create(
		"projen.ProjenrcJson",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewProjenrcJson_Override(p ProjenrcJson, project Project, options *ProjenrcJsonOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.ProjenrcJson",
		[]interface{}{project, options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func ProjenrcJson_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjenrcJson_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.ProjenrcJson",
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
func ProjenrcJson_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjenrcJson_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.ProjenrcJson",
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
func ProjenrcJson_Of(project Project) ProjenrcFile {
	_init_.Initialize()

	if err := validateProjenrcJson_OfParameters(project); err != nil {
		panic(err)
	}
	var returns ProjenrcFile

	_jsii_.StaticInvoke(
		"projen.ProjenrcJson",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjenrcJson) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcJson) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcJson) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjenrcJson) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

