package java

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/java/internal"
)

// Configures a maven project to produce a .jar archive with sources and javadocs.
// Experimental.
type MavenPackaging interface {
	projen.Component
	// The directory containing the package output, relative to the project outdir.
	// Experimental.
	Distdir() *string
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

// The jsii proxy struct for MavenPackaging
type jsiiProxy_MavenPackaging struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_MavenPackaging) Distdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MavenPackaging) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MavenPackaging) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewMavenPackaging(project projen.Project, pom Pom, options *MavenPackagingOptions) MavenPackaging {
	_init_.Initialize()

	if err := validateNewMavenPackagingParameters(project, pom, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_MavenPackaging{}

	_jsii_.Create(
		"projen.java.MavenPackaging",
		[]interface{}{project, pom, options},
		&j,
	)

	return &j
}

// Experimental.
func NewMavenPackaging_Override(m MavenPackaging, project projen.Project, pom Pom, options *MavenPackagingOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.MavenPackaging",
		[]interface{}{project, pom, options},
		m,
	)
}

// Test whether the given construct is a component.
// Experimental.
func MavenPackaging_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMavenPackaging_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.java.MavenPackaging",
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
func MavenPackaging_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMavenPackaging_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.java.MavenPackaging",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MavenPackaging) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MavenPackaging) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MavenPackaging) Synthesize() {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MavenPackaging) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

