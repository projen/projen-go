package java

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/java/internal"
)

// Adds the maven-compiler plugin to a POM file and the `compile` task.
// Experimental.
type MavenCompile interface {
	projen.Component
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
}

// The jsii proxy struct for MavenCompile
type jsiiProxy_MavenCompile struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_MavenCompile) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewMavenCompile(project projen.Project, pom Pom, options *MavenCompileOptions) MavenCompile {
	_init_.Initialize()

	if err := validateNewMavenCompileParameters(project, pom, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_MavenCompile{}

	_jsii_.Create(
		"projen.java.MavenCompile",
		[]interface{}{project, pom, options},
		&j,
	)

	return &j
}

// Experimental.
func NewMavenCompile_Override(m MavenCompile, project projen.Project, pom Pom, options *MavenCompileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.MavenCompile",
		[]interface{}{project, pom, options},
		m,
	)
}

func (m *jsiiProxy_MavenCompile) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MavenCompile) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MavenCompile) Synthesize() {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		nil, // no parameters
	)
}

