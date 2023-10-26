package java

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/java/internal"
)

// Allows writing projenrc files in java.
//
// This will install `org.projen/projen` as a Maven dependency and will add a
// `synth` task which will compile & execute `main()` from
// `src/main/java/projenrc.java`.
// Experimental.
type Projenrc interface {
	projen.ProjenrcFile
	// The name of the java class that includes the projen entrypoint.
	// Experimental.
	ClassName() *string
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

func (j *jsiiProxy_Projenrc) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
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
func NewProjenrc(project projen.Project, pom Pom, options *ProjenrcOptions) Projenrc {
	_init_.Initialize()

	if err := validateNewProjenrcParameters(project, pom, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Projenrc{}

	_jsii_.Create(
		"projen.java.Projenrc",
		[]interface{}{project, pom, options},
		&j,
	)

	return &j
}

// Experimental.
func NewProjenrc_Override(p Projenrc, project projen.Project, pom Pom, options *ProjenrcOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.Projenrc",
		[]interface{}{project, pom, options},
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
		"projen.java.Projenrc",
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
func Projenrc_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjenrc_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.java.Projenrc",
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
		"projen.java.Projenrc",
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

